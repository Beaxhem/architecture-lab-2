package tests

import (
	"fmt"
	"testing"

	. "github.com/Beaxhem/architecture-lab-2/implementation"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type ImplementationSuite struct{}

var _ = Suite(&ImplementationSuite{})

func (s *ImplementationSuite) TestPostfixToPrefix(c *C) {
	res, err := PostfixToPrefix("4 2 - 3 * 5 +")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "+ * - 4 2 3 5")
}

func (s *ImplementationSuite) TestPostfixToPrefixLongExpressions(c *C) {
	res, err := PostfixToPrefix("4 2 - 3 * 5 + 6 / 2")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "+ / + * - 4 2 3 5 6 2")
}

func (s *ImplementationSuite) TestNotEnoughOperands(c *C) {
	res, err := PostfixToPrefix("+ * + 2 2 3 5")
	c.Assert(err, ErrorMatches, "not enough operands")
	c.Assert(res, Equals, "")
}

func (s *ImplementationSuite) TestNotNumbers(c *C) {
	res, err := PostfixToPrefix("4 2 - 3 * A +")
	c.Assert(err, ErrorMatches, "A is not a number")
	c.Assert(res, Equals, "")
}

func (s *ImplementationSuite) TestEmptyExpression(c *C) {
	res, err := PostfixToPrefix("")
	c.Assert(err, ErrorMatches, "empty string")
	c.Assert(res, Equals, "")
}

func (s *ImplementationSuite) TestOperatorOnly(c *C) {
	res, err := PostfixToPrefix("+")
	c.Assert(err, ErrorMatches, "not enough operands")
	c.Assert(res, Equals, "")
}

func (s *ImplementationSuite) TestOneNumberOnly(c *C) {
	res, err := PostfixToPrefix("1")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "1")
}

func (s *ImplementationSuite) TestTwoNumbersOnly(c *C) {
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
