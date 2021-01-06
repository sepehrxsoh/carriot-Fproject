package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type customer struct{
	vehicleID int
	vehicleName string
	plateNumber string
	DateTime string
	Color string
	ProductDate string
}

func CreatVehicle() {
	vehicleList :=[]customer{}
	min := 10
	max := 100
	n := random(min , max)
	for i:= 0;i < n ; i++ {
		cust := new(customer)
		cust.vehicleID = i
		cust.vehicleName = cars()
		cust.plateNumber = plateNumber()
		cust.DateTime = time.Stamp
		cust.Color = color()
		cust.ProductDate = Pdate()
		vehicleList = append(vehicleList,*cust)
	}
	fmt.Println(vehicleList)
	}

func random(min int , max int) int {
	n := rand.Intn(max - min) + min
	return n
}

func plateNumber() string{
	letter := [5]string{"a","b","c","d","f"}
	l := random(0,5)
	first :=random(10,100)
	second := random(100,1000)
	third := random(10,100)
	i := strconv.Itoa(first)
	j := strconv.Itoa(second)
	f := strconv.Itoa(third)
	plate := i +" "+ letter[l]+ " " +  j + " ir "+f
	return plate
}

var Month = [...]string{"01","02","03","04","05","06","07","08","09","10","11","10"}
var Year = [...]string{"2020","2019","2018","2017","2016","2015","2014","2013"}

func Pdate() string {
	n := random(0,12)
	my := len(Year)
	m := random(0,my)
	x := Month[n]
	y := Year[m]
	out := x + "/"+y
	return out
}

var colors = [...]string{"red","blue","black","gray","white","yellow","orange"}

func color() string {
	l := len(colors)
	n := random(0,l)
	return colors[n]
}
var car =[...]string{"peraid","peykan","samand","pejuet","rana","saina","L90","benz","bmw"}

func cars() string {
	l := len(car)
	n := random(0,l)
	return car[n]
}


func main(){
	CreatVehicle()
}