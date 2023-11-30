package main

import (
	"testing"
)

type Person1122 struct {
	level int
}

func (p *Person1122) Add() {
	p.level++
}
func TestMain(t *testing.T) {
	p := &Person1122{
		level: 0,
	}
	// once := sync.Once{}
	for i := 0; i < 100; i++ {
		// go once.Do(p.Add)
		go p.Add()
	}
	want := 1
	got := p.level
	if want != got {

		t.Errorf("got :%d,want :%d", got, want)
	} else {

		t.Logf("got :%d,want :%d", got, want)
	}

}
