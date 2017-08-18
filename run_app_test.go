package biigo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestRunner struct {
	isRun bool
}

func (runner *TestRunner) RunApp() {
	runner.isRun = true
}

func (runner *TestRunner) Name() string {
	return "runner test module"
}

func TestRunApp(t *testing.T) {
	m := &TestRunner{}
	AddModule(m).Run()
	time.Sleep(100 * time.Microsecond)
	assert.True(t, m.isRun)
}
