package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/pinpoint-apm/pinpoint-go-agent"
	"github.com/pinpoint-apm/pinpoint-go-agent/plugin/goredis"
	"github.com/pinpoint-apm/pinpoint-go-agent/plugin/http"
)

var redisClient *ppgoredis.Client
var redisClusterClient *ppgoredis.ClusterClient

func redisv6(w http.ResponseWriter, r *http.Request) {
	c := redisClient.WithContext(r.Context())
	redisPipeIncr(c.Pipeline())
}

func redisv6Cluster(w http.ResponseWriter, r *http.Request) {
	c := redisClusterClient.WithContext(r.Context())
	redisPipeIncr(c.Pipeline())
}

func redisPipeIncr(pipe redis.Pipeliner) {
	incr := pipe.Incr("foo")
	pipe.Expire("foo", time.Hour)
	_, er := pipe.Exec()
	fmt.Println(incr.Val(), er)
}

func main() {
	opts := []pinpoint.ConfigOption{
		pinpoint.WithAppName("GoRedisTest"),
		pinpoint.WithAgentId("GoRedisTestAgent"),
		pinpoint.WithConfigFile(os.Getenv("HOME") + "/tmp/pinpoint-config.yaml"),
	}
	cfg, _ := pinpoint.NewConfig(opts...)
	agent, err := pinpoint.NewAgent(cfg)
	if err != nil {
		log.Fatalf("pinpoint agent start fail: %v", err)
	}
	defer agent.Shutdown()

	addrs := []string{"localhost:6379", "localhost:6380"}

	//redis client
	redisOpts := &redis.Options{
		Addr: addrs[0],
	}
	redisClient = ppgoredis.NewClient(redisOpts)

	//redis cluster client
	redisClusterOpts := &redis.ClusterOptions{
		Addrs: addrs,
	}
	redisClusterClient = ppgoredis.NewClusterClient(redisClusterOpts)

	http.HandleFunc("/redis", pphttp.WrapHandlerFunc(redisv6))
	http.HandleFunc("/rediscluster", pphttp.WrapHandlerFunc(redisv6Cluster))

	http.ListenAndServe(":9000", nil)

	redisClient.Close()
	redisClusterClient.Close()
}
