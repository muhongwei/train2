package hashtable

import (
	"fmt"
)
//定义哈希结点
type HashNode struct{
	key int
	value string
	next *HashNode
}
//定义哈希链表
type HashNodeList struct{
	head *HashNode
	tail *HashNode
	curr *HashNode
	size int
	hash int
}
//定义哈希表
type HashTable struct{
	table []HashNodeList     //储存链表的数组
	count int                //哈希表内存储的哈希结点个数
	threshold int            //用于判断是否需要扩充哈希表的容量
	loadFactor float32       //加载因子
	modCount int 
	tableSize int            //哈希表中储存链表的数组大小
	capacity int             //哈希表的容量
}
//返回哈希表储存的结点个数
func (ht *HashTable) GetCount() int{
	return ht.count
}
//返回哈希表储存的阈值
func (ht *HashTable) GetThreshold() int{
	return ht.threshold
}
//返回哈希表的容量
func(ht *HashTable) GetCapacity() int{
	return ht.capacity
}
//返回哈希表的tablesize
func (ht *HashTable) GetTablesize() int{
	return ht.tableSize
}
//初始化哈希表
func (ht *HashTable) InitHashtable(initCapacity int,tableSize int){
	if(initCapacity<0){
		fmt.Println("illegal capacity")
	}
	if(initCapacity==0){
		initCapacity=1
	}
	ht.loadFactor = 0.75
	ht.count = 0
	ht.capacity = initCapacity
	ht.tableSize = tableSize
	ht.table = make([]HashNodeList,ht.tableSize)
	for i:=0;i<ht.tableSize;i++{
		var hl HashNodeList
		hl.Init(i)
		ht.table[i]=hl
		
	}	
	ht.threshold = int(ht.loadFactor * float32(ht.capacity))
}
//向哈希表里面添加元素
func (ht *HashTable) Put(key int,value string) string{
	var hd HashNode
	hd.key=key
	hd.value=value
	hash:=Hash(key,ht)

	var hl HashNodeList
	hl=ht.table[hash]
	
	pointer:=hl.head
	
	for pointer.next != nil{
		pointer=pointer.next
		if (hl.hash==hash) && (pointer.key==key){
			oldValue:=pointer.value
			pointer.value=value
			return oldValue
		}
	}
	ht.AddToHashtable(hd,hash)
	return "nil"
}
//返回键对应的值
func (ht *HashTable) Get(key int) string{
	hash:=Hash(key,ht)
	var hl HashNodeList
	hl=ht.table[hash]
	
	pointer:=hl.head
	
	for pointer.next != nil{
		pointer=pointer.next
		if (hl.hash==hash) && (pointer.key==key){
			return pointer.value
		}
	}
	return "没有key对应得结点"

}
//散列方法
func Hash(key int,ht *HashTable) int{
	return key & (ht.tableSize-1)
}
//哈希表的扩容
func (ht *HashTable) Rehash(){
	
	oldCapacity := ht.capacity
	newCapacity := oldCapacity*2
	oldTablesize := ht.tableSize
	newTableszie := oldTablesize*2

	var newHt HashTable
	newHt.InitHashtable(newCapacity,newTableszie)
	for i:=0;i<oldTablesize;i++{
		var hl HashNodeList
		hl = ht.table[i]

		pointer:=hl.head
		
		for pointer.next != nil{
			pointer=pointer.next
			key := pointer.key
			value := pointer.value
			newHt.Put(key,value)			
		}

	}
	*ht=newHt
}
func (hl *HashNodeList) Init(i int){
	var head HashNode=HashNode{key:0,value:"head",next:nil}
	hl.head=&head
	hl.tail=&head
	hl.curr=&head
	hl.hash=i
	hl.size=0
}
//将哈希结点添加到对应的哈希链表中
func (ht *HashTable)AddToHashtable(hd HashNode,index int){
	
	hl:=ht.table[index]
	pointer:=hl.head
	
	var hdtemp HashNode
	hdtemp.key=hd.key
	hdtemp.value=hd.value

	hdtemp.next=pointer.next

	pointer.next=&hdtemp
	hl.head.next=&hdtemp

	hl.size++
	ht.count++

	if(ht.count>=ht.threshold){
		fmt.Println("rehash")
		ht.Rehash()
	}
}