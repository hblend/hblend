package A

import "fmt"

func (c *Component) tag_include(token *gotreescript.Token) {
	s := c.require(token)

	fmt.Println("<h1>" + s.Name + "</h1>")

	*c.Html += *s.Html
}
