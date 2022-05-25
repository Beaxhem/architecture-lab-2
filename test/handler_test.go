package tests

import (
	"testing"

	// . "github.com/Beaxhem/architecture-lab-2/handler"
	. "gopkg.in/check.v1"
)

func HandlerTest(t *testing.T) { TestingT(t) }

type HandlerSuite struct{}

var _ = Suite(&HandlerSuite{})
