# Fast and Lightweight HTTP-Proxy
This project utilizes **Gin** to tunnel connections **man-in-the-middle** style.

- Every incoming request are proxied and response will return to the **INITIATOR**.
- One direction only. *(Direction = **TUNNEL_IP**)*.

*May be useful for developing on local machine, as you can tunnel HTTPS to HTTP connections.*

<img src="https://github.com/pedrobartolini/golang_tunnel/blob/main/readme/diagram.png">

### Environment variables:
```env
TUNNEL_IP = Where proxy will forward requests to
PORT = Port proxy will listen to
```

## Instructions
Build project
```batch
go build
```
Run build
```batch
./main
```
**Remember to set your environment variables before running**

## Note for HTTPS connections
If you're using web services like GCP, AWS, Azure, Render, etc, HTTPS setup may be done automatically.

**If you're not using an webserver with https redirection, you'll need to implement it manually into this project.**
