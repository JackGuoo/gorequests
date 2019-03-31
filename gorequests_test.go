package gorequests

import (
	"strings"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	println("Timeout example1")
	req := Requests()
	req.Debug = 1

	// Timeout Second
	const timeout = 200 * time.Nanosecond
	req.SetTimeout(timeout)
	_, err := req.Get("https://httpbin.org/get")
	if err != nil {
		if strings.Contains(err.Error(), "Client.Timeout") {
			t.Skipf("host too slow to get fast resource in %v", timeout)
		}
		t.Fatal(err)
	}

}
