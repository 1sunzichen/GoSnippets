package main

import (
	"fmt"
	"testing"

	"dario.cat/mergo"
)

type Person1122 struct {
	Hobby []string
	Age   int
}

func TestMain(t *testing.T) {
	p := &Person1122{
		Hobby: []string{"你好", "北京"},
		Age:   0,
	}
	m := make(map[string]interface{})
	mergo.Map(&m, p)
	fmt.Println(m, "m")
	//声明结构体 首字母大写   转换后为小写开头
	got := m["age"]
	want := 0

	if got != want {
		t.Errorf("got %d,want %d", got, want)
	} else {

		t.Logf("got %d,want %d", got, want)

	}
}
