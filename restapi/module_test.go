package restapi

import (
	"testing"

	"github.com/biigo/biigo"
	"github.com/stretchr/testify/assert"

	"github.com/go-openapi/spec"
)

func TestModule(t *testing.T) {
	m := NewModule("/test/api", "/test/api/docs.json", &spec.Info{})

	assert.Implements(t, (*biigo.AppModule)(nil), m)
	assert.Implements(t, (*biigo.AppInitor)(nil), m)
}
