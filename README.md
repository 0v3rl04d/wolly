# wolly
A simple and really working WOL module fo Go

## Import module
```
go get github.com/0v3rl04d/wolly
```
## Usage Example

```go
package main

import (
    "log"
    "github.com/0v3rl04d/wolly/wol"
)

func main() {
    mp, err := wol.CreateMagicPacket("FF:FF:FF:FF:FF:FF")
    if err != nil {
        log.Panic(err)
    }
    err = wol.SendMagic(mp, "192.168.1.255", 9)
    if err != nil {
	    log.Panic(err)
    }
}
```