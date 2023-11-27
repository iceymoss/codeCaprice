package main

import (
	"fmt"
	"sync"
)

/*
【编程实现】问答路：在我们玩苗残的过程中，一股都有副本日053或古区城日口85，多个玩丽或者是大量的玩家间时攻击
该BOSS。在击杀日OSS后，会对击O85的这个玩家进行特殊奖励。使用居的系纸，件，设计实理一个接口，接受玩家的攻击
请求，保证击杀BOSS的玩家有且只有一个
*/

//思路：为了设计一个接口，确保击杀BOSS的玩家有且只有一个，可以考虑使用互斥锁（Mutex）来实现并发安全。Go语言实现：

// Boss 接口定义
type Boss interface {
	Attack(playerID int)
	KillBoss() (int, error)
}

// BossImpl 实现结构体
type BossImpl struct {
	mu         sync.Mutex     //互斥锁：保证并发安全
	lifeBar    int            //血条
	killed     bool           //是否被击杀
	killingID  int            //击杀的用户
	killReward map[int]string //击杀获取的奖励
}

// NewBossImpl 创建Boss实例
func NewBossImpl(lifeBar int) *BossImpl {
	return &BossImpl{
		lifeBar:    lifeBar,
		killReward: make(map[int]string),
	}
}

// Attack 玩家攻击BOSS
func (b *BossImpl) Attack(playerID int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.killed {
		// BOSS已经被击败
		fmt.Printf("玩家 %d 攻击了已被击败的BOSS\n", playerID)
		return
	}
	b.lifeBar--
	fmt.Printf("玩家 %d 攻击了BOSS\n", playerID)
	if b.lifeBar == 0 {
		fmt.Printf("玩家 %d 成功击败了BOSS\n", playerID)
		b.killingID = playerID
	}
}

// KillBoss 击败BOSS并获取奖励
func (b *BossImpl) KillBoss() (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.killed {
		// BOSS已经被击败
		return 0, fmt.Errorf("BOSS已被击败")
	}

	if b.killingID == 0 {
		// 没有玩家在攻击BOSS
		return 0, fmt.Errorf("没有玩家在攻击BOSS")
	}

	// BOSS被击败，记录奖励
	b.killed = true
	reward := b.killingID
	b.killReward[reward] = "特殊奖励"

	return reward, nil
}

func main() {
	boss := NewBossImpl(10)

	// 玩家攻击BOSS的并发例子
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(playerID int) {
			defer wg.Done()
			boss.Attack(playerID)
		}(i)
	}

	// 等待所有攻击完成
	wg.Wait()

	// 击败BOSS并获取奖励
	reward, err := boss.KillBoss()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("玩家 %d 获得了 %s\n", reward, boss.killReward[reward])
	}
}
