package linknodelist
import "fmt"
//定义错误常量
const(
    ERROR = -66666666
)
//定义节点
type LinkNode struct{
	value int
	next *LinkNode
}
//定义链表
type LinkNodeList struct{
	first *LinkNode    //头指针，指向该链表的第一个元素
    last *LinkNode     //尾指针，指向该链表的最后一个元素
    curr *LinkNode
	size int      //链表的长度
}

//初始化
func (ls *LinkNodeList)InitList(){
	var head LinkNode=LinkNode{value:0,next:nil}
    ls.first=&head
    ls.last=&head
    ls.curr=&head
    ls.size=0	
}
//显示链表所以内容
func (ls *LinkNodeList)Show_list(){	
    pointer:=ls.first
    if(pointer.next==nil){
        fmt.Println("链表为空")
    }
	for pointer.next != nil{
		pointer=pointer.next
		fmt.Print(pointer.value,"   ")
	}
	
}
//将当前位置移动到链表首
func (ls *LinkNodeList)MoveToStart(){
    ls.curr=ls.first
}
//将当前位置移动到链表尾
func (ls *LinkNodeList)MoveToEnd(){
    ls.curr=ls.last
}
//将当前位置移动到指定位置
func(ls *LinkNodeList)MoveToPos(pos int){
    if((pos>=0)&&(pos<=ls.size)){
        ls.curr=ls.first
        for i:=0;i<pos;i++{
            ls.curr=ls.curr.next
        }
    }
}
//将当前位置向前移动一个位置
func (ls *LinkNodeList) Prev(){
    if(ls.first==ls.curr){
        return
    }
    temp:=ls.first
    for temp.next!=ls.curr{
        temp=temp.next
    }
    ls.curr=temp
}
//返回当前位置
func (ls *LinkNodeList)CurrPos() int{
    temp := ls.first
    var i int
    for i=0;ls.curr != temp;i++{
        temp=temp.next
    }
    return i
}
//将当前位置向后移动一个位置
func (ls *LinkNodeList) Next(){
    if(ls.curr != ls.last){
        ls.curr=ls.curr.next
    }
}
//尾插入
func (ls *LinkNodeList)Push_back(data int){
    pointer:=ls.last
    
	var ld LinkNode
	ld.value=data
	pointer.next=&ld

    ls.last=&ld

    ls.size++
}
//头插入
func (ls *LinkNodeList)Push_front(data int){
	pointer:=ls.first

	var ld LinkNode
	ld.value=data
	ld.next=pointer.next

	pointer.next=&ld
    ls.first.next=&ld
    
    ls.size++
}
//尾删除
func (ls *LinkNodeList)Pop_back() bool{
    if(ls.size==0){
        fmt.Println("链表已空")
        return false
    }
    pointer := ls.first
    for pointer.next != ls.last{
        pointer=pointer.next
    }
    pointer.next=nil
    ls.last=pointer
    ls.size--
    return true
}
//头删除
func(ls *LinkNodeList) Pop_front() bool{
    if(ls.size==0){
        fmt.Println("链表已空")
        return false
    }
    pointer:=ls.first
    pointer.next=pointer.next.next
    if(ls.size==1){
        ls.last=ls.first
    }
    ls.size--

    return true
}
//查找元素
func (ls *LinkNodeList)Find(data int) *LinkNode{
    pointer:=ls.first.next
    for pointer!=nil && pointer.value != data{
        pointer=pointer.next
    }
	return pointer
}
//修改元素
func (ls *LinkNodeList)Modify(data int,newValue int) bool{
    pointer:=ls.Find(data)
    if(pointer!=nil){
        pointer.value=newValue
        return true
    }else{
        return false
    }
}
//删除元素
func (ls *LinkNodeList)Delete_val() int{
    // pointer:=ls.Find(data)
    // if (pointer!=nil){
    //     if(pointer==ls.first){
    //         ls.Pop_front()
    //     }else if(pointer==ls.last){
    //         ls.Pop_back()
    //     }else{
    //         pointer.value=pointer.next.value
    //         pointer.next=pointer.next.next
    //     }
    //     ls.size--
    //     fmt.Printf("%d已经被删除",data)
    // }else{
    //     fmt.Printf("要删除的元素%d不存在",data)
    // }
    
    if(ls.curr.next == nil){
        fmt.Println("没有元素")
        return ERROR
    }
    it:=ls.curr.next.value
    if(ls.curr.next == ls.last){
        ls.last=ls.curr
    }
    ls.curr.next=ls.curr.next.next
    ls.size--
    return it    
}
//清除链表
func (ls *LinkNodeList)Clear(){
    pointer:=ls.first.next
    for pointer!=nil{
        ls.first.next=pointer.next
        pointer=ls.first.next
    }
    ls.last=ls.first
    ls.size=0
}
//链表长度
func (ls *LinkNodeList)Length() int{
	return ls.size
}
//指定位置插入
func (ls *LinkNodeList)Insert(data int,index int){
    if(index<=0||index>ls.size+1){
        fmt.Println("输入位置错误")
    }else{
        pointer:=ls.first
        var linknode LinkNode
        linknode.value=data
        for i:=1;i<index;i++{
            pointer=pointer.next
            ls.curr=pointer
        }
        linknode.next=pointer.next
        pointer.next=&linknode
        if(ls.curr==ls.last){
            ls.last=ls.curr.next
        }
        ls.size++
    }

}
//获取当前位置元素值
func (ls *LinkNodeList)GetValue()int{
    if(ls.curr.next==nil){
        return ERROR
    }else{
        return ls.curr.next.value
    }
    

}

