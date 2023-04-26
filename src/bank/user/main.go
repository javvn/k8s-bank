package main

import (
	"context"
	"fmt"
	"jawncorp.com/bank/user/infrastructure"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	envTLSCert = "TLS_CERT"
	envTLSKey  = "TLS_KEY"
)

func main() {

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM)

	router := infrastructure.Router()

	go func() {
		//if os.Getenv(envTLSCert) == "" || os.Getenv(envTLSKey) == "" {
		router.Logger.Fatal(router.Start(":3000"))
		//} else {
		//	router.Logger.Fatal(router.StartTLS(":443",
		//		os.Getenv(envTLSCert), os.Getenv(envTLSKey)))
		//}
	}()

	<-quit
	fmt.Println("Caught SIGTERM, shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := router.Shutdown(ctx); err != nil {
		router.Logger.Fatal("Error:", err)
	}
	fmt.Println("Exited app")
}
