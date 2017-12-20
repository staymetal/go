package main

import "fmt"

func main(){


var list [4]int

list[0] = 1
list[1] = 2

fmt.Println(list)

for value := range list {
	fmt.Println(value)

}










}
