package main

import(
	"fmt"
	"hashtable"
	"strconv"
)
func main(){
	var ht hashtable.HashTable
	ht.InitHashtable(4,1)
	fmt.Println()
	fmt.Println("------------------------------------------------------------------------------------")
	for i:=0;i<10;i++{
		ht.Put(i,strconv.Itoa(i)+" hello")
		fmt.Println()
		fmt.Println("哈希表结点个数：",ht.GetCount())
		fmt.Println("哈希表的容量：",ht.GetCapacity())
		fmt.Println("哈希表阈值：",ht.GetThreshold())
		fmt.Println("哈希表tablesize：",ht.GetTablesize())
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("哈希表结点个数为：",ht.GetCount())
	fmt.Println()
	for i:=0;i<10;i++{
		fmt.Println(i,"---->",ht.Get(i))
	}
	fmt.Println()
	fmt.Println("key为20的value：",ht.Get(20))
	fmt.Println("key为8的value：",ht.Get(8))
	
	
	
}