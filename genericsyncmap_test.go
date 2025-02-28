package genericsyncmap

import "testing"

func TestStoreAndLoad(t *testing.T) {
	var gsm GenericSyncMap[int, string]

	gsm.Store(1, "one")
	gsm.Store(2, "two")

	value, ok := gsm.Load(1)
	if ok != true || value != "one" {
		t.Errorf("gsm.Load(1) = (%q, %v) want (\"one\", true)", value, ok)
	}

	value, ok = gsm.Load(2)
	if ok != true || value != "two" {
		t.Errorf("gsm.Load(1) = (%q, %v) want (\"two\", true)", value, ok)
	}

	value, ok = gsm.Load(3)
	if ok != false || value != "" {
		t.Errorf("gsm.Load(1) = (%q, %v) want (\"\", true)", value, ok)
	}
}

func TestLoadAndDelete(t *testing.T) {
	var gsm GenericSyncMap[int, string]

	gsm.Store(1, "one")
	gsm.Store(2, "two")

	value, ok := gsm.LoadAndDelete(1)
	if ok != true || value != "one" {
		t.Errorf("gsm.LoadAndDelete(1) = (%q, %v) want (\"\", true)", value, ok)

	}

	value, ok = gsm.LoadAndDelete(1)
	if ok != false || value != "" {
		t.Errorf("gsm.LoadAndDelete(1) = (%q, %v) want (\"\", false)", value, ok)
	}

}
