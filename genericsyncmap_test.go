package genericsyncmap

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLoad(t *testing.T) {

	var gsm GenericSyncMap[int, string]

	gsm.Store(1, "one")

	tests := map[string]struct {
		key       int
		wantValue string
		wantOK    bool
	}{
		"successful load":   {key: 1, wantValue: "one", wantOK: true},
		"unsuccessful load": {key: 2, wantValue: "", wantOK: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			value, ok := gsm.Load(tc.key)

			diff := cmp.Diff(tc.wantValue, value)
			if diff != "" {
				t.Fatal(diff)
			}

			diff = cmp.Diff(tc.wantOK, ok)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestClear(t *testing.T) {

	var gsm GenericSyncMap[int, string]

	gsm.Store(1, "one")

	gsm.Clear()

	tests := map[string]struct {
		key       int
		wantValue string
		wantOK    bool
	}{
		"unsuccessful load": {key: 1, wantValue: "", wantOK: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			value, ok := gsm.Load(tc.key)

			diff := cmp.Diff(tc.wantValue, value)
			if diff != "" {
				t.Fatal(diff)
			}

			diff = cmp.Diff(tc.wantOK, ok)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestLoadOrStore(t *testing.T) {
	var gsm GenericSyncMap[int, string]

	tests := map[string]struct {
		key        int
		value      string
		wantActual string
		wantLoaded bool
	}{
		"first store":   {key: 1, value: "one", wantActual: "", wantLoaded: false},
		"second load":   {key: 1, value: "one", wantActual: "one", wantLoaded: true},
		"new key store": {key: 2, value: "two", wantActual: "", wantLoaded: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual, loaded := gsm.LoadOrStore(tc.key, tc.value)

			diff := cmp.Diff(tc.wantActual, actual)
			if diff != "" {
				t.Fatal(diff)
			}

			diff = cmp.Diff(tc.wantLoaded, loaded)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}

}

func TestLoadAndDelete(t *testing.T) {
	var gsm GenericSyncMap[int, string]

	gsm.Store(1, "one")

	tests := map[string]struct {
		key        int
		wantValue  string
		wantLoaded bool
	}{
		"first load":  {key: 1, wantValue: "one", wantLoaded: true},
		"second load": {key: 1, wantValue: "", wantLoaded: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			value, loaded := gsm.LoadAndDelete(tc.key)

			diff := cmp.Diff(tc.wantValue, value)
			if diff != "" {
				t.Fatal(diff)
			}

			diff = cmp.Diff(tc.wantLoaded, loaded)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}

}

func TestDelete(t *testing.T) {
	var gsm GenericSyncMap[int, string]

	gsm.Store(1, "one")

	tests := map[string]struct {
		key               int
		initialValue      string
		initialLoadOK     bool
		afterDeleteValue  string
		afterDeleteLoadOK bool
	}{
		"first load": {key: 1, initialValue: "one", initialLoadOK: true, afterDeleteValue: "", afterDeleteLoadOK: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			value, ok := gsm.Load(tc.key)

			diff := cmp.Diff(tc.initialValue, value)
			if diff != "" {
				t.Fatal(diff)
			}

			diff = cmp.Diff(tc.initialLoadOK, ok)
			if diff != "" {
				t.Fatal(diff)
			}

			gsm.Delete(tc.key)

			value, ok = gsm.Load(tc.key)

			diff = cmp.Diff(tc.afterDeleteValue, value)
			if diff != "" {
				t.Fatal(diff)
			}

			diff = cmp.Diff(tc.afterDeleteLoadOK, ok)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
