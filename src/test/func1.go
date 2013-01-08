/*************************************************************************
	> File Name: func1.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Mon 07 Jan 2013 10:07:57 PM CST
 ************************************************************************/
package main
import "fmt"

type Log struct{
	msg string
}

type Customer struct{
	Name string
	log *Log
}

func (l *Log)Add(s string){
	l.msg += "\n"+s
}

func (l *Log)String() string{
	return l.msg
}

func (c *Customer)Log() *Log{
	return c.log
}

func main(){
	c := &Customer{"Barak Obam",&Log{"1 - Yes we can!"}}
	fmt.Println(c)
	c.Log().Add("2 - After me the world will be a better place!")
	fmt.Println(c.log)
	fmt.Println(c.Log())
}
