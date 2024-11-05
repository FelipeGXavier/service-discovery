package discovery

import (
	"context"
	"fmt"
	"log"

	"go.etcd.io/etcd/clientv3"
)

func RegisterService(etcdClient *clientv3.Client, serviceName string, serviceIpv4Address string) error {
	servicePrefix := fmt.Sprintf("/services/%s", serviceName)
	leaseTtl := int64(60)
	lease, err := etcdClient.Grant(context.Background(), leaseTtl)
	if err != nil {
		return err
	}
	_, err = etcdClient.Put(context.Background(), servicePrefix, serviceIpv4Address, clientv3.WithLease(lease.ID))
	if err != nil {
		return err
	}
	heartBeat(etcdClient, lease)
	return nil
}

func heartBeat(etcdClient *clientv3.Client, lease *clientv3.LeaseGrantResponse) error {
	keepAliveCh, err := etcdClient.KeepAlive(context.Background(), lease.ID)
	if err != nil {
		log.Fatal("error keeping lease alive:", err)
	}
	go func() {
		for {
			kaResp := <-keepAliveCh
			if kaResp == nil {
				log.Panicln("lease has expired")
				break
			} else {
				fmt.Println("lease renewed:", kaResp.TTL)
			}
		}
	}()
	return nil
}
