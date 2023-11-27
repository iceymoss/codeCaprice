package main

import "fmt"

/*
软件工程】编程题：一个游戏角色有三种状态：行走、攻击、防街。在行走状态时。按A键进入攻击状态，按B键进入防御状态；
在攻击状态时，按A键进入行走状态，按B键进入防街状态；在防状态时，按A键进入攻击状态，按B键进入行走状态。设计实现
相关的功能类
*/

// State 接口定义
type State interface {
	PressA()
	PressB()
}

// GameCharacter 游戏角色
type GameCharacter struct {
	state State
}

// SetState 设置角色的状态
func (c *GameCharacter) SetState(newState State) {
	c.state = newState
}

// PressA 按A键
func (c *GameCharacter) PressA() {
	c.state.PressA()
}

// PressB 按B键
func (c *GameCharacter) PressB() {
	c.state.PressB()
}

// WalkState 行走状态
type WalkState struct {
	character *GameCharacter
}

// PressA 实现按A键的行为
func (s *WalkState) PressA() {
	fmt.Println("进入攻击状态")
	s.character.SetState(&AttackState{s.character})
}

// PressB 实现按B键的行为
func (s *WalkState) PressB() {
	fmt.Println("进入防御状态")
	s.character.SetState(&DefendState{s.character})
}

// AttackState 攻击状态
type AttackState struct {
	character *GameCharacter
}

// PressA 实现按A键的行为
func (s *AttackState) PressA() {
	fmt.Println("进入行走状态")
	s.character.SetState(&WalkState{s.character})
}

// PressB 实现按B键的行为
func (s *AttackState) PressB() {
	fmt.Println("进入防御状态")
	s.character.SetState(&DefendState{s.character})
}

// DefendState 防御状态
type DefendState struct {
	character *GameCharacter
}

// PressA 实现按A键的行为
func (s *DefendState) PressA() {
	fmt.Println("进入攻击状态")
	s.character.SetState(&AttackState{s.character})
}

// PressB 实现按B键的行为
func (s *DefendState) PressB() {
	fmt.Println("进入行走状态")
	s.character.SetState(&WalkState{s.character})
}

func main() {
	// 测试
	character := &GameCharacter{state: &WalkState{}}
	character.PressA() // 进入攻击状态
	character.PressB() // 进入防御状态
	character.PressA() // 进入行走状态
	character.PressB() // 进入防御状态
}
