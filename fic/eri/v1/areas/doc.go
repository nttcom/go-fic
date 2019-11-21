/*
Package areas contains functionality for working with
FIC Area resources.

Example to List Areas

	listOpts := areas.ListOpts{}

	allPages, err := areas.List(client, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allAreas, err := areas.ExtractAreas(allPages)
	if err != nil {
		panic(err)
	}

	for _, area := range allAreas {
		fmt.Printf("%+v", area)
	}
*/
package areas
