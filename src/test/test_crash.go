/*************************************************************************
	> File Name: test_crash.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Sun 06 Jan 2013 11:05:53 PM CST
 ************************************************************************/
package main
import(
	"fmt"
	"errors"
	"reflect"
)
func main(){
	e := errors.New("Error Content")
	fmt.Println(reflect.TypeOf(e))
	i := 32
	fmt.Println(reflect.TypeOf(i))
	var i2 *int = &i
	fmt.Println(reflect.TypeOf(i2))

}
