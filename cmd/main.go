package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	fmt.Println("onlineshop is running")
}
func runServer(server *http.Server, port int, serverName string) {
	fmt.Printf("%s starting lstening on port %d", serverName, port)
	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		fmt.Errorf("could not create listener on %s ,error : %s", serverName, err.Error())
	}
	err = server.Serve(ln)
	if err != nil {
		fmt.Errorf("error on serving %s, error : %s", serverName, err.Error())
	}
}