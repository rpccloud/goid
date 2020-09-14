# goid
Get current goroutine id 

## Usage
```go
package main

import (
  "fmt"
  "github.com/rpccloud/goid"
)

func main() {
  fmt.Println("Current Goroutine ID:", goid.GoRoutineId())
}
```

## Benchmark
```bash
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/rpccloud/goid
BenchmarkGoRoutineId-12         1000000000               0.413 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/rpccloud/goid        1.040s
```

