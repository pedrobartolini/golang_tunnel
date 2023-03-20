# Fast and Lightweight HTTP-Proxy
This project utilizes **Gin** to tunnel connections **man-in-the-middle** style.

- Every incoming request are proxied and response will return to the **INITIATOR**.
- One direction only. *(Direction = **TUNNEL_IP**)*.

*May be useful for developing on local machine, as you can tunnel HTTPS to HTTP connections.*

<img src="https://github.com/pedrobartolini/golang_tunnel/blob/main/git/diagram.png">

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

## For HTTPS connections
If you're using web services like GCP, AWS, Azure, Render, etc, HTTPS setup may be done automatically.

**If your web service doesn't do this or if you're not using a web service, you'll need to implement SSL manually into this project.**

Then follow steps in [**Instructions**](#instructions) section.


