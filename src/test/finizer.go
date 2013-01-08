/*************************************************************************
	> File Name: finizer.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Mon 07 Jan 2013 10:32:19 PM CST
 ************************************************************************/
package main
import(
	"fmt"
	"runtime"
)

type T struct{
	A string
	B string
	C string
}

func (t *T)Members() string{

	return t.A + t.B + t.C
}

func testGc(){
	
	t := &T{"A_","B_","C_"}
	fmt.Println(t.Members())
	runtime.SetFinalizer(t,func(t *T){
		fmt.Println("Gb collection.. remove t from memory")
	})
}

func main(){
	testGc()
}
