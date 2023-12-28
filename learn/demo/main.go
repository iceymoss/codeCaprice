package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

const lockKey = "lock"
const timerKey = "timer"
const listKey = "element_list"

func main() {
	// 创建 Redis 客户端连接
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379", // Redis 服务器地址
	})

	// 关闭连接
	defer client.Close()

	// 初始化元素列表
	initElementList(client)

	// 模拟不同请求
	for i := 0; i < 10; i++ {
		go makeRequest(client)
	}

	// 阻塞主程序
	select {}
}

func initElementList(client *redis.Client) {
	// 初始化元素列表
	elements := []string{"Element1", "Element2", "Element3", "Element4", "Element5", "Element6"}

	// 将元素列表写入 Redis List
	client.RPush(context.Background(), listKey, elements)
}

func makeRequest(client *redis.Client) {
	for {
		// 获取分布式锁
		ok, err := acquireLock(client, lockKey)
		if err != nil {
			fmt.Println("获取锁出错:", err)
			return
		}

		if ok {
			// 获取计时器状态
			timer, err := getTimerState(client, timerKey)
			if err != nil {
				fmt.Println("获取计时器状态出错:", err)
				releaseLock(client, lockKey)
				return
			}

			// 获取当前时间
			now := time.Now()

			// 判断是否需要返回下一个元素
			if now.After(timer.Expiration) {
				// 计时器已过期，可以返回下一个元素
				nextElement, err := getNextElement(client, listKey)
				if err != nil {
					fmt.Println("获取下一个元素出错:", err)
					releaseLock(client, lockKey)
					return
				}

				// 更新计时器状态，重置计时器
				updateTimerState(client, timerKey, 0, time.Now().Add(30*time.Minute), nextElement)

				// 打印返回的元素
				fmt.Printf("返回下一个元素: %s\n", nextElement)

				// 释放锁
				releaseLock(client, lockKey)
				return
			}

			// 计时器未过期，返回当前元素
			fmt.Printf("返回当前元素: %s\n", timer.Element)
			releaseLock(client, lockKey)
			return
		}

		// 释放锁
		releaseLock(client, lockKey)
	}
}

func acquireLock(client *redis.Client, key string) (bool, error) {
	// 使用 SETNX 操作尝试获取锁
	result, err := client.SetNX(context.Background(), key, "locked", 30*time.Second).Result()
	if err != nil {
		return false, err
	}
	return result, nil
}

func releaseLock(client *redis.Client, key string) error {
	// 使用 DEL 操作释放锁
	_, err := client.Del(context.Background(), key).Result()
	if err != nil {
		fmt.Println("释放锁出错:", err)
	}
	return err
}

type TimerState struct {
	Index      int
	Expiration time.Time
	Element    string
}

func getTimerState(client *redis.Client, key string) (*TimerState, error) {
	// 从 Redis 中获取计时器状态
	index, err := client.Get(context.Background(), key+"_index").Int()
	if err != nil {
		return nil, err
	}

	expiration, err := client.Get(context.Background(), key+"_expiration").Time()
	if err != nil {
		return nil, err
	}

	element, err := client.Get(context.Background(), key+"_element").Result()
	if err != nil {
		return nil, err
	}

	return &TimerState{
		Index:      index,
		Expiration: expiration,
		Element:    element,
	}, nil
}

func updateTimerState(client *redis.Client, key string, newIndex int, newExpiration time.Time, newElement string) {
	// 更新计时器状态
	client.Set(context.Background(), key+"_index", newIndex, 0)
	client.Set(context.Background(), key+"_expiration", newExpiration, 0)
	client.Set(context.Background(), key+"_element", newElement, 0)
}

func getNextElement(client *redis.Client, key string) (string, error) {
	// 从 Redis List 中获取下一个元素
	nextElement, err := client.LPop(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return nextElement, nil
}
