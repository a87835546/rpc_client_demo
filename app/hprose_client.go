package app

import (
	"fmt"
	"github.com/hprose/hprose-go"
	"log"
	"time"
)

type Person struct {
	Name     string
	Age      int
	Sex      bool
	birthday int64
}
type clientStub struct {
	Hello      func(string) string
	Swap       func(int, int) (int, int)
	Sum        func(...int) (int, error)
	AddTest    func(Person) (Person, error)
	GetPersons func() []Person
}

func init() {
	client := hprose.NewClient("http://127.0.0.1:8080/")
	log.Printf("client --->>> %v\n", client)
	var ro *clientStub
	client.UseService(&ro)
	fmt.Println(ro.Hello("World111"))
	fmt.Println(ro.Swap(1, 2))
	fmt.Println(ro.Sum(1, 2, 3, 4, 5))
	fmt.Println(ro.Sum(1))
	p := Person{
		"zhansan",
		18, false,
		time.Now().UnixNano(),
	}
	if p1, err := ro.AddTest(p); err != nil {
		fmt.Errorf("err -->>> %s", err.Error())
	} else {
		fmt.Printf("person --->>>> %v\n  p -->>> %v \n", p1, p)
	}

	fmt.Printf("persons --->>>> %v\n", ro.GetPersons())
}
