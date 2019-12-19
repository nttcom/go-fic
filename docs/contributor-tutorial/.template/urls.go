package RESOURCE

import "github.com/nttcom/go-fic"

func listURL(client *fic.ServiceClient) string {
	return client.ServiceURL("resource")
}

func getURL(client *fic.ServiceClient, id string) string {
	return client.ServiceURL("resource", id)
}

func createURL(client *fic.ServiceClient) string {
	return client.ServiceURL("resource")
}

func deleteURL(client *fic.ServiceClient, id string) string {
	return client.ServiceURL("resource", id)
}

func updateURL(client *fic.ServiceClient, id string) string {
	return client.ServiceURL("resource", id)
}
