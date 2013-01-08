/*************************************************************************
	> File Name: client.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Sat 03 Nov 2012 08:51:40 PM CST
 ************************************************************************/
package main
import(
	"bufio"
	"fmt"
	"strings"
	"net"
	"os"
)

var flag chan int
func main(){
	flag = make(chan int,0)
	//prompt user to insert command
	stdin := bufio.NewReader(os.Stdin)
	for{
		fmt.Println(`
		conn <ip:port>
		`)
		 line,_,_:= stdin.ReadLine()
		 line2 := string(line)
		 args := strings.Split(line2," ")
		 fmt.Println("args = ",args,"length = ",len(args))
		 if len(args) != 2 || args[0] != "conn"{
			 fmt.Println("Unknow command:",args)
			 continue
		 }else{
			tcpAddr,err := net.ResolveTCPAddr("tcp4",args[1])
			if err != nil{
				fmt.Fprintf(os.Stderr,"Fatal resolve error: %s",err.Error())
				continue
			}
			conn,err := net.DialTCP("tcp",nil,tcpAddr)
			if err != nil{
				fmt.Fprintf(os.Stderr,"Fatal connection error: %s",err.Error())
				continue
			}
			go handleConnection(conn)
			//success and break
			break
		 }
	}
	//waiting to disconnect
	<-flag
	fmt.Println("Connection has been closed...")
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
				flag<-1
			}
			fmt.Println("<<",string(bs[:l]))
		}
	}(conn)

}
