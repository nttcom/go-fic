/*
Package ports contains functionality for working with
FIC Port resources.

Example to List Connections

	import (
		con "github.com/nttcom/go-fic/fic/eri/v1/port_to_port_connections"
	)

	listOpts := con.ListOpts{}

	allPages, err := con.List(fakeclient.ServiceClient(), nil).AllPages()
	if err != nil {
		panic(err)
	}

	allConnections, err := con.ExtractPorts(allPages)
	if err != nil {
		panic(err)
	}

	for _, c := range allConnections {
		fmt.Printf("%+v", c)
	}


Example to Get Connection

	connectionID := "484cda0e-106f-4f4b-bb3f-d413710bbe78"
	c, err := con.Get(fakeclient.ServiceClient(), connectionID).Extract()
	if err != nil {
		panic(err)
	}


Example to Create a Connection

	createOpts := con.CreateOpts{
		Name: "YourConnectionName",
		Source: con.Source{
			PortID: "F010123456789",
			VLAN:   101,
		},
		Destination: con.Destination{
			PortID: "F019876543210",
			VLAN:   102,
		},
		Bandwidth: "100M",
	}
	c, err := con.Create(fakeclient.ServiceClient(), createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a Connection

	connectionID := "484cda0e-106f-4f4b-bb3f-d413710bbe78"
	res := con.Delete(fakeclient.ServiceClient(), idConnection1)
	if err != nil {
		panic(err)
	}

*/
package port_to_port_connections
