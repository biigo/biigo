package biigo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type FooModule struct {
	MyName string
}

func (module FooModule) Name() string {
	return "fooModule"
}

func (module FooModule) GetModuleName() string {
	return module.MyName
}

func TestSetGetModule(t *testing.T) {
	fooModule := FooModule{MyName: "testModule"}
	AddModule(fooModule)

	m := Module(fooModule.Name()).(FooModule)
	assert.Equal(t, fooModule.MyName, m.GetModuleName())
}
