package common

import "fmt"

type Doer interface {
	Do() error
}

type myDoer struct{}

func NewMyDoer() myDoer {
	return myDoer{}
}

func (d myDoer) Do() error {
	fmt.Println("I am doing something...")
	return nil
}
