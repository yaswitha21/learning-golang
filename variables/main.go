package main

import "fmt"

var(

	t = 23
	i =1 /* i is defined in global scope and unused ,
	 we can declare same variable in two different scopes and it wont show error
	but when we declare same variable more than once it will show error

	FIRST it will check local scope and then it will go for global scope.
*/
)
func main(){
	/* var is the keyword
	foo is the variable name
    int is the built-in type
	45 is the value we initialized 
	*/

	var foo int = 45// declaration with initialization
	var bar int  
	// declaration without initialization   
var deb = 23 
 var x,y int= 2,7 // multi value declaration with intialization
 var i,j int//i is defined in main scope  
  // multi value declaration without intialization

 var name string = "yash" // name :=  yash both are equal, this is called short variable declaration and 
                          // this is used in loops that we goona use in particular block

						  a,b := 2,"test"
t = 56 // we have updated the value of t , which is declared in global scope.
	fmt.Println(foo,deb)
	fmt.Printf("the type of deb is %T \n",deb)
	fmt.Printf("the value of bar is %v\n", bar)
	fmt.Println(x,y,i,j)
	fmt.Println(name)
	fmt.Println(a,b)
	fmt.Println(t)


	// go is strictly typed language and you cant change the type of variables once it is declared.
	//i.e, here b is declared as string and you cant change it to int or something else
	

	
}