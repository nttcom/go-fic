package router_to_uno_connections

import (
	"github.com/nttcom/go-fic"
)

func resourceURL(c *fic.ServiceClient, id string) string {
	return c.ServiceURL("router-to-uno-connections", id)
}

func rootURL(c *fic.ServiceClient) string {
	return c.ServiceURL("router-to-uno-connections")
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

func updateURL(c *fic.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func deleteURL(c *fic.ServiceClient, id string) string {
	return resourceURL(c, id)
}
