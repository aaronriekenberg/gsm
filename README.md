# gsm

Enhanced version of Go's [sync.Map](https://pkg.go.dev/sync#Map) with generic types and iterators.

## Features:
1. Wrapping of all sync.Map methods with methods having generic parameter and return types.
2. Iterator methods `Range()`, `Keys()`, and `Values()`
3. [Unit tests](https://github.com/aaronriekenberg/gsm/blob/main/gsm_test.go) of every method
4. [Documentation](https://pkg.go.dev/github.com/aaronriekenberg/gsm) including [runnable example](https://pkg.go.dev/github.com/aaronriekenberg/gsm#example-GenericSyncMap)

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

