package main
import(
    "linknodelist"
    "fmt"
)

func main(){
    fmt.Println("==================================================================================================")
    fmt.Println("==================================================================================================")
    fmt.Println("-------------实现方法如下：----------------")
    fmt.Println("-------------MoveToStart()----------------")
    fmt.Println("-------------MoveToEnd()------------------")
    fmt.Println("-------------MoveToPos()------------------")
    fmt.Println("-------------Prev()-----------------------")
    fmt.Println("-------------Next()-----------------------")
    fmt.Println("-------------Currpos()--------------------")
    fmt.Println("-------------Initlist()-------------------")
    fmt.Println("-------------Showlist()-------------------")
    fmt.Println("-------------Push_back(data int)----------")
    fmt.Println("-------------Push_front(data int)---------")
    fmt.Println("-------------Pop_back()-------------------")
    fmt.Println("-------------Pop_front()------------------")
    fmt.Println("-------------Find(data int)---------------")
    fmt.Println("-------------Modify(data int,newValue int)")
    fmt.Println("-------------Delete_val()-----------------")
    fmt.Println("-------------Clear()----------------------")
    fmt.Println("-------------Length()---------------------")
    fmt.Println("-------------Insert(data int,index int)---")
    fmt.Println("-------------GetValue()-------------------")
    fmt.Println("==================================================================================================")
    fmt.Println("==================================================================================================")
    fmt.Println()
    var ls linknodelist.LinkNodeList

    fmt.Println("----------------InitList()--------   -----")
    
    ls.InitList()	
    for i:=0;i<10;i++{
        if(i%2==0){
            ls.Push_back(i)
        }else{
            ls.Push_front(i)
        }
	}
    fmt.Println("----------------Show_list()----------------")
    fmt.Println()
    ls.Show_list()
    fmt.Println()
    fmt.Println()
    fmt.Println("----------------MoveToFirst()--------------")
    fmt.Println("----------------MoveToEnd()----------------")
    fmt.Println("----------------MoveToPos(pos int)---------")
    fmt.Println("----------------CurrPos--------------------")
    fmt.Println()
   // ls.Show_list()
    fmt.Println("movetofirst:")
    ls.MoveToStart()
    fmt.Println("当前位置为：",ls.CurrPos())
    fmt.Println("movetoend:")
    ls.MoveToEnd()
    fmt.Println("当前位置为：",ls.CurrPos())
    fmt.Println("movetopos(6):")
    ls.MoveToPos(6)
    fmt.Println("当前位置为：",ls.CurrPos())
    fmt.Println()
    fmt.Println()
    fmt.Println("----------------Length()-------------------")
    fmt.Println()
    length:=ls.Length()
    fmt.Println("the size of the linknodelist is:",length)
    fmt.Println()
    fmt.Println("----------insert(data int,pos int)指定插入位-")
    fmt.Println()
    fmt.Println("插入66到为第11个结点")
    ls.Insert(66,11)
    ls.Show_list()
    fmt.Println()
    fmt.Println("插入77到为第1个结点")
    ls.Insert(77,1)
    ls.Show_list()
    fmt.Println()
    fmt.Println("插入88到为第5个结点")
    ls.Insert(88,5)
    ls.Show_list()
    fmt.Println() 
    fmt.Println("插入99到为第20个结点")
    ls.Insert(99,20)
    ls.Show_list()   
    fmt.Println()
    fmt.Println()
    fmt.Println("-----------------Pop_back()-----------------")
    fmt.Println()
    if(ls.Pop_back()){
        ls.Show_list()
    }
    fmt.Println()
    fmt.Println()
    fmt.Println("-----------------Pop_front()----------------")
    fmt.Println()
    if(ls.Pop_front()){
        ls.Show_list()
    }
    fmt.Println()
    fmt.Println()
    fmt.Println("-----------------Delete_val()删除当前位置元素-")
    fmt.Println()
    fmt.Println("删除元素1:")
    var tag bool
    for ls.MoveToStart();ls.CurrPos()<ls.Length();ls.Next(){
        if(ls.GetValue()==1){
            ls.Delete_val()
            tag=true
            break
        }
    }
    if tag{
        fmt.Println("删除成功")
        tag=false
    }else{
        fmt.Println("没有找到要删除的元素")
    }
    fmt.Println()
    ls.Show_list()
    fmt.Println()
    fmt.Println("删除元素20:")
    for ls.MoveToStart();ls.CurrPos()<ls.Length();ls.Next(){
        if(ls.GetValue()==20){
            ls.Delete_val()
            tag=true
            break
        }
    }
    if tag{
        fmt.Println("删除成功")
    }else{
        fmt.Println("没有找到要删除的元素")
    }
    fmt.Println()
    ls.Show_list()
    fmt.Println()
    fmt.Println()
    fmt.Println("-------------Find(data int)：返回指针-------")
    fmt.Println()
    fmt.Println("Find(8):",ls.Find(8))
    fmt.Println("Find(7):",ls.Find(7))
    fmt.Println()
    fmt.Println("--------------Modify(old int,new int) ------")
    fmt.Println()
    fmt.Print("修改2为22:")
    if(ls.Modify(2,22)){
        fmt.Println("修改成功")
    }else{
        fmt.Println("修改失败")
    }
    ls.Show_list()
    fmt.Println()
    fmt.Print("修改10为100:")
    if(ls.Modify(10,100)){
        fmt.Println("修改成功")
    }else{
        fmt.Println("修改失败")
    }
    ls.Show_list()    
    fmt.Println()
    fmt.Println("----------------clear()----------------------")
    fmt.Println()
    ls.Clear()
    ls.Show_list()
    fmt.Println()
    fmt.Println("============================================================================================")
    fmt.Println("============================================================================================")
}
