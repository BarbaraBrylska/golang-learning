/*************************************************************************
	> File Name: method3.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Mon 07 Jan 2013 10:01:02 PM CST
 ************************************************************************/
package main
import (
	"fmt"
	"math"
)

type Point struct{
	x,y float64
	name string
}

func (p *Point)Abs() float64{
	return math.Sqrt(p.x*p.x + p.y*p.y)
}
func (p *Point)Name() string{
	return p.name
}
type NamedPoint struct{
	Point
	name string
}

func main(){
	n := &NamedPoint{Point{3,4,"Base_name"},"Pythagoras"}
	fmt.Println(n.Abs())
	fmt.Println(n.Name())
}
