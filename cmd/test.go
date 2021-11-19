package main

import (
	"fmt"
	fm "fmt"
	ft "fmt"
	"log"
	md "math/rand"
	ra "math/rand"
	rd "math/rand"
	"os"
	sv "strconv"
	"time"
)

func main() {
	//var test = 42.0
	//fmt.Println(reflect.TypeOf(test))
	var test = 10 / 3.0
	ft.Printf("%010.5f\n", test)

	fm.Printf("%-20v %-5v %-12v %-6v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("================================================")
	spaceLine := []string{
		"Space Adventures",
		"SpaceX",
		"Virgin Galactic",
		"Shen Zhou",
	}
	tripType := []string{
		"One-way",
		"Round-trip",
	}

	rd.Seed(time.Now().UnixNano())
	num, err := sv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < num; i++ {
		line := spaceLine[ra.Intn(4)]
		speed := rd.Intn(15) + 16
		tripIndex := rd.Intn(2) + 1
		trip := tripType[tripIndex-1]
		days := (62100000 / speed / 60 / 60 / 24) * tripIndex
		price := (md.Intn(15) + 36) * tripIndex

		fmt.Printf("%-20v %-5v %-12v $%4v %4v\n", line, days, trip, price, speed)
	}

}
