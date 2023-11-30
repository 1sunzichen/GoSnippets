package main

import (
	"errors"
	"fmt"
)

type RandomBalance1123 struct {
	curIndex int
	rss      []string
}

func (r *RandomBalance1123) Add(Params ...string) error {
	if len(Params) == 0 {
		return errors.New("params  len 1 at least")
	}
	addr := Params[0]
	r.rss = append(r.rss, addr)
	return nil
}

func (r *RandomBalance1123) Next() string {
	l := len(r.rss)
	if l == 0 {
		return ""
	}
	if r.curIndex >= l {
		r.curIndex = 0
	}
	r.curIndex = (r.curIndex + 1) % l
	return r.rss[r.curIndex]
}

func (r *RandomBalance1123) Get() string {
	return r.Next()
}

func main() {
	addrs := &RandomBalance1123{rss: []string{"1", "2", "3"}}
	addrs.Add("4")
	fmt.Println(addrs.rss)
}
