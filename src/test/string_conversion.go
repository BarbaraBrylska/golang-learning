/*************************************************************************
	> File Name: string_conversion.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Sun 06 Jan 2013 10:38:09 PM CST
 ************************************************************************/
package main
import(
	"fmt"
	"strconv"
)
func main(){
	var orig string = "666"
	var an int
	var newS string
	fmt.Printf("The size of ints is: %d\n",strconv.IntSize)
	an,_ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n",an)
	an += 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n",newS)
}
