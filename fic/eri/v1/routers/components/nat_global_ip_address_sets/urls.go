package nat_global_ip_address_sets

import (
	"github.com/nttcom/go-fic"
)

func listURL(c *fic.ServiceClient, routerID, natID string) string {
	return c.ServiceURL("routers", routerID, "nats", natID, "global-ip-address-sets")
}

func createURL(c *fic.ServiceClient, routerID, natID string) string {
	return c.ServiceURL("routers", routerID, "nats", natID, "global-ip-address-sets")
}

func deleteURL(c *fic.ServiceClient, routerID, natID, globalIPAddressSetID string) string {
	return c.ServiceURL("routers", routerID, "nats", natID, "global-ip-address-sets", globalIPAddressSetID)
}

func getURL(c *fic.ServiceClient, routerID, natID, globalIPAddressSetID string) string {
	return c.ServiceURL("routers", routerID, "nats", natID, "global-ip-address-sets", globalIPAddressSetID)
}
