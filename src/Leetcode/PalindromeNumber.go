package main

import "fmt"

func palindrome(x int) bool{
	res := 0
	ori :=x
	for x > 0 {
		temp := x % 10
		res = res * 10 + temp
		x = x/10
	}
	if (res == ori) {
		return true
	} else {
		return false
	}
}

func main(){
	x := 121
	res := palindrome(x)
	fmt.Println(res)
}