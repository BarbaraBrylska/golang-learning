/*************************************************************************
	> File Name: time_base.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Sun 06 Jan 2013 10:44:27 PM CST
 ************************************************************************/
package main
import(
	"fmt"
	"time"
)
var week time.Duration
func main(){
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%04d.\n",t.Day(),t.Month(),t.Year())
	//06.01.2013
	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())
	week = 60 * 60 * 24 * 7 * 1e9 
	week_from_now := t.Add(week)
	fmt.Println(week_from_now)
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2006 15:04"))
	s := t.Format("20060102")
	fmt.Println(t,"=>",s)
}
