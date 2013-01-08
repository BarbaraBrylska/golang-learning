/*************************************************************************
	> File Name: type.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Tue 08 Jan 2013 09:57:15 PM CST
 ************************************************************************/
package main
import "fmt"
import "strconv"
type T struct{
	a,b int
}
type IT interface{
	show() string
}

func (t *T)show() string{
	return strconv.Itoa(t.a + t.b)
}

func main(){
	var it IT = new(T)	
	switch type_ := it.(type){
		default:
			fmt.Printf("Type of it is %T, and value is %v",type_,type_)

	}
	var s interface{}= &StringStruct{"hello world"}
	if sv,ok := s.(Stringer); ok {
		fmt.Printf("\ns implements String(): %s\n",sv.String())
	}
}
type StringStruct struct{
	A string
}
func (s *StringStruct)String()string {
	return s.A
}
type Stringer interface {
	String() string
}
