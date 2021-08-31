package main

import "fmt"

type EventListener interface {
	OnFire()
}

type Event interface {
	Register(EventListener)
}

type Mail struct {
	listener EventListener
}

func (m *Mail) Register(listener EventListener) {
	m.listener = listener
}

func (m *Mail) OnRecv() {
	m.listener.OnFire()
}

type Alarm struct {
}

func (a *Alarm) OnFire() {
	fmt.Println("알람~! 알람~!")
}

func main() {
	var mail = &Mail{}
	var listener EventListener = &Alarm{}

	mail.Register(listener)
	mail.OnRecv()
}
