package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
		kv     clientv3.KV
		putOp  clientv3.Op
		getOp  clientv3.Op
		opResp clientv3.OpResponse
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

	kv = clientv3.NewKV(client)

	//OP operation
	putOp = clientv3.OpPut("/cron/jobs/job8", "12312312")

	//执行OP
	if opResp, err = kv.Do(context.TODO(), putOp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("写入Revision", opResp.Put().Header.Revision)

	getOp = clientv3.OpGet("/cron/jobs/job8")

	//执行OP
	if opResp, err = kv.Do(context.TODO(), getOp); err != nil {
		fmt.Println(err)
		return
	}

	//打印
	fmt.Println("数据Revision", opResp.Get().Kvs[0].ModRevision) //create rev= mod rev
	fmt.Println("数据value", opResp.Get().Kvs[0].Value)
}
