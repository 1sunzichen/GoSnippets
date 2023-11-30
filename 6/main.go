package main

import (
	"fmt"
	"sync"
)

// sync.Once{}
type Person struct {
	level int
}

func (p *Person) AddLevel() {

	p.level++
	fmt.Printf("Level:%d", p.level)
}
func main() {
	p := &Person{level: 0}
	once := sync.Once{}
	w := sync.WaitGroup{}
	w.Add(100)
	for i := 0; i < 100; i++ {

		go once.Do(p.AddLevel)
		w.Done()
	}
	w.Wait()
}
