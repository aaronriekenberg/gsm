# gsm

Enhanced version of Go's [sync.Map](https://pkg.go.dev/sync#Map) with generic types and iterators.

## Example:

```
package main

import (
	"fmt"

	"github.com/aaronriekenberg/gsm"
)

func main() {
	var gsm gsm.GenericSyncMap[int, string]

	gsm.Store(1, "one")

	value, ok := gsm.Load(1)
	fmt.Printf("value = %q ok = %v\n", value, ok)

	value, ok = gsm.Load(2)
	fmt.Printf("value = %q ok = %v\n", value, ok)

	swapped := gsm.CompareAndSwap(1, "one", "updatedOne")
	value, ok = gsm.Load(1)
	fmt.Printf("swapped = %v value = %q ok = %v\n", swapped, value, ok)

}
```

output:

```
value = "one" ok = true
value = "" ok = false
swapped = true value = "updatedOne" ok = true
```

## Features:
1. Wrapping of all sync.Map methods with methods having generic parameter and return types.
2. Iterator methods `Range()`, `Keys()`, and `Values()`
3. [Unit tests](https://github.com/aaronriekenberg/gsm/blob/main/gsm_test.go) of every method
4. [Documentation](https://pkg.go.dev/github.com/aaronriekenberg/gsm) including [runnable example](https://pkg.go.dev/github.com/aaronriekenberg/gsm#example-GenericSyncMap)
