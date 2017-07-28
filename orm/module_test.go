package orm

import (
	"testing"

	"github.com/biigo/biigo"
	"github.com/stretchr/testify/assert"
)

func TestModule(t *testing.T) {
	m := NewModule("mysql", "testUrl")

	assert.Implements(t, (*biigo.Module)(nil), m)
	assert.Implements(t, (*biigo.AppInitor)(nil), m)
}
