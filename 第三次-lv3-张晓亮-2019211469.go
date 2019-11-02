/*package main

import (
	"fmt"
	"time"
)

func main(){
	go animation()

	for i:=2;i<=10000;i++ {
		time.Sleep(1)
		flag := 1
		for j := 2; j <= i/2; j++ {
			if (i%j == 0) {
				flag = 0
			}
		}
		if (flag != 0) {
			fmt.Println(i)
		}
	}
}

func animation(){
	for{
		for _, r := range "-.~"{
			fmt.Printf("\r%c", r)
		}
	}
}
*/

package main

import (
	"fmt"
	"time"
)

func sushu(){
	for i:=2;i<=10000;i++ {

		flag := 1
		for j := 2; j <= i/2; j++ {
			if (i%j == 0) {
				flag = 0
			}
		}
		if (flag != 0) {
			fmt.Println(i)
		}
	}
}

func main(){
	for a :=1;a<=10000;a++{
		time.Sleep(1)
		go sushu()
	}
}
