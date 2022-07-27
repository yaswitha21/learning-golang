package main

import "fmt"
func main(){
//calling a function with '()'
	simpleFunction()
// f := simpleFunction  //using as variable but not calling it , here we have assigned this simplefunction to variable f ,
// and calling it with variable f 
//f()
withParams("yash",123)
}
//defining a simple function with no return types
func simpleFunction(){
	fmt.Println("simple function")
	 
}
//defining a function with parameters

func withParams(p1 string,p2 int ){

	fmt.Println("inside withParams ::", p1,p2)


}