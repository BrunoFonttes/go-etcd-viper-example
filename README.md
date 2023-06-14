# go-etcd-viper-example

Example of dynamic env loading in go using etcd and viper

How to run:

```
docker run -d -v /usr/share/ca-certificates/:/etc/ssl/certs -p 4001:4001 -p 2380:2380 -p 2379:2379 \
 --name etcd quay.io/coreos/etcd:v2.3.8 \
 -name etcd0 \
 -advertise-client-urls http://127.0.0.1:2379,http://127.0.0.1:4001 \
 -listen-client-urls http://0.0.0.0:2379,http://0.0.0.0:4001 \
 -initial-advertise-peer-urls http://127.0.0.1:2380 \
 -listen-peer-urls http://0.0.0.0:2380 \
 -initial-cluster-token etcd-cluster-1 \
 -initial-cluster etcd0=http://127.0.0.1:2380 \
 -initial-cluster-state new
 ```

Installing etcd:

Follow the instructions in [https://etcd.io/docs/v3.5/install/](https://etcd.io/docs/v3.5/install/)


Updating key with config file:

```
etcdctl put democonfig "{\       
    \"url\":\"https://helloworld.com11\",\
    \"port\":9005,\
    \"timeout\":13\
}"
```
Running app:

```
etcd_host=http://127.0.0.1:2379 etcd_watch_key=democonfig go run main.go
```
