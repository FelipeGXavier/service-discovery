package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"go.etcd.io/etcd/clientv3"
)

func TestCallService(etcdClient *clientv3.Client) {
	_, err := etcdClient.Get(context.Background(), "/services/service1")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := etcdClient.Get(context.Background(), "", clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}

	serviceAddress := resp.Kvs[0]

	r, err := http.Get(fmt.Sprintf("http://%s/", string(serviceAddress.Value)))

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	log.Println(string(body))
}
