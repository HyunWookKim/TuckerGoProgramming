package main

type Database interface {
	Get()
	Set()
}

type CDatabase struct {
}

func (c CDatabase) GetData() {

}

func (c CDatabase) SetDta() {

}

type Wrapper struct {
	cdb CDatabase
}

func (w Wrapper) Get() {
	w.cdb.GetData()
}

func (w Wrapper) Set() {
	w.cdb.SetDta()
}

func main() {
	var cdatabase CDatabase
	var database Database
	database = Wrapper{cdatabase}
	database.Get()
	database.Set()
}
