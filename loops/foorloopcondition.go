package main

import "fmt"

func main(){
	i :=0
	for {
		i++
		if i==5 {
			fmt.Println("next")
			continue
		}

		fmt.Println(i)
		if i==7 {
			fmt.Println("finished")
			break


	}


	}



}