/*************************************************************************
	> File Name: server.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Sat 03 Nov 2012 11:44:13 AM CST
 ************************************************************************/
package ipc
import(
	"encoding/json"
	"fmt"
)

type Request struct{
	Method string "method"
	Params string "params"
}

type Response struct{
	Code string "code"
	Body string "body"
}

type Server interface{
	Name() string
	Handle(method,params string) *Response
}

type IpcServer struct{
	Server
}

func NewIpcServer(server Server) *IpcServer{
	return &IpcServer{server}
}

func (server *IpcServer)Connect() chan string{
	session := make(chan string,0)

	go func(c chan string){
		for{
			request := <-c

			if request == "CLOSE"{
				break
			}

			var req Request
			err := json.Unmarshal([]byte(request),&req)
			if err != nil{
				fmt.Println("Invalid request format:",request)
			}
			resp := server.Handle(req.Method,req.Params)
			b,err := json.Marshal(resp)
			c <- string(b)
		}	
		fmt.Println("Session Closed...")
	}(session)

	fmt.Println("A new session has been created successfully.")
	return session
}
