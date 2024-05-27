package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	// 加载TLS证书和私钥
	cert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	// 加载CA证书
	caCert, err := os.ReadFile("certs/etcd-root-ca.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// 创建TLS配置
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	fmt.Print(&tlsConfig)

	// 创建客户端配置
	cfg := clientv3.Config{
		Endpoints:   []string{"https://192.168.188.101:2379", "https://192.168.188.102:2379", "https://192.168.188.103:2379"},
		DialTimeout: 5 * time.Second,
		TLS:         tlsConfig,
	}

	// 创建客户端
	client, err := clientv3.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// 使用客户端进行操作，例如放置一个键值对
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = client.Put(ctx, "example_key", "example_value")
	if err != nil {
		log.Fatal(err)
	}

	// 获取刚才放置的键值对
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.Get(ctx, "hello")
	if err != nil {
		log.Fatal(err)
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s = %s\n", ev.Key, ev.Value)
	}
}
