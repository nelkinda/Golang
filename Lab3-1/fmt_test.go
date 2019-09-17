package Lab3_1

import (
	"fmt"
	"testing"
)

func Test_fmt(t *testing.T) {
	p := Point{ x : 10, y : 20 }
	if expected, actual := "(10, 20)", fmt.Sprintf("TODO", p); expected != actual {
		t.Errorf("Expected and actual differ:\n<%s>\n<%s>\n", expected, actual)
	}
}