package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/291700351/core"
	"github.com/redis/go-redis/v9"
)

func main() {
	c := core.LoadRedisConfig("./config.yaml")

	helper := core.NewRedisHelper(c)
	defer helper.Release()

	helper.Set("Name", "Lee")
	helper.Set("Age", 33)
	helper.Set("Time", time.Now().Add(time.Hour*24))

	if ok, value := helper.GetString("Name"); ok {
		fmt.Println("name", value, "type", reflect.TypeOf(value))
	}
	if ok, value := helper.GetFloat32("Age"); ok {
		fmt.Println("age", value, "type", reflect.TypeOf(value))
	}
	if ok, value := helper.GetTime("Time"); ok {
		fmt.Println("time", value, "type", reflect.TypeOf(value))
	}

	if ok, value := helper.Keys("*"); ok {
		for _, v := range value {
			fmt.Println("v", v, "type", reflect.TypeOf(v))
		}
	}
	if ok, value := helper.Keys("*"); ok {
		for _, v := range value {
			helper.Delete(v)
		}
	}
	fmt.Println("--------------------------")

	if ok, value := helper.Keys("*"); ok {
		for _, v := range value {
			fmt.Println("v", v, "type", reflect.TypeOf(v))
		}
	}

	// clusterClient()

}

func c() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	defer rdb.Close()

	ctx := context.Background()
	v, e := rdb.Ping(ctx).Result()

	fmt.Println("value", v, "error", e)
}

func clusterClient() {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{Addrs: []string{"127.0.0.1:6379", "127.0.0.1:6379", "127.0.0.1:6379"}})

	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("连接集群出错", err)
		return
	}

	// rdb.ForEachShard(timeoutCtx, func(ctx context.Context, client *redis.Client) error {
	// 	r := client.Ping(ctx)
	// 	v, e := r.Result()
	// 	fmt.Println("value", v, "error", e)

	// 	return r.Err()
	// })
}
