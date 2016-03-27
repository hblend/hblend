package location

import (
	"testing"

	. "gopkg.in/check.v1"
)

type World struct{}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func (w *World) SetUpSuite(c *C) {}

func (w *World) SetUpTest(c *C) {}

func (w *World) TearDownSuite(c *C) {}

var _ = Suite(&World{})
