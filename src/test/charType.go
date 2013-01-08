/*************************************************************************
	> File Name: charType.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Sun 06 Jan 2013 10:12:41 PM CST
 ************************************************************************/
package main
import(
	"fmt"
)
func main(){
	var ch int = '\u0041'
	var ch2 int = '\u03b2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n",ch,ch2,ch3)//integer
	fmt.Printf("%c - %c - %c\n",ch,ch2,ch3)//character
	fmt.Printf("%X - %X - %X\n",ch,ch2,ch3)//UTF-8 bytes
	fmt.Printf("%U - %U - %U\n",ch,ch2,ch3)//UTF-8 code point
}
