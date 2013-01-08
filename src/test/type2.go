/*************************************************************************
	> File Name: type2.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Mon 07 Jan 2013 09:33:01 PM CST
 ************************************************************************/
package main
import "fmt"
type integer int
func main(){
	var a int = 34
	var b integer = integer(a)
	fmt.Println(a,b)
}
