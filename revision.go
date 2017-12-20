package main

import "fmt"

func main(){

num := 0

name := "surya"

for(num < 100){

	fmt.Println(num)
	num = num + 1
}

fmt.Println("Hi " ,name)

var age int

fmt.Println("Whats your age?")
fmt.Scanf("%d",&age)

if(age < 18){

	fmt.Println("You cant drive")
}else{
	fmt.Println("You can drive")
}


}
