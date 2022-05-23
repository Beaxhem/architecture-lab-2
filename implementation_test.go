package lab2

import (
	"fmt"
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

func (s *MySuite) TestPostfixToPrefixLongExpressions(c *C) {
	res, err := PostfixToPrefix("4 2 - 3 * 5 + 6 / 2 +")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "+ / + * - 4 2 3 5 6 2")
}

func (s *MySuite) TestNotEnoughOperands(c *C) {
	res, err := PostfixToPrefix("+ * + 2 2 3 5")
	c.Assert(err, ErrorMatches, "not enough operands")
	c.Assert(res, Equals, "")
}

func (s *MySuite) TestNotNumbers(c *C) {
	res, err := PostfixToPrefix("4 2 - 3 * A +")
	c.Assert(err, ErrorMatches, "A is not a number")
	c.Assert(res, Equals, "")
}

func (s *MySuite) TestEmptyExpression(c *C) {
	res, err := PostfixToPrefix("")
	c.Assert(err, ErrorMatches, "empty string")
	c.Assert(res, Equals, "")
}

func (s *MySuite) TestOperatorOnly(c *C) {
	res, err := PostfixToPrefix("+")
	c.Assert(err, ErrorMatches, "not enough operands")
	c.Assert(res, Equals, "")
}

func (s *MySuite) TestOneNumberOnly(c *C) {
	res, err := PostfixToPrefix("1")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "1")
}

func (s *MySuite) TestTwoNumbersOnly(c *C) {
	res, err := PostfixToPrefix("1 2")
	c.Assert(err, ErrorMatches, "not enough operands")
	c.Assert(res, Equals, "")
}

func ExamplePostfixToPrefix() {
	res, _ := PostfixToPrefix("4 2 - 3 * 5 +")
	fmt.Println(res)

	// Output:
	// + * - 4 2 3 5
}
