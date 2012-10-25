/*************************************************************************
	> File Name: play.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Thu 25 Oct 2012 10:14:42 PM CST
 ************************************************************************/
package mp
import "fmt"
type Player interface{
	Play(source string)
}

func Play(source,mtype string){
	var p Player
	switch mtype{
	case "MP3":
		p = &MP3Player{}
	case "WAV":
		p = &WAVPlayer{}
	default:
		fmt.Println("Unsupported music type",mtype)
	}
	p.Play(source)
}
