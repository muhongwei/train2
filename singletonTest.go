package main

import (
	"fmt"
	"singleton"
)
func main(){
	obj := singleton.Instance()

	obj.Setname("mumu")
	fmt.Println("obj.Say():")
	obj.Say()

	obj2 :=singleton.Instance()
	fmt.Println("obj2.Say():")
	obj2.Say()

}