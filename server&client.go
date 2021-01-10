package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type custcars struct {
	vehicleID   int
	vehicleName string
	plateNumber string
	DateTime    string
	Color       string
	ProductDate string
}

var vehicles []custcars

func addVehicle(custs custcars) custcars {
	vehicles = append(vehicles, custs)
	return custs
}

func rndm(min int, max int) int {
	n := rand.Intn(max-min) + min
	return n
}

func plateNO() string {
	letter := [5]string{"a", "b", "c", "d", "f"}
	l := rndm(0, 5)
	first := rndm(10, 100)
	second := rndm(100, 1000)
	third := rndm(10, 100)
	i := strconv.Itoa(first)
	j := strconv.Itoa(second)
	f := strconv.Itoa(third)
	plate := i + " " + letter[l] + " " + j + " ir " + f
	return plate
}

var month = [...]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "10"}
var year = [...]string{"2020", "2019", "2018", "2017", "2016", "2015", "2014", "2013"}

func Prdate() string {
	n := rndm(0, 12)
	my := len(year)
	m := rndm(0, my)
	x := month[n]
	y := year[m]
	out := x + "/" + y
	return out
}

var clrs = [...]string{"red", "blue", "black", "gray", "white", "yellow", "orange"}

func clr() string {
	l := len(clrs)
	n := rndm(0, l)
	return clrs[n]
}

var Car = [...]string{"peraid", "peykan", "samand", "pejuet", "rana", "saina", "L90", "benz", "bmw"}

func Cars() string {
	l := len(Car)
	n := rndm(0, l)
	return Car[n]
}

func servers(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/vehicles" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "only GET method accepted", http.StatusNotFound)
		return
	}
	body, _ := ioutil.ReadAll(r.Body)
	st := string(body)
	to, err3 := strconv.Atoi(st)
	if err3 == nil {
		if to < len(vehicles) {
			finallist := vehicles[:to]
			fmt.Fprint(w, finallist)
			res := strconv.Itoa(len(finallist))
			fmt.Fprint(w, "\n lenght of list:"+res)
		}
		if to >= len(vehicles) {
			fmt.Fprint(w, vehicles)
			res := strconv.Itoa(len(vehicles))
			fmt.Fprint(w, "\n lenght of list:"+res)
		}
	}
	if err3 != nil {
		fmt.Fprint(w, "please input Number")
	}
}

func main() {
	min := 10
	max := 100
	n := rndm(min, max)
	for i := 0; i < n; i++ {
		x := custcars{i,
			Cars(),
			plateNO(),
			time.Stamp,
			clr(),
			Prdate(),
		}
		addVehicle(x)
	}
	http.HandleFunc("/vehicles", servers)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
