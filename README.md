# regulus

Regulus is a simple, lightweight web api for obtaining client information.

## Installation

### As a binary

```bash
go install github.com/snowmerak/regulus@latest
```

### As a library

```bash
go get github.com/snowmerak/regulus
```

```go
package main

import "github.com/snowmerak/regulus/regulus"

func main() {
	server := regulus.New(":8080")

	if err := server.Start(); err != nil {
		panic(err)
	}
}
```

## Usage

```bash
regulus --port <port>
```

- `--port` - The port to run the server on. Default is `:8080`.

## Endpoints

### GET /ip

Returns the client's IP address.  
Supports the `X-Forwarded-For` and `X-ProxyUser-Ip`, `X-Real-Ip` headers.

Supported formats:
- str: /ip?format=str

```
127.0.0.1
```

- json: /ip?format=json

```
{
  "ip": "127.0.0.1"
}
```

- xml: /ip?format=xml

```
<ip>127.0.0.1</ip>
```

- yaml: /ip?format=yaml

```
ip: 127.0.0.1
```

- toml: /ip?format=toml

```
ip = 127.0.0.1
```
