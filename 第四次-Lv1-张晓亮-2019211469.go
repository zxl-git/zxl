package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var a string
	m1:=make(map[int]time.Time)
	//fmt.Println("输入result得到结果")
	for i := 1; i <= 1000; i++ {
		fmt.Scanf("%s", &a)
		if (a!="result"){
			fmt.Println("Input ok!")
			if b,err :=strconv.Atoi(a);err == nil {
				m1[i] = time.Unix(int64(b), 0)
			}
		}
		if(a=="result"){
			break
		}
	}
	fmt.Println("the result are :")
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}