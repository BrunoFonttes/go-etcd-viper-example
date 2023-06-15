# go-etcd-viper-example

Example of dynamic env loading in go using etcd and viper

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
