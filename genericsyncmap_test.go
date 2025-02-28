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
