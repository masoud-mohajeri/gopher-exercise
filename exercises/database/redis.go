package database

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func connectToRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client
}

func Redis() {
	client := connectToRedis()

	ctx := context.Background()

	client.Set(ctx, "SimpleKeyValue", "KV:Done!", 5*time.Second)
	val := client.Get(ctx, "SimpleKeyValue").Val()
	fmt.Println(val)

	listName := "SetTest1!"
	client.RPush(ctx, listName, "R1", "R2")
	client.LPush(ctx, listName, "L1", "L2")
	fmt.Println("POP from set:", client.LPop(ctx, listName).Val())
	fmt.Println("list Exists? ", client.Exists(ctx, listName).Val())
	fmt.Println("list Exists? ", client.Exists(ctx, "no").Val())

	sortedSetName := "ZSet1"
	client.ZAdd(ctx, sortedSetName, redis.Z{Score: 11, Member: "Eleven"})
	client.ZAdd(ctx, sortedSetName, redis.Z{Score: 1, Member: "One"})
	fmt.Println("ZRange ", client.ZRange(ctx, sortedSetName, 0, 12).Val())

	hashTableName := "myHashTable"
	client.HSet(ctx, hashTableName, "Zero", 0, "One", 1)
	fmt.Println("Get from hash table, Zero:", client.HGet(ctx, hashTableName, "Zero").Val())
}
