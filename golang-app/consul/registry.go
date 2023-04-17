package consul

import (
	"fmt"
	"log"
	"strconv"
	"os"
)

type ServiceDefinition struct {
	ConsulAddr  string
	ServiceId   string
	ServiceHost string
	ServicePort int
}

func GetEnv(name string, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultValue
	}
	return val
}

func GetEnvAsInt(name string, defaultValue int64) (int64, error) {
	val := os.Getenv(name)
	if val == "" {
		return defaultValue, nil
	}
	return strconv.ParseInt(val, 10, 0)
}

func init() {
	var sd ServiceDefinition
	host := GetEnv("CONSUL_HOST", "localhost")
	port, _ := GetEnvAsInt("CONSUL_PORT", 8500)
	name := GetEnv("APP_NAME", "GoApp")
	sd.ConsulAddr = host + ":" + strconv.FormatInt(port, 10)
	sd.ServiceId = name
	sd.ServiceHost = "0.0.0.0"
	sd.ServicePort = 8000

	cli, err := NewClient(sd.ConsulAddr)
	if err != nil {
		log.Fatalf("can't initiate consul client: %+v\n", err)
	}
	if err := cli.Register(sd.ServiceId); err != nil {
		log.Println("error registering... ", err)
	}
	log.Println("service registered: ", sd.ServiceId)

	consulStore := NewKVClient(cli)

	srvString := fmt.Sprintf("%s:%d", sd.ServiceHost, sd.ServicePort)

	if err := consulStore.PutKV("apigw", "apigw.example.com"); err != nil {
		log.Println("error creating key: ", err)

	}

	if err := consulStore.PutKV(sd.ServiceId, srvString); err != nil {
		log.Println("error creating key: ", err)

	}

	url, err := consulStore.GetKV("apigw")
	if err != nil {
		log.Println("could not get kv for apigw: ", err)
	}
	log.Println("apigw baseUrl: ", url)

}
