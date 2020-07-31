package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func main()  {
	coordinateWithContext()
}
func coordinateWithContext(){
	total:=12
	var num int32
	fmt.Println(num)
	cxt,cancelFunc:=context.WithCancel(context.Background())
	for i:=1;i<=total;i++{
		go addNum(&num,i, func() {
			if atomic.LoadInt32(&num)==int32(total){
				fmt.Println("发送撤销信号")
				cancelFunc()
			}else {
				fmt.Println("没有撤销")
			}
		})
	}
	<-cxt.Done()
	fmt.Println("end")
}

//func coordinateWithGroup(){
//	total:=12
//	stride:=3
//	var num int32
//	fmt.Println("启动了几个goroutine",num)
//	var wg sync.WaitGroup
//	for i:=1;i<=total;i=i+stride{
//		wg.Add(stride)
//		for j:=0;j<stride;j++ {
//			go addNum(&num,i+j,wg.Done)
//			fmt.Println(j)
//		}
//		wg.Wait()
//	}
//	fmt.Println("end")
//}

func addNum(numP *int32,id int,deferFunc func())  {
	defer func() {
		deferFunc()
	}()
	for i:=0;;i++{
		currNum:=atomic.LoadInt32(numP)
		newNum:=currNum+1
		time.Sleep(time.Millisecond * 200)
		if atomic.CompareAndSwapInt32(numP,currNum,newNum){
			fmt.Printf("The number: %d [%d-%d]\n", newNum, id, i)
			break
		}else {

		}
	}
}
