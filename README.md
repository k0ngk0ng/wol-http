# Wake-On-LAN using HTTP

## Requirement

`go version >= 1.12`

## Build

```bash
$ git clone https://github.com/k0ngk0ng/wol-http.git
$ cd wol-http
$ go build
```

## Start

`wol-http [-p port]`

Note: the port number is 8000 by default.

## Usage

```bash
$ curl 'http://localhost:8000/wol/wake?mac=xx:xx:xx:xx:xx:xx&bcast=255.255.255.255'
```
### query string 
- mac: the mac address
- bcast: the broadcast address, default is 255.255.255.255