[![Build Status](https://travis-ci.org/artnez/structconfig.svg?branch=master)](http://travis-ci.org/artnez/structconfig)

```go
package main

import (
  "log"
  "os"

  "github.com/artnez/structconfig"
)

type Config struct {
  ServerHost string  `env:"SERVER_HOST"`
  ServerPort int     `env:"SERVER_PORT"`
  Timeout    float64 `env:"TIMEOUT"`
}

func main() {
  config := &Config{
    ServerHost: "localhost",
    ServerPort: 5000,
  }
  structconfig.FromEnviron(config, os.Environ())
  log.Printf("[config] %s", structconfig.String(config))
}
```
