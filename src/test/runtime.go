/*************************************************************************
	> File Name: runtime.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Sun 06 Jan 2013 11:16:57 PM CST
 ************************************************************************/
package main
import(
	"fmt"
	"runtime"
)
func main(){
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	var str string
	var str2 string = ""
	fmt.Println(str == "",str2 == "",len(str))
}
