package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"log"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	defer cli.Close()

	kv := clientv3.NewKV(cli)
	getOp := clientv3.OpGet("key", clientv3.WithPrefix())
	opReponse, err := kv.Do(context.TODO(), getOp)
	if err != nil {
		log.Fatal("error ", err)
	}
	for k, v := range opReponse.Get().Kvs {
		fmt.Println(k, v)
	}

	opReponse, err = kv.Do(context.TODO(), clientv3.OpDelete("key"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(opReponse.Del().Deleted)

	lease := clientv3.NewLease(cli)
	leaseResponse, err := lease.Grant(context.TODO(), 2)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(3 * time.Second)

	fmt.Println("ttl ", leaseResponse.TTL)

	fmt.Println(leaseResponse.TTL)
	leaseChan, err := lease.KeepAlive(context.TODO(), leaseResponse.ID)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case v := <-leaseChan:
				if v == nil {
					fmt.Println("v is nil")
					goto ERR
				}
				fmt.Println(v)
			}
		}
	ERR:
	}()

	time.Sleep(30 * time.Second)

}
