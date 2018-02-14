package main

import "testing"

func TestScanHost(t *testing.T) {
	cs, err := ScanHost("securenetwork.it:443")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", cs)
	for _, c := range cs {
		t.Log(c.Name)
	}
}
