# gsm

Enhanced version of Go's [sync.Map](https://pkg.go.dev/sync#Map) with generic types and iterators.

## Features:
1. Wrapping of all sync.Map methods with methods having generic parameter and return types.
1. Iterator methods `Range()`, `Keys()`, and `Values()`
1. [Unit tests](https://github.com/aaronriekenberg/gsm/blob/main/gsm_test.go) of every method
1. [Documentation](https://pkg.go.dev/github.com/aaronriekenberg/gsm) including [runnable example](https://pkg.go.dev/github.com/aaronriekenberg/gsm#example-GenericSyncMap)
1. Very fast, on Apple M2 in [parallel](https://pkg.go.dev/testing#B.RunParallel) benchmarks. 
    1. Adds 0.1 nanoseconds overhead vs `sync.Map`
    1. 60x faster than `sync.RWMutex.RLock` and builtin go map
    1. See [benchmarks for raw output](#benchmarks)

## Example:

```go
package main

import (
	"fmt"
	"slices"

	"github.com/aaronriekenberg/gsm"
)

func main() {

	type name string

	type person struct {
		title string
		age   int
	}

	nameToPerson := gsm.GenericSyncMap[name, person]{}

	nameToPerson.Store("alice", person{title: "engineer", age: 25})
	nameToPerson.Store("bob", person{title: "manager", age: 35})

	value, ok := nameToPerson.Load("alice")
	fmt.Printf("alice value = %+v ok = %v\n", value, ok)
	fmt.Printf("alice title = %q\n", value.title)

	value, ok = nameToPerson.Load("bob")
	fmt.Printf("bob value = %+v ok = %v\n", value, ok)
	fmt.Printf("bob age = %v\n", value.age)

	swapped := nameToPerson.CompareAndSwap(
		"alice",
		person{title: "engineer", age: 25},
		person{title: "manager", age: 25},
	)
	value, ok = nameToPerson.Load("alice")
	fmt.Printf("swapped = %v value = %+v ok = %v\n", swapped, value, ok)

	keys := slices.Sorted(nameToPerson.Keys())
	fmt.Printf("keys = %v", keys)
}
```

output:

```
alice value = {title:engineer age:25} ok = true
alice title = "engineer"
bob value = {title:manager age:35} ok = true
bob age = 35
swapped = true value = {title:manager age:25} ok = true
keys = [alice bob]
```

## Benchmarks:

```
$ go test -bench=. -benchmem ./...

goos: darwin
goarch: arm64
pkg: github.com/aaronriekenberg/gsm
cpu: Apple M2
BenchmarkSyncMapParallelLoad-8          	769333302	         1.428 ns/op	       0 B/op	       0 allocs/op
BenchmarkGenericSyncMapParallelLoad-8   	758809900	         1.588 ns/op	       0 B/op	       0 allocs/op
BenchmarkRWMutexMapParallelLoad-8       	12766698	        96.07 ns/op	      63 B/op	       1 allocs/op
```
