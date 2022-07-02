package main

import "fmt"

type Strategy interface {
	Execute()
}

type strategyA struct {
}

func (a strategyA) Execute() {
	fmt.Println("A Plan Executed")
}

func NewStrategyA() Strategy {
	return &strategyA{}
}

type strategyB struct {
}

func (b strategyB) Execute() {
	fmt.Println("B Plan Executed")
}

func NewStrategyB() Strategy {
	return &strategyB{}
}

type Context struct {
	strategy Strategy
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) Execute() {
	c.strategy.Execute()
}

func main() {
	s_A := NewStrategyA()
	c := NewContext()
	c.SetStrategy(s_A)
	c.Execute()

	s_B := NewStrategyB()
	c.SetStrategy(s_B)
	c.Execute()
}
