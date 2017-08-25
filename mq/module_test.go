package mq

import (
	"testing"

	"github.com/biigo/biigo"
	"github.com/stretchr/testify/assert"
)

func TestModule(t *testing.T) {
	m := NewModule()

	assert := assert.New(t)
	assert.Implements((*biigo.AppModule)(nil), m)
	assert.Implements((*biigo.AppConfiger)(nil), m)
	assert.Implements((*biigo.AppInitor)(nil), m)
}
