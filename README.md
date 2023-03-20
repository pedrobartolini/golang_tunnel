# Fast and lightweight HTTP-Proxy written in Golang
This project utilizes Gin to tunnel connections man-in-the-middle style.
Both requests and responses are proxied.<br>
May be useful for developing on local machine, as you can tunnel HTTPS to HTTP connections.<br>
<img src="https://github.com/pedrobartolini/golang_tunnel/blob/main/git/dada.png" width="500">

## Instructions for HTTP connections
Just compile and run.<br>

## Instructions for HTTPS connections
Depending on the web service you're using, SSL setup may done automatically. (Google Cloud, AWS, Azure, Render, etc)<br>
If your web service doesn't do this, or if you're not using a web service, you'll need to implement SSL manually into this project.

### Environment variables:
**TUNNEL_IP** =  *Ip that proxy will forward requests to<br>*
**PORT** = *Port that proxy will listen to*
