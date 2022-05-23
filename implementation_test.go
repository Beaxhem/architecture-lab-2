package lab2

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestPostfixToPrefix(c *C) {
	res, err := PostfixToPrefix("4 2 - 3 * 5 +")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "+ * - 4 2 3 5")
}

func (s *MySuite) TestPostfixToPrefixFail(c *C) {
	res, err := PostfixToPrefix("+ 2 2")
	c.Assert(err, NotNil)
	c.Assert(res, Equals, "")
}
