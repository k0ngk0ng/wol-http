# Wake-On-LAN using HTTP

## Installation

```bash
$ git clone https://github.com/k0ngk0ng/wol-http.git
$ cd wol-http
$ go build
$ ./wol-http
```

## Usage

```bash
$ curl 'http://localhost:8000/wol/wake?mac=xx:xx:xx:xx:xx:xx&bcast=255.255.255.255'
```