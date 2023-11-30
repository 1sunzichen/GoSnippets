package main

import (
	"errors"
	"strconv"
)

type WeightBalance struct {
	curIndex int
	rss      []*WeightNode
	rsw      []int
}
type WeightNode struct {
	addr            string
	Weight          int //初始化 对节点约定的权重
	currentWeight   int //节点临时权重 每轮都会发生变换
	effectiveWeight int //有效权重 默认与weight 相同 totalWeight=sum(effect)
}

func (r *WeightBalance) Add(Params ...string) error {
	if len(Params) != 2 {
		errors.New("need 2 params")
	}
	parInt, err := strconv.ParseInt(Params[1], 10, 64)
	if err != nil {
		return err
	}
	node := &WeightNode{
		addr:   Params[0],
		Weight: int(parInt),
	}
	node.effectiveWeight = node.Weight
	r.rss = append(r.rss, node)
	return nil
}
func (r *WeightBalance) Next() string {
	var best *WeightNode
	total := 0
	for i := 0; i < len(r.rss); i++ {
		w := r.rss[i]
		total += w.effectiveWeight
		w.currentWeight += w.effectiveWeight
		if w.effectiveWeight < w.Weight {
			w.effectiveWeight++
		}
		if best == nil || w.currentWeight > best.currentWeight {
			best = w
		}
	}
	if best == nil {
		return ""
	}
	best.currentWeight -= total
	return best.addr
}
func (r *WeightBalance) Get(string) (string, error) {
	return r.Next(), nil
}
func main() {

}
