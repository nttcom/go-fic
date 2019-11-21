package ports

import (
	fic "github.com/nttcom/go-fic"
)

func resourceURL(c *fic.ServiceClient, id string) string {
	return c.ServiceURL("ports", id)
}

func rootURL(c *fic.ServiceClient) string {
	return c.ServiceURL("ports")
}

func getURL(c *fic.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func listURL(c *fic.ServiceClient) string {
	return rootURL(c)
}

func createURL(c *fic.ServiceClient) string {
	return rootURL(c)
}

func deleteURL(c *fic.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func activateURL(c *fic.ServiceClient, id string) string {
	return c.ServiceURL("ports", id, "activate")
}
