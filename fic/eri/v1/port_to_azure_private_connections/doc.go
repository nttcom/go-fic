/*
Package ports contains functionality for working with
FIC Connection resources.

Example to List Connections

	import (
		con "github.com/nttcom/go-fic/fic/eri/v1/port_to_azure_private_connections"
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
			Primary: con.Primary{
				PortID: "F010123456789",
				VLAN:   1025,
			},
			Secondary: con.Secondary{
				PortID: "F019876543210",
				VLAN:   1057,
			},
			ASN: "65530",
		},
		Destination: con.Destination{
			Interconnect: "Tokyo-1",
			ServiceKey:   "6191af11-82f9-4c15-9894-9a69c8f8628a",
			SharedKey:    "a8268f8c96a9",
			QosType:      "guarantee",
		},
		Bandwidth:                   "100M",
		PrimaryConnectedNetworkAddress:   "10.10.0.0/30",
		SecondaryConnectedNetworkAddress: "10.20.0.0/30",
	}

	c, err := con.Create(fakeclient.ServiceClient(), createOpts).Extract()
	if err != nil {
		panic(err)
	}


Example to Delete a Connection

	connectionID := "484cda0e-106f-4f4b-bb3f-d413710bbe78"
	res := con.Delete(fakeclient.ServiceClient(), connectionID)
	if err != nil {
		panic(err)
	}

*/
package port_to_azure_private_connections
