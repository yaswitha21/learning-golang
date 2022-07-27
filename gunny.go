package main

import "fmt"

func main(){
	myvar1, myvar2 := 39,45
	myvar3,myvar2 := 45,100
	fmt.Printf("the values of myvar1 and myvar2 are : %d,%d \n", myvar1,myvar2)
	fmt.Printf("the values of myvar3 and myvar2 are : %d,%d \n", myvar3,myvar2)
}
