/*************************************************************************
	> File Name: struct.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Mon 07 Jan 2013 07:17:46 PM CST
 ************************************************************************/
package main
import(
	"fmt"
)

type T struct{
	A int
	B string
}

func main(){
	var t T
	//t.A = 12
	//t.B = "Hello World"
	fmt.Printf("%v\n",t)
}
