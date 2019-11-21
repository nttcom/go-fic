/*
Package switches contains functionality for working with
FIC Switch resources.

Example to List Switches

	listOpts := switches.ListOpts{}

	allPages, err := switches.List(client, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allSwitches, err := switches.ExtractSwitches(allPages)
	if err != nil {
		panic(err)
	}

	for _, switch := range allSwitches {
		fmt.Printf("%+v", area)
	}
*/
package switches
