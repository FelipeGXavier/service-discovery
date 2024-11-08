package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/FelipeGXavier/service-discovery/etcd/go/service1/pkg/discovery"
	"github.com/FelipeGXavier/service-discovery/etcd/go/service1/pkg/netutils"
	"go.etcd.io/etcd/clientv3"
)

func main() {
	log.Println("start service 01")
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

	log.Println(ipv4Address + ":3333")

	err = discovery.RegisterService(etcdClient, "service1", ipv4Address+":3333")

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World")
	})

	err = http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}

}
