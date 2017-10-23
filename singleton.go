package singleton

import (
	"fmt"
	"sync"
)

var _instance *object
var once sync.Once
//var lock *sync.Mutex = &sync.Mutex{}

type object struct {
    name string
}

func Instance() *object {
//单例模型
//    if _instance == nil {
//        _instance = new(object)
//    }
//    
    //双重锁，解决并发问题
    //   if _instance == nil{
	// 	  lock.Lock()
	// 	  defer lock.Unlock()
	// 	  if _instance == nil{
	// 		  _instance = new(object)
	// 	  }
	//   }
	
	//sync.Once的Do方法中的函数只会执行一次
	  once.Do(func(){
		  _instance = new (object)
	  })
	  return _instance
  
}

func (p *object) Setname(name string) {
    p.name = name
}

func (p *object) Say() {
    fmt.Println("my name is ",p.name)
}
