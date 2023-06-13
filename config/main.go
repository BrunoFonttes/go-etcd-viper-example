package config

import (
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type config struct {
	url     string
	port    int
	timeout int
}

func (c *config) Url() string  { return c.url }
func (c *config) Port() int    { return c.port }
func (c *config) Timeout() int { return c.timeout }

func (e config) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.url, validation.Required),
		validation.Field(&e.port, validation.Required),
		validation.Field(&e.timeout, validation.Required),
	)
}

func loadConf() error {
	Conf = config{
		url:     viper.GetString("url"),
		port:    viper.GetInt("port"),
		timeout: viper.GetInt("timeout"),
	}
	if err := Conf.Validate(); err != nil {
		return err
	}
	return nil
}

var Conf config

func Load() {
	etcdHost := "http://127.0.0.1:2379"
	etcdWatchKey := "/democonfig"

	viper.AddRemoteProvider("etcd", etcdHost, etcdWatchKey)
	viper.SetConfigType("json") // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"

	if err := viper.ReadRemoteConfig(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	go func() {
		for {
			if err := viper.WatchRemoteConfig(); err != nil {
				fmt.Printf("unable to read remote config: %v", err)
				continue
			}
			if err := loadConf(); err != nil {
				fmt.Println(err)
				panic(err)
			}
			fmt.Println(Conf.port, Conf.timeout, Conf.url)
			time.Sleep(time.Second * 10)
		}
	}()
}
