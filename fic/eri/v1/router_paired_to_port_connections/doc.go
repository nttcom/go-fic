/*
Package ports contains functionality for working with
FIC Connection resources.

Example to List Connections

	import (
		con "github.com/nttcom/go-fic/fic/eri/v1/router_to_port_connections"
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
			RouterID:  "F020123456789",
			GroupName: "group_1",
			Primary: con.SourceHAInfo{
				IPAddress: "10.0.0.1/30",
				ASPathPrepend: con.ASPathPrepend{
					In:  &four,
					Out: &four,
				},
				MED: &con.MED{
					Out: 10,
				},
			},
			Secondary: con.SourceHAInfo{
				IPAddress: "10.0.0.5/30",
				ASPathPrepend: con.ASPathPrepend{
					In:  &two,
					Out: &one,
				},
				MED: &con.MED{
					Out: 20,
				},
			},
			RouteFilter: con.RouteFilter{
				In:  "fullRoute",
				Out: "fullRouteWithDefaultRoute",
			},
		},
		Destination: con.Destination{
			Primary: con.DestinationHAInfo{
				PortID:    "F010123456789",
				VLAN:      101,
				IPAddress: "10.0.0.2/30",
				ASN:       "65000",
			},
			Secondary: con.DestinationHAInfo{
				PortID:    "F019876543210",
				VLAN:      102,
				IPAddress: "10.0.0.6/30",
				ASN:       "65000",
			},
		},
		Bandwidth: "100M",
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

	var two = interface{}(2)
	var null = interface{}(nil)
	updateOpts := con.UpdateOpts{
		Source: con.SourceForUpdate{
			Primary: con.SourceHAInfoForUpdate{
				ASPathPrepend: con.ASPathPrepend{
					In:  &null,
					Out: &two,
				},
				MED: &con.MED{
					Out: 30,
				},
			},
			Secondary: con.SourceHAInfoForUpdate{
				ASPathPrepend: con.ASPathPrepend{
					In:  &null,
					Out: &two,
				},
				MED: &con.MED{
					Out: 40,
				},
			},
			RouteFilter: con.RouteFilter{
				In:  "noRoute",
				Out: "fullRoute",
			},
		},
	}
	connectionID := "484cda0e-106f-4f4b-bb3f-d413710bbe78"
	c, err := con.Update(fakeclient.ServiceClient(), connectionID, updateOpts).Extract()
	if err != nil {
		panic(err)
	}

*/
package router_to_port_connections
