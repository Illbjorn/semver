# Overview

`semver` is a basic zero-allocation semantic version parsing package.

# Example

```go
package main

import "github.com/illbjorn/semver"
import "fmt"

func main() {
  input := "v1.22.333-rc0"

  ver, err := semver.Parse(input)
  assert(err == nil)
  assert(ver.Prefix == "v")
  assert(ver.Major == 1)
  assert(ver.Minor == 22)
  assert(ver.Patch == 333)
  assert(ver.More == "-rc0")
}

func assert(condition bool) {
  if !condition {
    panic(0)
  }
}
```

TODO: More examples.

# Benchmarks

```log
goos: linux
goarch: amd64
pkg: github.com/illbjorn/semver
cpu: AMD Ryzen 9 5900X 12-Core Processor
BenchmarkParse
BenchmarkParse-24       46279713                26.25 ns/op            0 B/op          0 allocs/op
```

```log
goos: windows
goarch: amd64
pkg: github.com/illbjorn/semver
cpu: AMD Ryzen 9 5900X 12-Core Processor
BenchmarkParse
BenchmarkParse-24       48085399                24.87 ns/op            0 B/op          0 allocs/op
```
