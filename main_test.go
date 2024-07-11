package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_printThis(t *testing.T) {
	stdout := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go printThis("Found string", &wg)

	wg.Wait()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)
	output = strings.ReplaceAll(output, "\n", "")

	os.Stdout = stdout
	assert.Equal(t, "Found string", output)
}
