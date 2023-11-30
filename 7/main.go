package main

import (
	"fmt"

	"dario.cat/mergo"
)

type Person struct {
	Hobby []string
	Name  string
	Age   int
}

type Duck interface {
	Quack()
}

func (p *Person) Quack() {
	fmt.Println("a")
}

type Cat struct{}

func (c *Cat) Quack() {
	fmt.Println("喵")
}
func main() {
	p1 := &Person{Hobby: []string{"打篮球", "羽毛球"}, Age: 90}
	m1 := make(map[string]interface{})
	err := mergo.Map(&m1, p1)
	if err != nil {
		fmt.Println("我错了 我错了")
	}
	m2 := map[string]interface{}{
		"name": "11",
		"Hobby2": []string{
			"1111", "111",
		},
	}
	p2 := &Person{}
	err = mergo.Map(p2, m2)

	if err != nil {
		fmt.Println("我错了 我错了")
	}
	fmt.Printf("m1: %+v \n", m1)
	fmt.Printf("p2: %+v \n", p2)
	cat := &Cat{}
	cat.Quack()
	p2.Quack()
}
