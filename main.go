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
		panic("Error: env TUNNEL_IP not defined")
	}

	port := os.Getenv("PORT")
	if port == "" {
		panic("Error: env PORT not defined")
	}

	upstream, err := url.Parse(tunnel_ip)
	if err != nil {
		panic("Error: failed to parse tunnel ip")
	}

	proxy := httputil.NewSingleHostReverseProxy(upstream)
	proxy.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	gin.SetMode(gin.ReleaseMode)
	api := gin.Default()

	api.Use(func(client *gin.Context) {
		proxy.ServeHTTP(client.Writer, client.Request)
	})

	if api.Run(":"+port) != nil {
		panic("Error: failed to initialize api. Port:" + port)
	}
}
