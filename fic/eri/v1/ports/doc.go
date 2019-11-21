/*
Package ports contains functionality for working with
FIC Port resources.

Example to List Ports

	listOpts := ports.ListOpts{}

	allPages, err := ports.List(client, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allPorts, err := ports.ExtractPorts(allPages)
	if err != nil {
		panic(err)
	}

	for _, port := range allPorts {
		fmt.Printf("%+v", port)
	}

Example to Create a Port

	createOpts := ports.CreateOpts{
		Name:          "YourPortName",
		SwitchName:    "SwitchName",
		NumberOfVLANs: 32,
		PortType:      "1G",
	}

	port, err := ports.Create(client, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a Port

	portID := "484cda0e-106f-4f4b-bb3f-d413710bbe78"
	err := ports.Delete(client, portID).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to Activate a Port

	port, err := ports.Activate(client, map[string]interface{}{}).Extract()
	if err != nil {
		panic(err)
	}

*/
package ports
