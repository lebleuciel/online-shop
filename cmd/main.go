package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	backendapi "github.com/lebleuciel/online-shop/internal/backend"
)

func main() {
	fmt.Println("onlineshop is running")
	backendServer, _ := setupHttpServers()
	go runServer(backendServer, 8080, "backend Server")
	shutDown := make(chan os.Signal, 0)
	signal.Notify(shutDown, syscall.SIGHUP, syscall.SIGINT)
	<-shutDown
	fmt.Printf("shutting down the server")
}
func runServer(server *http.Server, port int, serverName string) {
	fmt.Printf("%s starting listening on port %d", serverName, port)
	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		_ = fmt.Errorf("could not create listener on %s ,error : %s", serverName, err.Error())
	}
	err = server.Serve(ln)
	if err != nil {
		_ = fmt.Errorf("error on serving %s, error : %s", serverName, err.Error())
	}
}
func setupHttpServers() (*http.Server, *http.Server) {
	fmt.Println("setting up servers")
	backendServerHandler, err := backendapi.NewBackendServer()
	if err != nil {
		_ = fmt.Errorf("error on setting backend server, error : %s", err.Error())
	}
	backendServer := &http.Server{
		Addr:    ":8080",
		Handler: backendServerHandler,
	}
	return backendServer, nil
}
