package web

import (
	"log"
	"time"

	"github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/net/context"
	"github.com/coreos/etcd/client"
)

func GetEtcdInfo() {
	cfg := client.Config{
		Endpoints: []string{"http://cloud_0001:2379", "http://cloud_0002:2379", "http://cloud_0003:2379"},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi := client.NewKeysAPI(c)
	// set "/goo" key with "bar" value
	log.Print("Setting '/goo' key with 'bar' value")
	resp, err := kapi.Set(context.Background(), "/goo", "Gooooooooooooooooooooooooooogle", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Set is done. Metadata is %q\n", resp)
	}
	// get "/foo" key's value
	log.Print("Getting '/goo' key value")
	resp, err = kapi.Get(context.Background(), "/goo", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Get is done. Metadata is %q\n", resp)
		// print value
		log.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
	}
}

//func main() {
//	GetEtcdInfo()
//}
