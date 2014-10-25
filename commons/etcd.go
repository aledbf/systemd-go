package commons

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/coreos/go-etcd/etcd"
)

// Connect to etcd using environment variables
// COREOS_PRIVATE_IPV4
// ETCD_PORT
func ConnectToEtcd() *etcd.Client {
	host := Getopt("COREOS_PRIVATE_IPV4", "127.0.0.1")
	etcdPort := Getopt("ETCD_PORT", "4001")
	etcdUrl := "http://" + host + ":" + etcdPort
	log.Info("connecting to etcd in URL: %s", etcdUrl)
	return etcd.NewClient([]string{etcdUrl})
}

func SetDefaultEtcd(client *etcd.Client, key, value string) {
	_, err := client.Set(key, value, 0)
	if err != nil {
		log.Warn(err)
	}
}

func MkdirEtcd(client *etcd.Client, path string) {
	_, err := client.CreateDir(path, 0)
	if err != nil {
		log.Warn(err)
	}
}

// Wait for a set of keys to exist in etcd before return
func WaitForEtcdKeys(client *etcd.Client, keys []string) {
	wait := true

	for {
		for _, key := range keys {
			_, err := client.Get(key, false, false)
			if err != nil {
				log.Debugf("key \"%s\" error %v", key, err)
				wait = true
			}
		}

		if !wait {
			break
		}

		log.Println("waiting for missing etcd keys...")
		time.Sleep(1 * time.Second)
		wait = false
	}
}

// Get a value from etcd. If the value foesn't exists returns an empty string
func GetEtcd(client *etcd.Client, key string) string {
	result, err := client.Get(key, false, false)
	if err != nil {
		return ""
	}

	return result.Node.Value
}

// Return a list of values contained in a key or nil
func GetListEtcd(client *etcd.Client, key string) etcd.Nodes {
	result, err := client.Get(key, true, false)
	if err != nil {
		return nil
	}

	return result.Node.Nodes
}

func SetEtcd(client *etcd.Client, key, value string, ttl uint64) {
	_, err := client.Set(key, value, ttl)
	if err != nil {
		log.Debugf("%v", err)
	}
}

// Publish a service to etcd periodcally
func PublishService(
	client *etcd.Client,
	host string,
	etcdPath string,
	externalPort string,
	ttl uint64,
	timeout time.Duration) {

	for {
		SetEtcd(client, etcdPath+"/host", host, ttl)
		SetEtcd(client, etcdPath+"/port", externalPort, ttl)
		time.Sleep(timeout)
	}
}
