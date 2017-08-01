package mail

import (
	"testing"

	"github.com/biigo/biigo"
	"github.com/stretchr/testify/assert"
)

func TestModule(t *testing.T) {
	m := Module(new(MockMailSender))
	assert.Implements(t, (*biigo.AppModule)(nil), m)
}
