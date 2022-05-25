package tests

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Beaxhem/architecture-lab-2/handler"
	. "gopkg.in/check.v1"
)

func HandlerTest(t *testing.T) { TestingT(t) }

type HandlerSuite struct{}

var _ = Suite(&HandlerSuite{})

func (s *HandlerSuite) TestPostfixToPrefix(c *C) {
	var buff = new(bytes.Buffer)
	handler := handler.ComputeHandler{
		Reader: strings.NewReader("4 2 - 3 * 5 +"),
		Writer: buff,
	}
	err := handler.Compute()
	c.Assert(err, IsNil)
	c.Assert(buff.String(), Equals, "+ * - 4 2 3 5\n")
}

func (s *HandlerSuite) TestErrorIsReturnedOnBadInput(c *C) {
	var buff = new(bytes.Buffer)
	handler := handler.ComputeHandler{
		Reader: strings.NewReader("4 2 - 3 * 5"),
		Writer: buff,
	}
	err := handler.Compute()
	c.Assert(err, ErrorMatches, "not enough operands")
}
