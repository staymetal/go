package main

import "fmt"

func main(){

number := 10

var o *int

o =  &number

fmt.Println("The value of number is ",number)
fmt.Println("The address of number is ",o)
fmt.Println("Adding 10 to number")
*o = *o + 10
fmt.Println("Value of number is ",*o)


















}
