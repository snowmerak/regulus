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
