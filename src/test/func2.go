/*************************************************************************
	> File Name: func2.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Mon 07 Jan 2013 10:13:39 PM CST
 ************************************************************************/
package main
import(
	"fmt"
)
type Log struct{
	msg string
}

type Customer struct{
	Name string
	Log
}

func (l *Log)Add(s string){
	l.msg += "\n" + s
}

func (c *Customer)String() string{
	return c.Name + "\nLog:" + fmt.Sprintln(c.Log)
}

func (l *Log)String() string{
	return l.msg
}

func main(){
	c := &Customer{"Barak Obama",Log{"1 - Yes we can!"}}
	c.Add("2 - After me the world will be a better place!")
	fmt.Println(c)
}
