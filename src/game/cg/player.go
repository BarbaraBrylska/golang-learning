/*************************************************************************
	> File Name: player.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Sat 03 Nov 2012 03:15:18 PM CST
 ************************************************************************/
package cg
import "fmt"

type Player struct{
	Name string "name"
	Level int "level"
	Exp int "exp"
	Room int "room"

	mq chan *Message //message queue
}

func NewPlayer() *Player{
	m := make(chan *Message,1024)
	player := &Player{"",0,0,0,m}

	go func(p *Player){
		for{
			msg := <-p.mq
			fmt.Println(p.Name,"received message:",msg.Content)
		}	
	}(player)

	return player
}
