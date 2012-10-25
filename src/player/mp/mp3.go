/*************************************************************************
	> File Name: mp3.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Thu 25 Oct 2012 10:18:13 PM CST
 ************************************************************************/
package mp
import(
	"fmt"
	"time"
)

type MP3Player struct{
	stat int
	progress int
}

func (p *MP3Player)Play(source string){
	fmt.Println("Playing MP3 music",source)
	p.progress = 0
	for p.progress < 100{
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		p.progress += 10
	}
	fmt.Println("\nFinished playing",source)
}
