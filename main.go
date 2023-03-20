package main

import (
	"crypto/tls"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	tunnel_ip := os.Getenv("TUNNEL_IP")
	if tunnel_ip == "" {
		panic("Error: failed to get env TUNNEL_IP")
	}

	port := os.Getenv("PORT")
	if port == "" {
		panic("Error: failed to get env PORT")
	}

	upstream, err := url.Parse(tunnel_ip)
	if err != nil {
		panic("Error: failed to parse TUNNEL_IP")
	}

	proxy := httputil.NewSingleHostReverseProxy(upstream)
	proxy.Transport = &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		ForceAttemptHTTP2: true,
		IdleConnTimeout:   300 * 1000000000,
	}

	gin.SetMode(gin.ReleaseMode)
	proxy_server := gin.Default()

	proxy_server.Use(func(client *gin.Context) {
		proxy.ServeHTTP(client.Writer, client.Request)
	})

	if proxy_server.Run(":"+port) != nil {
		panic("Error: failed to start proxy.")
	}
}
