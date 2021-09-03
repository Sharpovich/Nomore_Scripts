package main

import (
	"fmt"
	"log"
	"time"

	"gitlab.com/sharpovichjob/gother/csver"
	gorutinelessons "gitlab.com/sharpovichjob/gother/gorutineLessons"
	"gitlab.com/sharpovichjob/gother/interfaceslesson"
	"gitlab.com/sharpovichjob/gother/jsonmanager"
	"gitlab.com/sharpovichjob/gother/newbattery"
	"gitlab.com/sharpovichjob/gother/notpublicer"
	"gitlab.com/sharpovichjob/gother/publicer"
	"gitlab.com/sharpovichjob/gother/timetotime"
)

var (
	a string = "1000011110"
)

func main() {
	gorutinelessons.MainerGo9()
	gorutinelessons.MainerGo8()
	gorutinelessons.MainerGo7()
	go gorutinelessons.MainerGo6()
	go gorutinelessons.MainerGo210()
	go gorutinelessons.MainerGo4()
	go gorutinelessons.MainerGo3()
	go gorutinelessons.MainerGo2()
	go gorutinelessons.MainerGo()
	go timetotime.UnixDateConvert()
	go timetotime.Duration()
	go timetotime.TimeinTwo()
	go timetotime.TimEr()
	go jsonmanager.JsonEr()
	go jsonmanager.JsonManager()
	go csver.CsvOuter()
	var batteryForTest newbattery.Stringer = &newbattery.Battery{Charge: a}
	go fmt.Println(batteryForTest.String())
	go log.Println(publicer.NewArray())
	time.Sleep(1 * time.Second)
	go log.Println("Program compile")
	go notpublicer.NonArray()
	go interfaceslesson.Interfaces_lesson()
}
