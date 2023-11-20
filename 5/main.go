package main

//负载均衡算法 随机均衡
import (
	"errors"
	"fmt"
	"math/rand"
)

type RandomBalance struct {
	curIndex int
	rss      []string
}

func (r *RandomBalance) Add(Params ...string) error {
	if len(Params) == 0 {
		return errors.New("Params is nil")
	}
	addr := Params[0]
	r.rss = append(r.rss, addr)
	return nil
}
func (r *RandomBalance) Next() string {
	if len(r.rss) == 0 {
		return ""
	}
	r.curIndex = rand.Intn(len(r.rss))
	return r.rss[r.curIndex]
}
func (r *RandomBalance) Get() (string, error) {
	return r.Next(), nil
}
func main() {
	rb := &RandomBalance{rss: []string{"192.168.1.1", "192.168.1.2", "192.168.1.3"}}
	ip, _ := rb.Get()
	fmt.Printf("IP:%s", ip)
}
