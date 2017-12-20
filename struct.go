package main

import "fmt"

func main(){


mycar := Car{1000,"Mercedes",220}

fmt.Printf("My Car is a %s which has %d litre fuel and runs on a top speed of %d km/hr",mycar.brand,mycar.fuel,mycar.speed)




}

type Car struct {

	fuel int
	brand  string
	speed  int

}
