package main

import "fmt"

const PI = 3.14

var name, sex, title string
var age int
var i int

const (
	lable = 1
)

type Books struct {
	title string
	author string
	subject string
	book_id string
}

type Phone interface {
	call()
}

type NokiaPhone struct {

}

type testInt func(int) bool

func main(){
	slice := []int{1,2,3,4,5,6,7}
	fmt.Println("slice=", slice)

	odd := filter(slice, isOdd)
	fmt.Println("odd is =", odd)

	even := filter(slice, isEven)
	fmt.Println("even is = ", even)
}

func isOdd(integer int) bool{
	if integer%2 == 0{
		return false
	}
	return true
}

func isEven(integer int) bool{
	if integer %2 ==0{
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _,value := range slice{
		if f(value){
			result = append(result, value)
		}
	}
	return result
}

func (nokiaphone NokiaPhone) call()  {
	fmt.Println("Nokia")
}

type Iphone struct {

}

func (iphone Iphone) call()  {
	fmt.Println("iphone")
}




func myFunc()(int){
	sum := 1

	for sum < 10{
		sum += sum
	}
	return sum
}

func addone(a *int) int{
	*a = *a+1
	return *a
}

func swap(x,y string) (string,string){
	return y,x
}

func getAverage(arr []int, size int) float32 {
	var i,sum int
	var avg float32

	for i=0;i<size;i++{
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}

func Factorial(n uint64) (result uint64)  {
	if (n > 0){
		result = n * Factorial(n-1)
		return result
	}

	return 1
}

func fibonacci(n int) int  {
	if n < 2{
		return n
	}

	return fibonacci(n-2) + fibonacci(n-1)
}

