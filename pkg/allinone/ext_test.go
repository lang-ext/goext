package allinone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToJsonString(t *testing.T) {
	type A struct {
		A1 string
	}
	a := A{A1: "x"}
	assert.Equal(t,UncheckToJsonString(a), `{"A1":"x"}`)
}
