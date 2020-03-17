/*
Package nat_global_ip_sets contains functionality for working with
FIC Global IP Address Set in Component-NAT resources.

Example to List Global IP Address Sets

	listOpts := nat_global_ip_address_sets.ListOpts{}

	idRouter := "router1-id"
	idNAT := "nat1-id"
	allPages, err := nat_global_ip_address_sets.List(client, idRouter, idNAT, nil).AllPages()
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

	gip, err := nat_global_ip_address_sets.Get(client, idRouter, idNAT, idGIP).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", gip)


Example to Create a Global IP Address Set

	idRouter := "router1-id"
	idNAT := "nat1-id"

	createOpts := nat_global_ip_address_sets.CreateOpts{
		Name:              "src-set-02",
		Type:              "sourceNapt",
		NumberOfAddresses: 5,
	}
	gip, err := nat_global_ip_address_sets.Create(client, idRouter, idNAT, createOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", gip)


Example to Delete a Global IP Address Set

	idRouter := "router1-id"
	idNAT := "nat1-id"
	idGIP := "gip1-id"

	gip, err := nat_global_ip_address_sets.Delete(client, idRouter, idNAT, idGIP).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", gip)

*/
package nat_global_ip_address_sets
