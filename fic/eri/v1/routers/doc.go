/*
Package routers contains functionality for working with
FIC Router resources.

Example to List Routers

	listOpts := routers.ListOpts{}

	allPages, err := routers.List(client, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allRouters, err := routers.ExtractRouters(allPages)
	if err != nil {
		panic(err)
	}

	for _, router := range allRouters {
		fmt.Printf("%+v", router)
	}

Example to Create a Router

	createOpts := routers.CreateOpts{


	}

	port, err := routers.Create(client, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a Router

	routerID := "484cda0e-106f-4f4b-bb3f-d413710bbe78"
	err := routers.Delete(client, routerID).ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package routers
