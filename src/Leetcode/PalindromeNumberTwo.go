package main

import (
	"strconv"
	"fmt"
)

func PalindromeNumberTwo(x int) bool{
	if x < 0{
		return false
	}

	s := strconv.Itoa(x)

	for i,j := 0, len(s) - 1; i < j;i,j = i+i,j-1{
		if s[i] != s[j]{
			return false
		}
	}

	return true
}

func main() {
	x := 121
	res := PalindromeNumberTwo(x)
	fmt.Println(res)
}