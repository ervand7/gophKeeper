package requests

func (c *Client) Logout() (resultMessage string) {
	if c.token == "" {
		return "not authorized"
	}
	c.reset()
	return "success logout"
}
