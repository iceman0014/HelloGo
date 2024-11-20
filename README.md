# HelloGo
the repository of go project

just say hello to go

## Install
import code 
```bash
go get github.com/felixliang/HelloGo@latest
```

install cmd
```bash
go install github.com/felixliang/HelloGo@latest
```

## Example
Here is a simple example as follows:

```go
package main

import (
    "fmt"
    "github.com/felixliang/HelloGo/hello"
)
func main() {
    result := hello.SayHello("world")
    fmt.Println(result)
}
```

The output of this program will be:
```
Hello, world!
``` 

## License
[MIT](https://choosealicense.com/licenses/mit/)
