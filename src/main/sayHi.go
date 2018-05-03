package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money	float32
}

func (h Human) sayHi(){
	fmt.Printf("Hi, I am %s ,phone num %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("la la ", lyrics)
}

func (e Employee) sayHi() {
	fmt.Printf("Hi, I am %s, I work in %s, salary %s",e.name, e.company, e.money)
}

type Men interface {
	sayHi()
	Sing(lyrics string)
}

func main(){
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	var i Men

	i = mike
	fmt.Println("This is mike, a Student:")
	i.sayHi()
	i.Sing("lolipop")

	i = tom
	fmt.Println("This is tom")
	i.sayHi()
	i.Sing("born this way")

	fmt.Println("use slice")

	x := make([]Men,3)
	x[0],x[1],x[2] = paul,sam,mike

	for _,value := range x{
		value.sayHi()
	}
}
