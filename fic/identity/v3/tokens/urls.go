package tokens

import fic "github.com/nttcom/go-fic"

func tokenURL(c *fic.ServiceClient) string {
	return c.ServiceURL("auth", "tokens")
}
