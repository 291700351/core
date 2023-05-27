package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/291700351/core"
	"github.com/291700351/core/examples/dao"
	"github.com/291700351/core/examples/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	//core.InitMysql(logger.Info, "root", "99233123", "127.0.0.1", 3306, "test", nil)
	core.InitSqlite()
	core.GetOrmEngine().AutoMigrate(model.User{})
	userDao := dao.Use(core.GetOrmEngine()).User

	if u, err := userDao.WithContext(context.Background()).Where(userDao.Name.Like("%Lee%")).Find(); nil == err {
		for i, u2 := range u {
			fmt.Printf("index = %d, u = %v\n-------\n", i, u2)
		}
	}

}

type User struct {
	gorm.Model
	Id   int64
	Name string
	Age  int
}

type UserDao struct {
	core.Dao[User]
}

func testGen() {
	core.InitMysql(logger.Info, "root", "99233123", "127.0.0.1", 3306, "test", nil)
	core.Gen()
}
func testDao() {
	core.InitMysql(logger.Info, "root", "99233123", "127.0.0.1", 3306, "test", nil)
	e := core.GetOrmEngine()

	e.AutoMigrate(&User{})
	dao := UserDao{}
	dao.DB = e

	// e.Sync(new(User))
	// dao := core.Dao[User]{
	// 	Db: e,
	// }

	user := User{
		Name: "Lee",
		Age:  33,
	}

	dao.Save(&user)

	if ok, u := dao.List(); ok {
		for _, v := range u {
			fmt.Println("--->", ":", v.Name)
		}
	}

}

func testRedis() {
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
