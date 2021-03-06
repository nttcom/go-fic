package port_to_azure_private_connections

import (
	"github.com/nttcom/go-fic"
)

func resourceURL(c *fic.ServiceClient, id string) string {
	return c.ServiceURL("port-to-azure-private-connections", id)
}

func rootURL(c *fic.ServiceClient) string {
	return c.ServiceURL("port-to-azure-private-connections")
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
