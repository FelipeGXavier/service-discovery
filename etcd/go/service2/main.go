package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/FelipeGXavier/service-discovery/etcd/go/service1/pkg/discovery"
	"github.com/FelipeGXavier/service-discovery/etcd/go/service1/pkg/netutils"
	"go.etcd.io/etcd/clientv3"
)

func main() {
	log.Println("start service 02")
	etcdClient, err := clientv3.New(
		clientv3.Config{
			Endpoints:   []string{"localhost:2379"},
			DialTimeout: 60 * time.Second,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	defer etcdClient.Close()

	ipv4Address, err := netutils.GetFirstNonLoopbackAddress()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(ipv4Address)

	err = discovery.RegisterService(etcdClient, "service1", ipv4Address)

	if err != nil {
		log.Fatal(err)
	}

	TestCallService(etcdClient)

	resp, err := etcdClient.Get(context.Background(), "", clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through each key-value pair and print the key
	for _, kv := range resp.Kvs {
		fmt.Printf("Key: %s, Value: %s\n", kv.Key, kv.Value)
	}

	for {

	}

}
