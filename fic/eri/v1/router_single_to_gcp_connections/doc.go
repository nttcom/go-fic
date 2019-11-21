/*
Package ports contains functionality for working with
FIC Connection resources.

Example to List Connections

	import (
		con "github.com/nttcom/go-fic/fic/eri/v1/router_single_to_gcp_connections"
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
			RouterID:  "F010123456789",
			GroupName: "group_1",
			RouteFilter: con.RouteFilter{
				In:  "noRoute",
				Out: "defaultRoute",
			},
			Primary: con.Primary{
				ASPathPrepend: con.ASPathPrepend{
					In:  &one,
					Out: &null,
				},
			},
		},
		Destination: con.Destination{
			Primary: con.DestinationHAInfo{
				Interconnect: "Equinix-TY2-1",
				ASN:          "65530",
			},
			AWSAccountID: "123456789012",
			QosType:      "guarantee",
		},
		Bandwidth:                      "100M",
		PrimaryConnectedNetworkAddress: "10.0.0.0/30",
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


Example to Update a Connection

	updateOpts := con.UpdateOpts{
		Source: con.SourceForUpdate{
			RouteFilter: con.RouteFilter{
				In:  "noRoute",
				Out: "privateRoute",
			},
			Primary: con.Primary{
				ASPathPrepend: con.ASPathPrepend{
					In:  &two,
					Out: &one,
				},
			},
		},
	}
	connectionID := "484cda0e-106f-4f4b-bb3f-d413710bbe78"
	c, err := con.Update(fakeclient.ServiceClient(), connectionID, updateOpts).Extract()
	if err != nil {
		panic(err)
	}

*/
package router_single_to_gcp_connections
