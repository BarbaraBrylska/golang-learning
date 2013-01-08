/*************************************************************************
	> File Name: server.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Sat 03 Nov 2012 08:02:00 PM CST
 ************************************************************************/
package main
import (
	"fmt"
	"net"
	"os"
	"bufio"
)

const help = `
1)exec <command> <args> --execute a command on server
2)quit|exit --close the connection
`

func main(){
	//listening the port
	ln,err := net.Listen("tcp",":5555")
	fmt.Println("server listen at port 5555")
	if err != nil{
		checkError(err)	
	}
	//get the connect
	for{
		conn,err := ln.Accept()
		fmt.Println("Server accept a connection")
		if err != nil{
			fmt.Fprintf(os.Stderr,"Accept Fatal Error: %s",err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

func checkError(err error){
	if err != nil{
		fmt.Fprintf(os.Stderr,"Fatal error: %s",err.Error())
		os.Exit(1)
	}
}



//handle read and write
func handleConnection(conn net.Conn){

	//handle read
	go func(conn net.Conn){
		reader := bufio.NewReader(os.Stdin)	
		for{
			fmt.Print(">>")
			line,_,_ := reader.ReadLine()
			conn.Write([]byte(line))
		}
	}(conn)
	//handle write
	go func(conn net.Conn){
		bs := make([]byte,1024)	
		for{
			l,err := conn.Read(bs)
			if err != nil{
				fmt.Fprintf(os.Stderr,"Fatal error: %s",err.Error())
				conn.Close()
			}
			fmt.Println("<<",string(bs[:l]))
		}
	}(conn)

}
