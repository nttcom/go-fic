package areas

import (
	"github.com/nttcom/go-fic"
)

func rootURL(c *fic.ServiceClient) string {
	return c.ServiceURL("areas")
}

func listURL(c *fic.ServiceClient) string {
	return rootURL(c)
}
