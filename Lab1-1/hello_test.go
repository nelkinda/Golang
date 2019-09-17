package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	originalStdout := os.Stdout
	defer func() { os.Stdout = originalStdout }()

	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		outC <- buf.String()
		_ = r.Close()
	}()

	main()
	_ = w.Close()

	if expected, actual := "Hello, world!\n", <-outC; expected != actual {
		t.Errorf("Expected and actual differ:\n<%s>\n<%s>\n", expected, actual)
	}
}
