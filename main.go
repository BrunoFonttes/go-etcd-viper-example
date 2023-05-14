package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/viper"
	// _ "github.com/spf13/viper/remote"
	"go.etcd.io/etcd/clientv3"
)

func main() {
	fmt.Println("Hello")
	etcdHost := "http://127.0.0.1:2379"
	etcdWatchKey := "/democonfig"
	viper.AddRemoteProvider("etcd3", etcdHost, etcdWatchKey)
	// viper.SetConfigType("json") // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	// if err := viper.ReadRemoteConfig(); err != nil {
	// 	panic(err)
	// }

	fmt.Println("connecting to etcd - " + etcdHost)

	etcd, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://" + etcdHost},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("connected to etcd - " + etcdHost)

	defer etcd.Close()

	watchChan := etcd.Watch(context.Background(), etcdWatchKey)
	fmt.Println("set WATCH on " + etcdWatchKey)

	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			fmt.Printf("Event received! %s executed on %q with value %q\n", event.Type, event.Kv.Key, event.Kv.Value)
		}
	}

	http.ListenAndServe(":8090", nil)
}
