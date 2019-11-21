package switches

import (
	fic "github.com/nttcom/go-fic"
)

func resourceURL(c *fic.ServiceClient, id string) string {
	return c.ServiceURL("switches", id)
}

func rootURL(c *fic.ServiceClient) string {
	return c.ServiceURL("switches")
}

func listURL(c *fic.ServiceClient) string {
	return rootURL(c)
}
