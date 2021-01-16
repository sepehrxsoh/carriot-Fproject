package main

import (
	"context"
	"encoding/json"
	//"fmt"
	"github.com/sepehrxsoh/carriot-Fproject/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type Customer struct {
	vehicleID   int
	vehicleName string
	plateNumber string
	DateTime    string
	Color       string
	ProductDate string
}

var VehicleList []Customer

func CreatVehicle(cust Customer) Customer {
	VehicleList = append(VehicleList, cust)
	return cust
}

func Random(min int, max int) int {
	n := rand.Intn(max-min) + min
	return n
}

func plateNumber() string {
	letter := [5]string{"a", "b", "c", "d", "f"}
	l := Random(0, 5)
	first := Random(10, 100)
	second := Random(100, 1000)
	third := Random(10, 100)
	i := strconv.Itoa(first)
	j := strconv.Itoa(second)
	f := strconv.Itoa(third)
	plate := i + " " + letter[l] + " " + j + " ir " + f
	return plate
}

func Pdate() string {
	n := Random(0, 12)
	my := len(Year)
	m := Random(0, my)
	x := Month[n]
	y := Year[m]
	out := x + "/" + y
	return out
}

var Month = [...]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "10"}
var Year = [...]string{"2020", "2019", "2018", "2017", "2016", "2015", "2014", "2013"}

var colors = [...]string{"red", "blue", "black", "gray", "white", "yellow", "orange"}

func Color() string {
	l := len(colors)
	n := Random(0, l)
	return colors[n]
}

var car = [...]string{"peraid", "peykan", "samand", "pejuet", "rana", "saina", "L90", "benz", "bmw"}

func Cars() string {
	l := len(car)
	n := Random(0, l)
	return car[n]
}

func List() {
	min := 10
	max := 100
	n := Random(min, max)
	for i := 0; i < n; i++ {
		x := Customer{i,
			Cars(),
			plateNumber(),
			time.Stamp,
			Color(),
			Pdate(),
		}
		CreatVehicle(x)
	}
}

/*func (s *server) Makelist(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	pay := request.GetPayload()
	x := int64(len(VehicleList))
	if pay > x {
		pay = x
	}
	list := VehicleList[:pay]
	result, _ := json.Marshal(list)
	return &proto.Response{Result: string(result)}, nil
}
*/
type server struct{}

func (s *server) MakeList(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	//panic("implement me")
	pay := request.GetPayload()
	x := int64(len(VehicleList))
	if pay > x {
		pay = x
	}
	list := VehicleList[:pay]
	result, _ := json.Marshal(list)
	return &proto.Response{Result: string(result)}, nil
}

func main() {
	List()
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterListCustomersServer(srv, &server{})
	reflection.Register(srv)
	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}
