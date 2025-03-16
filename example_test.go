package gsm_test

import (
	"fmt"
	"slices"

	"github.com/aaronriekenberg/gsm"
)

func ExampleGenericSyncMap() {

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

	// Output:
	// alice value = {title:engineer age:25} ok = true
	// alice title = "engineer"
	// bob value = {title:manager age:35} ok = true
	// bob age = 35
	// swapped = true value = {title:manager age:25} ok = true
	// keys = [alice bob]
}
