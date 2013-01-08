/*************************************************************************
	> File Name: pointer.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Sun 06 Jan 2013 10:58:34 PM CST
 ************************************************************************/
package main
import(
	"fmt"
)
func main(){
	var a int = 32
	var b *int = &a
	fmt.Printf("%d 's address is %p",*b,b)
}
