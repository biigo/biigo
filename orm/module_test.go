package orm

import (
	"testing"

	"github.com/biigo/biigo"
	"github.com/stretchr/testify/assert"
)

func TestModule(t *testing.T) {
	m := NewModule(Config{})

	assert.Implements(t, (*biigo.AppModule)(nil), m)
	assert.Implements(t, (*biigo.AppInitor)(nil), m)
}
