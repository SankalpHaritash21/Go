//? var array_name = [length]datatype{values}; --> // here length is defined

package main

import (
	"fmt"
)

func printArrays() {
	var arr1 = [3]int{1,2,3}
	arr2 := [5]int{4,5,6,7,8}
	  
	fmt.Println(arr1)
    fmt.Println(arr2)
}

func Initialization() {
	arr1 := [5]int{} //not initialized
	arr2 := [5]int{1,2} //partially initialized
	arr3 := [5]int{1,2,3,4,5} //fully initialized
  
	fmt.Println(len(arr1))
	fmt.Println(arr2)
	fmt.Println(arr3)
  }