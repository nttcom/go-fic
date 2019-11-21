package nats

import (
	fic "github.com/nttcom/go-fic"
)

func actionURL(c *fic.ServiceClient, routerID, natID, action string) string {
	return c.ServiceURL("routers", routerID, "nats", natID, action)
}

func updateURL(c *fic.ServiceClient, routerID, natID string) string {
	return c.ServiceURL("routers", routerID, "nats", natID)
}

func getURL(c *fic.ServiceClient, routerID, natID string) string {
	return c.ServiceURL("routers", routerID, "nats", natID)
}

func listURL(c *fic.ServiceClient, routerID string) string {
	return c.ServiceURL("routers", routerID, "nats")
}
