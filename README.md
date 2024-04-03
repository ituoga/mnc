# mnc

micro wrapper around nats

env:

```
NATS_URL=nats://127.0.0.1:4222
```


package usage: 

```golang
package main
import "github.com/ituoga/mnc"

type Response struct {
    Name string `json:"name"`
}

type Request struct {
    Name string `json:"name"`
}

func main() {
    response, err := mnc.Call[Response]("function.on.topic", Request{"hello world"})
    if err != nil {
        panic(err)
    }
    log.Printf("%+#v", response)
}

```

and you could simply test that app with `natscli`

https://github.com/nats-io/natscli

by runing 
```
nats reply --echo "function.on.topic"
```

and `go run yoruapp.main.go`


