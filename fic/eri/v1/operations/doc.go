/*
Package operations contains functionality for working with
FIC Area resources.

Example to List Operations

	listOpts := operations.ListOpts{}

	allPages, err := operations.List(client, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allOperations, err := operations.ExtractOperations(allPages)
	if err != nil {
		panic(err)
	}

	for _, area := range allOperations {
		fmt.Printf("%+v", area)
	}


Example to Get Operation

	operationID := "484cda0e-106f-4f4b-bb3f-d413710bbe78"
	c, err := operations.Get(fakeclient.ServiceClient(), operationID).Extract()
	if err != nil {
		panic(err)
	}

*/
package operations
