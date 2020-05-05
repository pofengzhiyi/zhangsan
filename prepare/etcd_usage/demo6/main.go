package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	var (
		config         clientv3.Config
		client         *clientv3.Client
		err            error
		lease          clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId        clientv3.LeaseID
		putResp        *clientv3.PutResponse
		getResp        *clientv3.GetResponse
		keepResp       *clientv3.LeaseKeepAliveResponse
		keepRespChan   <-chan *clientv3.LeaseKeepAliveResponse
		kv             clientv3.KV
	)

	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"}, //集群服务
		DialTimeout: 5 * time.Second,
	}

	// 建立一个客户端
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	// 申请一个lease（租约）
	lease = clientv3.NewLease(client)

	//申请一个10秒的租约
	if leaseGrantResp, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}

	//拿到租约的ID
	leaseId = leaseGrantResp.ID

	//自动续租
	//ctx,_ := context.WithTimeout(context.TODO(),5 * time. Second)

	//续租了5秒，停止续租，10秒的生命周期 = 15秒的声明周期

	//5秒后取消自动续租
	if keepRespChan, err = lease.KeepAlive(context.TODO(), leaseId); err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			select {
			case keepResp = <-keepRespChan:
				if keepRespChan == nil {
					fmt.Println("租约已经失效")
					goto END
				} else { //每秒会续租，所以一次答应
					fmt.Println("收到自动续租应答", keepResp.ID)
				}
			}
		}
	END:
	}()

	//获取kv API子集
	kv = clientv3.NewKV(client)

	//Put一个KV，让它与租约关联起来，从而实现10秒自动退出
	if putResp, err = kv.Put(context.TODO(), "/cron/lock/job1", "", clientv3.WithLease(leaseId)); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("写入成功", putResp.Header.Revision)

	//定时看一下key过期没有
	for {
		if getResp, err = kv.Get(context.TODO(), "/cron/lock/job1"); err != nil {
			fmt.Println(err)
			return
		}
		if getResp.Count == 0 {
			fmt.Println("kv过期了")
			break
		}
		fmt.Println("还没过期", getResp.Kvs)
		time.Sleep(2 * time.Second)
	}

}
