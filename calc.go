package main

import "fmt"

func main(){
  fmt.Println("Enter the first number")
  var  num1 int
  fmt.Scanln(&num1)
  fmt.Println("Enter the second number")
  var num2 int
  fmt.Scanln(&num2)
  fmt.Printf("add\n")
  fmt.Printf("subration\n")
  fmt.Printf("multi\n")
  fmt.Printf("division\n")
  fmt.Printf("select any one operation\n")
  var s string
  fmt.Scanln(&s)


  switch s{
  case "add" :
	add  := num1 + num2
	fmt.Println("the result is :\n" , add)
  case "subration" :
    add  := num1 - num2
	fmt.Println("the result is :\n" , add)
  case "multi" :
	add  := num1 * num2
	fmt.Println("the result is :\n" , add)
  case "division" :
	add  := num1 / num2
	fmt.Println("the result is :\n" , add)  
  default :
   fmt.Println("invalid")
  }
  
}

