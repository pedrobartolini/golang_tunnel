# Fast and Lightweight HTTP-Proxy
This project utilizes **Gin** to tunnel connections **man-in-the-middle** style.<br>

- Every incoming request are proxied and response will return to the **INITIATOR**.
- One direction only. *(Direction = **TUNNEL_IP**)*.

*May be useful for developing on local machine, as you can tunnel HTTPS to HTTP connections.*
<img src="https://github.com/pedrobartolini/golang_tunnel/blob/main/git/diagram.png">

### Environment variables:
**TUNNEL_IP** =  *Server that proxy will forward requests to<br>*
**PORT** = *Port that proxy will listen to*

## Instructions for HTTP connections
Build project
```batch
go build
```

Run build.
```batch
./main
```

## Instructions for HTTPS connections
Depending on the web service you're using, SSL setup may done automatically. (Google Cloud, AWS, Azure, Render, etc)<br>
If your web service doesn't do this or if you're not using a web service, you'll need to implement SSL manually into this project.

