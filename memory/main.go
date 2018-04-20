package main

import "fmt"

const yard float64 = 1.09361

func main(){
	var meters float64
	fmt.Print("enter")
	fmt.Scan(&meters);
	yards := meters * yard
	fmt.Println(meters," meters is: ",yards, " yards")

}
