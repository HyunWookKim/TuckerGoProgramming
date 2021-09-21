package main

import "fmt"

type Flyweight struct {
	Somedata   string
	isDisposed bool
}

func (f *Flyweight) Reuse() {
	f.isDisposed = false
}

func (f *Flyweight) Dispose() {
	f.isDisposed = true
}

func (f *Flyweight) IsDisposed() bool {
	return f.isDisposed
}

type FlyweightFactory struct {
	pool     []*Flyweight
	AllocCnt int
}

func (fac *FlyweightFactory) Create() *Flyweight {
	var obj *Flyweight
	if len(fac.pool) > 0 {
		obj, fac.pool = fac.pool[len(fac.pool)-1], fac.pool[:len(fac.pool)-1]
		obj.Reuse()
	} else {
		obj = &Flyweight{}
		fac.AllocCnt++
	}
	return obj
}

func (fac *FlyweightFactory) Dispose(obj *Flyweight) {
	obj.Dispose()
	fac.pool = append(fac.pool, obj)
}

func NewFlyweightFactory(initSize int) *FlyweightFactory {
	return &FlyweightFactory{pool: make([]*Flyweight, 0, initSize)}
}

func main() {
	fac := NewFlyweightFactory(1000)
	for i := 0; i < 1000; i++ {
		obj := fac.Create()
		obj.Somedata = "Somedata"
		fac.Dispose(obj)
	}
	fmt.Println("AllocCnt:", fac.AllocCnt)
}
