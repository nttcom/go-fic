package operations

import (
	fic "github.com/nttcom/go-fic"
)

func resourceURL(c *fic.ServiceClient, id string) string {
	return c.ServiceURL("operations", id)
}

func rootURL(c *fic.ServiceClient) string {
	return c.ServiceURL("operations")
}

func getURL(c *fic.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func listURL(c *fic.ServiceClient) string {
	return rootURL(c)
}
