package main

import "fmt"

func main()  {
	for i :=0;i<1000;i++{
		for j :=0;j<1000;j++{
			fmt.Println(i," - ",j)
		}
	}
}