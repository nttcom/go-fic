/*
Package nat_global_ip_sets contains functionality for working with
FIC Global IP Address Set in Component-NAT resources.

Example to List Global IP Address Sets

	listOpts := nat_global_ip_address_sets.ListOpts{}

	idRouter := "router1-id"
	idNAT := "nat1-id"
	allPages, err := nat_global_ip_address_sets.List(fakeclient.ServiceClient(), idRouter, idNAT, nil).AllPages()
	if err != nil {
		panic(err)
	}

	allGIPs, err := nat_global_ip_address_sets.ExtractGlobalIPAddressSets(allPages)
	if err != nil {
		panic(err)
	}

	for _, gip := range allGIPs {
		fmt.Printf("%+v", gip)
	}


Example to Get Global IP Address Set

	idRouter := "router1-id"
	idNAT := "nat1-id"
	idGIP := "gip1-id"

	ip, err := nat_global_ip_address_sets.Get(fakeclient.ServiceClient(), idRouter, idNAT, idGIP).Extract()
	if err != nil {
		panic(err)
	}


Example to Create a Global IP Address Set

	idRouter := "router1-id"
	idNAT := "nat1-id"

	createOpts := nat_global_ip_address_sets.CreateOpts{
		Name:              "src-set-02",
		Type:              "sourceNapt",
		NumberOfAddresses: 5,
	}
	p, err := nat_global_ip_address_sets.Create(fakeclient.ServiceClient(), idRouter, idNAT, createOpts).Extract()
	if err != nil {
		panic(err)
	}


Example to Delete a Global IP Address Set

	idRouter := "router1-id"
	idNAT := "nat1-id"
	idGIP := "gip1-id"

	n, err := nat_global_ip_address_sets.Delete(client, idRouter, idNAT, idGIP)
	if err != nil {
		panic(err)
	}

*/
package nat_global_ip_address_sets
