package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	ctx, calfunc := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = redisdb.Ping(ctx).Result()
	calfunc()
	return
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Println("connect redis failed,err:", err)
	}

	//zest
	key := "rank"
	items := []*redis.Z{ //redis.Z为众多redis数据结构中的一种
		&redis.Z{Score: 99, Member: "PHP"},
		&redis.Z{Score: 96, Member: "Golang"},
		&redis.Z{Score: 97, Member: "Python"},
		&redis.Z{Score: 99, Member: "Java"},
	}
	ctx, canl := context.WithTimeout(context.Background(), 10*time.Hour)
	defer canl()
	//把这些元素都追加到key
	redisdb.ZAdd(ctx, key, items...)
	newScore, err := redisdb.ZIncrBy(ctx, key, 10, "Golang").Result() //给golang加10分
	if err != nil {
		fmt.Println("zincrby failed,err:", err)
		return
	}
	fmt.Println("Golang's score is ", newScore)
	// 取分数最高的3个
	ret := redisdb.ZRevRangeWithScores(ctx, key, 0, 2).Val()
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
	// 取95~100分的
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = redisdb.ZRangeByScoreWithScores(ctx, key, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}
