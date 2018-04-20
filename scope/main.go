package main

import "fmt"

func calc(y int)int{
	return y+10
}
var x int = 42
func varibales() func()int{
	b := 1
	return func()int{
		b++
		return b
	}
}

func main()  {

	incr := varibales()
	fmt.Println(incr())
	fmt.Println(incr())

}






