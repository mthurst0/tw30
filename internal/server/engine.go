package server

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ipAllowList(al map[string]bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		if !al[clientIP] {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status":  http.StatusForbidden,
				"message": "Permission denied",
			})
			zap.L().Info("Permission denied", zap.String("clientIP", clientIP))
			return
		}
	}
}

func engine() (*gin.Engine, error) {
	r := gin.Default()
	ips := os.Getenv("TW30_IP_ALLOW_LIST")
	if ips != "" {
		al := make(map[string]bool)
		al["127.0.0.1"] = true
		for _, ip := range strings.Split(ips, ",") {
			ip = strings.TrimSpace(ip)
			al[ip] = true
		}
		r.Use(ipAllowList(al))
	}
	if err := r.SetTrustedProxies(nil); err != nil {
		return nil, err
	}
	return r, nil
}

type ListenAddr struct {
	UseTLS bool
	Addr   string
}

func parseListenAddr(s string) (ListenAddr, error) {
	if strings.HasPrefix(s, "https://") {
		return ListenAddr{UseTLS: true, Addr: s[8:]}, nil
	}
	if strings.HasPrefix(s, "http://") {
		return ListenAddr{UseTLS: false, Addr: s[7:]}, nil
	}
	return ListenAddr{UseTLS: false, Addr: s}, nil
}
