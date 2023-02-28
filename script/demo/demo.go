package demo

import "fmt"

type DemoScript struct {
}

var DemoS DemoScript

func (s *DemoScript) China() func() {
	return func() {
		fmt.Println("Hello China")
	}
}

func (s *DemoScript) World() {
	fmt.Println("Hello World")
}

func (s *DemoScript) Run() {
	s.World()
}
