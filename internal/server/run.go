package server

import (
	"fmt"
	"os"

	"github.com/mthurst0/tw30/internal/appcore"

	"github.com/mthurst0/tw30/internal/webapp"

	"github.com/gin-gonic/gin"
)

func Run() error {
	baseRoutes, err := engine()
	if err != nil {
		return err
	}

	a := appcore.Must()
	webapp.AddRoutes(a, baseRoutes)

	listenAddr, err := parseListenAddr(os.Getenv("TW30_LISTEN_ADDR"))
	if err != nil {
		return err
	}
	if listenAddr.UseTLS {
		return runTLS(listenAddr.Addr, baseRoutes)
	} else {
		return runLocal(listenAddr.Addr, baseRoutes)
	}
}

func runTLS(addr string, r *gin.Engine) error {
	certPath := os.Getenv("TW30_CERT_PATH")
	if certPath == "" {
		return fmt.Errorf("TW30_CERT_PATH must be set")
	}
	keyPath := os.Getenv("TW30_PRIVATE_KEY_PATH")
	if keyPath == "" {
		return fmt.Errorf("TW30_PRIVATE_KEY_PATH must be set")
	}
	if err := r.RunTLS(addr, certPath, keyPath); err != nil {
		return err
	}
	return nil
}

func runLocal(addr string, r *gin.Engine) error {
	if err := r.Run(addr); err != nil {
		return err
	}
	return nil
}
