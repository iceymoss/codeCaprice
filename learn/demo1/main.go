package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	rdb *redis.Client
)

const (
	redisHashKey = "bottom:index:timestamp"
)

func init() {
	// 初始化Redis客户端
	rdb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379", // Redis服务器地址
		DB:   0,                // 默认数据库
	})
}

var list = []interface{}{"网格", "抢新", "自定义", "套利", "大单", "胜率"}

func main() {
	app := gin.Default()
	app.GET("/get-element", bottom)

	// 启动Iris应用
	app.Run(":8080")
}

func bottom(ctx *gin.Context) {
	var timestamp int64
	var index int
	jsonData, err := rdb.HGet(context.Background(), redisHashKey, "data").Result()
	if err != nil {
		if err == redis.Nil {
			// 获取当前时间
			currentTime := time.Now()
			// 加上30分钟
			timestamp = currentTime.Add(10 * time.Second).Unix()
			index = 0
			err = saveIndexTimestampToRedis(index, timestamp)
			if err != nil {
				fmt.Println("写Redis失败", err.Error())
				ctx.JSON(200, gin.H{"error": fmt.Sprintf("Error: %v", err)})
				return
			}
			ctx.JSON(200, gin.H{"message": list[0]})
			return
		}
		fmt.Println("获取Redis失败")
		ctx.JSON(200, gin.H{"error": fmt.Sprintf("Error: %v", err)})
		return
	} else {

	}

	// 解析JSON字符串
	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		ctx.JSON(500, gin.H{"error": fmt.Sprintf("Error: %v", err)})
		return
	}

	// 从map中获取索引和时间戳
	index = int(data["index"].(float64))
	timestamp = int64(data["timestamp"].(float64))

	curTime := time.Now()
	if timestamp < curTime.Unix() {
		//当前index到期
		index++
		if index > 5 {
			index = 0
		}
		curTime = curTime.Add(10 * time.Second)
		saveIndexTimestampToRedis(index, curTime.Unix())
	}
	ctx.JSON(200, gin.H{"message": list[index]})
}

func saveIndexTimestampToRedis(index int, timestamp int64) error {
	// 构建一个map，包含索引和时间戳
	data := map[string]interface{}{
		"index":     index,
		"timestamp": timestamp,
	}

	// 将map转换为JSON字符串
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// 使用HSET将JSON字符串存储到Redis的Hash中
	err = rdb.HSet(context.Background(), redisHashKey, "data", jsonData).Err()
	if err != nil {
		return err
	}

	return nil
}

func getIndexTimestampFromRedis() (int, int64, error) {
	// 从Redis的Hash中获取存储的JSON字符串
	jsonData, err := rdb.HGet(context.Background(), redisHashKey, "data").Result()
	if err != nil {
		if err == redis.Nil {
			return 0, 0, redis.Nil // 键不存在
		}
		return 0, 0, err
	}

	// 解析JSON字符串
	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return 0, 0, err
	}

	// 从map中获取索引和时间戳
	index := int(data["index"].(float64))
	timestamp := int64(data["timestamp"].(float64))

	return index, timestamp, nil
}
