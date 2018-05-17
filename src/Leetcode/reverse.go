package main

import (
	"fmt"
	"math"
)

func reverse(x int) int{
		lable := 1
		if (x < 0){
			lable = -1
		}

		res := 0
		for x > 0{
			temp := x % 10

			res = res * 10 + temp

			x = x/10
		}

		res = lable * res

		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
		return res
}

func main(){
	res := reverse(1234)
	fmt.Println(res)
}