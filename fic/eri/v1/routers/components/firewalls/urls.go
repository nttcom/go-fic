package firewalls

import (
	"github.com/nttcom/go-fic"
)

func actionURL(c *fic.ServiceClient, routerID, firewallID, action string) string {
	return c.ServiceURL("routers", routerID, "firewalls", firewallID, action)
}

func updateURL(c *fic.ServiceClient, routerID, firewallID string) string {
	return c.ServiceURL("routers", routerID, "firewalls", firewallID)
}

func getURL(c *fic.ServiceClient, routerID, firewallID string) string {
	return c.ServiceURL("routers", routerID, "firewalls", firewallID)
}

func listURL(c *fic.ServiceClient, routerID string) string {
	return c.ServiceURL("routers", routerID, "firewalls")
}
