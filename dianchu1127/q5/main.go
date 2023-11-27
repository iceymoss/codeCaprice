package main

import "fmt"

/*
【软件工程】编程题：比萨利造过程为准备、烘、切片、装盒，假设有一个比萨店。比萨店能制造厦门和媒圳两种风格的比萨，
每种风格比萨又分芝士、素食、始果、香肠四种类世，不同风格间种型的比萨也不相间，用合适的设计模式实现这样的比萨店
相关功能类。
*/

// Pizza 接口定义
type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

// PizzaStore 抽象比萨店
type PizzaStore interface {
	OrderPizza(pizzaType string) Pizza
}

// ---------- 美式风格 ----------

// NYStylePizzaStore 美式比萨店
type NYStylePizzaStore struct{}

type NYStyleClamPizza struct {
}

type NYStylePepperoniPizza struct {
}

// OrderPizza 订单比萨
func (s *NYStylePizzaStore) OrderPizza(pizzaType string) Pizza {
	var pizza Pizza
	switch pizzaType {
	case "cheese":
		pizza = &NYStyleCheesePizza{}
	case "veggie":
		pizza = &NYStyleVeggiePizza{}
	case "clam":
		pizza = &NYStyleClamPizza{}
	case "pepperoni":
		pizza = &NYStylePepperoniPizza{}
	default:
		fmt.Println("不支持的比萨类型")
		return nil
	}

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}

// 具体产品实现
type NYStyleCheesePizza struct{}

func (p *NYStyleCheesePizza) Prepare() {
	fmt.Println("准备纽约风格芝士比萨")
}

func (p *NYStyleCheesePizza) Bake() {
	fmt.Println("烘烤纽约风格芝士比萨")
}

func (p *NYStyleCheesePizza) Cut() {
	fmt.Println("切片纽约风格芝士比萨")
}

func (p *NYStyleCheesePizza) Box() {
	fmt.Println("装盒纽约风格芝士比萨")
}

// 具体产品实现
type NYStyleVeggiePizza struct {
	// 实现 Veggie 比萨的具体内容
}

func (p *NYStyleVeggiePizza) Prepare() {
	fmt.Println("准备纽约风格素食比萨")
}

func (p *NYStyleVeggiePizza) Bake() {
	fmt.Println("烘烤纽约风格素食比萨")
}

func (p *NYStyleVeggiePizza) Cut() {
	fmt.Println("切片纽约风格素食比萨")
}

func (p *NYStyleVeggiePizza) Box() {
	fmt.Println("装盒纽约风格素食比萨")
}

// ---------- 意大利风格 ----------

// ChicagoStylePizzaStore 芝加哥比萨店
type ChicagoStylePizzaStore struct{}

// OrderPizza 订单比萨
func (s *ChicagoStylePizzaStore) OrderPizza(pizzaType string) Pizza {
	var pizza Pizza
	switch pizzaType {
	case "cheese":
		pizza = &ChicagoStyleCheesePizza{}
	case "veggie":
		pizza = &ChicagoStyleVeggiePizza{}
	case "clam":
		pizza = &ChicagoStyleClamPizza{}
	case "pepperoni":
		pizza = &ChicagoStylePepperoniPizza{}
	default:
		fmt.Println("不支持的比萨类型")
		return nil
	}

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}

// 具体产品实现
type ChicagoStyleCheesePizza struct{}

func (p *ChicagoStyleCheesePizza) Prepare() {
	fmt.Println("准备芝加哥风格芝士比萨")
}

func (p *ChicagoStyleCheesePizza) Bake() {
	fmt.Println("烘烤芝加哥风格芝士比萨")
}

func (p *ChicagoStyleCheesePizza) Cut() {
	fmt.Println("切片芝加哥风格芝士比萨")
}

func (p *ChicagoStyleCheesePizza) Box() {
	fmt.Println("装盒芝加哥风格芝士比萨")
}

// 具体产品实现
type ChicagoStyleVeggiePizza struct {
	// 实现 Veggie 比萨的具体内容
}

func (p *ChicagoStyleVeggiePizza) Prepare() {
	fmt.Println("准备芝加哥风格素食比萨")
}

func (p *ChicagoStyleVeggiePizza) Bake() {
	fmt.Println("烘烤芝加哥风格素食比萨")
}

func (p *ChicagoStyleVeggiePizza) Cut() {
	fmt.Println("切片芝加哥风格素食比萨")
}

func (p *ChicagoStyleVeggiePizza) Box() {
	fmt.Println("装盒芝加哥风格素食比萨")
}

// 具体产品实现
type ChicagoStyleClamPizza struct{}

func (p *ChicagoStyleClamPizza) Prepare() {
	fmt.Println("准备芝加哥风格蛤蜊比萨")
}

func (p *ChicagoStyleClamPizza) Bake() {
	fmt.Println("烘烤芝加哥风格蛤蜊比萨")
}

func (p *ChicagoStyleClamPizza) Cut() {
	fmt.Println("切片芝加哥风格蛤蜊比萨")
}

func (p *ChicagoStyleClamPizza) Box() {
	fmt.Println("装盒芝加哥风格蛤蜊比萨")
}

// 具体产品实现
type ChicagoStylePepperoniPizza struct{}

func (p *ChicagoStylePepperoniPizza) Prepare() {
	fmt.Println("准备芝加哥风格香肠比")
}
func (p *ChicagoStylePepperoniPizza) Bake() {
	fmt.Println("烘烤芝加哥风格蛤蜊比萨")
}

func (p *ChicagoStylePepperoniPizza) Cut() {
	fmt.Println("切片芝加哥风格蛤蜊比萨")
}

func (p *ChicagoStylePepperoniPizza) Box() {
	fmt.Println("装盒芝加哥风格蛤蜊比萨")
}
