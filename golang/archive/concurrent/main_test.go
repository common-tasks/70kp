package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)
func TestUpdateMessage(t *testing.T) {

}
func Test_print(t *testing.T) {

	stdout := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go print("one", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)

	output := string(result)

	os.Stdout = stdout

	if !strings.Contains(output, "one") {
		t.Errorf("expected one but it is not there")
	}

}
