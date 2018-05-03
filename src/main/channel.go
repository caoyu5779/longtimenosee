package main

import "fmt"

func sum(a []int, c chan int){
	total := 0
	for _,v := range a{
		total += v
	}

	c <- total
}

func main()  {
	//a := []int{8,9,9,1,-1,0}
	//
	//c := make(chan int)
	//
	//go sum(a[:len(a) /2], c)
	//go sum(a[len(a)/2 :], c)
	//
	//x,y := <-c,<-c
	//fmt.Println(x, y)

	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0;i<10;i++{
			fmt.Println(<-c)
		}
		quit<-0
	}()

	fibonacci(c,quit)
}

func fibonacci(c,quit chan int)  {
	x,y := 1,1
	for {
		select {
			case c<-x:
				x,y = y,x+y
			case <-quit:
				fmt.Println("quit")
				return
		}
	}
}


