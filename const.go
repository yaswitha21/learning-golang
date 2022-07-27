package main

import "fmt"

const(
/* DATE_FORMAT, localConst we have used in package level 
and any variable or const starting with lowercase letter meaning that scope is only to that file.
and with uppercase scope is with the whole package . anyone outside the package can get access(it will be visible outside the package )
*/

	DATE_FORMAT = "mm/dd/yyyy" // its exported as starting with Uppercase letter

	localConst = 34// its only inside his package as it is starting with lowercase letter meaning that scope is only to that file
)

func main(){

	const name string = "yash"
	const name1 ="sri" // we cant put := here because it will create variable but not constant 
	const i = 200 // even though we dont use name1 it will just give warning
	                 // but wont give compilation error
fmt.Println(name,i)
fmt.Println(DATE_FORMAT)
fmt.Println(localConst)
} 

