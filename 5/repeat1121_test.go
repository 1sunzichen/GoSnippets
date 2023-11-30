package main

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
)

type RandomBalance1122 struct {
	curIndex int
	rss      []string
}

func (r *RandomBalance1122) Add(Params ...string) error {
	if len(Params) == 0 {
		errors.New("Params is not have")
	}
	addr := Params[0]
	r.rss = append(r.rss, addr)
	return nil
}

func (r *RandomBalance1122) Next() string {
	l := len(r.rss)
	if l == 0 {
		return ""
	}
	r.curIndex = rand.Intn(l)
	return r.rss[r.curIndex]
}

func (r *RandomBalance1122) Get() string {
	return r.Next()
}

func Test1121(t *testing.T) {
	rb := &RandomBalance1122{rss: []string{"1", "2", "193.168.1.3"}}
	fmt.Println(rb.Get())
	got := rb.Get()
	nowant := ""
	if got == nowant {
		t.Errorf("got :%s,nowant :%s", got, nowant)
	}
	t.Logf("got :%s,nowant :%s", got, nowant)
}
