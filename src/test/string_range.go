/*************************************************************************
	> File Name: string_range.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Sun 06 Jan 2013 11:27:17 PM CST
 ************************************************************************/
package main
import(
	"fmt"
)
func main(){
	var s string = "chinese: \u65e5\u672c\u8a9e"
	fmt.Println(s)
	for pos,char := range s {
		fmt.Printf("character %c starts at byte position %d\n",char,pos)
	}
	fmt.Println("index int(rune) rune char bytes")
	for index,rune := range s {
		fmt.Printf("%-2d     %d     %U  '%c' % X\n",index,rune,rune,rune,[]byte(string(rune)))
	}
}
