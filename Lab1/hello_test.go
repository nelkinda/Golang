package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestLab1Main(t *testing.T) {
	var r, w *os.File
	var err error
	if r, w, err = os.Pipe(); err != nil {
		panic(err)
	}
	originalStdout := os.Stdout
	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		outC <- buf.String()
	}()

	main()

	_ = w.Close()
	os.Stdout = originalStdout
	actual := <-outC
	expected := "Hello, world!\n"
	if actual != expected {
		t.Errorf("Expected and actual differ:\n<%s>\n<%s>\n", expected, actual)
	}
	fmt.Println("Output:")
	fmt.Print(actual)
}
