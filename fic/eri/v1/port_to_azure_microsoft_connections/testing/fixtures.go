package testing

import (
	"fmt"

	con "github.com/nttcom/go-fic/fic/eri/v1/port_to_azure_microsoft_connections"
)

const idConnection1 = "F010123456789"

var listResponse = fmt.Sprintf(`
{
    "connections": [
        {
            "id": "%s",
            "name": "YourConnectionName",
            "redundant": true,
            "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
            "area": "JPEAST",
            "operationStatus": "Processing",
            "bandwidth": "100M",
            "source": {
                "asn": "65530",
                "primary": {
                    "portId": "F010123456789",
                    "vlan": 1025
                },
                "secondary": {
                    "portId": "F019876543210",
                    "vlan": 1057
                }
            },
            "destination": {
                "interconnect": "Tokyo-1",
                "qosType": "guarantee",
                "serviceKey": "6191af11-82f9-4c15-9894-9a69c8f8628a",
                "sharedKey": "a8268f8c96a9",
                "advertisedPublicPrefixes": [
                    "100.100.1.1/32",
                    "100.100.1.2/32"
                ],
                "routingRegistryName": "APNIC"
            },
            "primaryConnectedNwAddress": "10.10.0.0/30",
            "secondaryConnectedNwAddress": "10.20.0.0/30"
        }
    ]
}
`,
	idConnection1,
)

var connection1 = con.Connection{
	ID:              idConnection1,
	TenantID:        "87e89b8f075a4ee1aa209f6ca6ce242c",
	OperationStatus: "Processing",
	Redundant:       true,
	Name:            "YourConnectionName",
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
		AdvertisedPublicPrefixes: []string{
			"100.100.1.1/32",
			"100.100.1.2/32",
		},
		RoutingRegistryName: "APNIC",
	},
	Bandwidth:                        "100M",
	PrimaryConnectedNetworkAddress:   "10.10.0.0/30",
	SecondaryConnectedNetworkAddress: "10.20.0.0/30",
	Area:                             "JPEAST",
}

var expectedConnectionsSlice = []con.Connection{
	connection1,
}

var getResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "name": "YourConnectionName",
        "redundant": true,
        "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
        "area": "JPEAST",
        "operationStatus": "Processing",
        "bandwidth": "100M",
        "source": {
            "asn": "65530",
            "primary": {
                "portId": "F010123456789",
                "vlan": 1025
            },
            "secondary": {
                "portId": "F019876543210",
                "vlan": 1057
            }
        },
        "destination": {
            "interconnect": "Tokyo-1",
            "qosType": "guarantee",
            "serviceKey": "6191af11-82f9-4c15-9894-9a69c8f8628a",
            "sharedKey": "a8268f8c96a9",
            "advertisedPublicPrefixes": [
                "100.100.1.1/32",
                "100.100.1.2/32"
            ],
            "routingRegistryName": "APNIC"
        },
        "primaryConnectedNwAddress": "10.10.0.0/30",
        "secondaryConnectedNwAddress": "10.20.0.0/30"
    }
}
`,
	idConnection1,
)

const createRequest = `
{
    "connection": {
        "name": "YourConnectionName",
        "source": {
            "primary": {
                "portId": "F010123456789",
                "vlan": 1025
            },
            "secondary": {
                "portId": "F019876543210",
                "vlan": 1057
            },
            "asn": "65530"
        },
        "destination": {
            "interconnect": "Tokyo-1",
            "serviceKey": "6191af11-82f9-4c15-9894-9a69c8f8628a",
            "sharedKey": "a8268f8c96a9",
            "qosType": "guarantee",
            "advertisedPublicPrefixes": [
                "100.100.1.1/32",
                "100.100.1.2/32"
            ],
            "routingRegistryName": "APNIC"
        },
        "bandwidth": "100M",
        "primaryConnectedNwAddress": "10.10.0.0/30",
        "secondaryConnectedNwAddress": "10.20.0.0/30"
    }
}
`

var createResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "name": "YourConnectionName",
        "redundant": true,
        "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
        "area": "JPEAST",
        "operationStatus": "Processing",
        "bandwidth": "100M",
        "source": {
            "asn": "65530",
            "primary": {
                "portId": "F010123456789",
                "vlan": 1025
            },
            "secondary": {
                "portId": "F019876543210",
                "vlan": 1057
            }
        },
        "destination": {
            "interconnect": "Tokyo-1",
            "qosType": "guarantee",
            "serviceKey": "6191af11-82f9-4c15-9894-9a69c8f8628a",
            "sharedKey": "a8268f8c96a9",
            "advertisedPublicPrefixes": [
                "100.100.1.1/32",
                "100.100.1.2/32"
            ],
            "routingRegistryName": "APNIC"
        },
        "primaryConnectedNwAddress": "10.10.0.0/30",
        "secondaryConnectedNwAddress": "10.20.0.0/30",
        "operationId": "0c6449e13ad7454ba1d487e7615a2b39"
    }
}
`,
	idConnection1,
)

var connectionCreated = con.Connection{
	ID:              idConnection1,
	TenantID:        "87e89b8f075a4ee1aa209f6ca6ce242c",
	OperationStatus: "Processing",
	Area:            "JPEAST",
	Redundant:       true,
	Name:            "YourConnectionName",
	Bandwidth:       "100M",
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
		AdvertisedPublicPrefixes: []string{
			"100.100.1.1/32",
			"100.100.1.2/32",
		},
		RoutingRegistryName: "APNIC",
	},
	PrimaryConnectedNetworkAddress:   "10.10.0.0/30",
	SecondaryConnectedNetworkAddress: "10.20.0.0/30",
	OperationID:                      "0c6449e13ad7454ba1d487e7615a2b39",
}

const updateRequest = `
{
    "connection": {
        "destination": {
            "advertisedPublicPrefixes": [
                "100.100.1.1/32",
                "100.100.1.2/32",
                "100.100.1.3/32"
            ],
            "routingRegistryName": "ARIN"
        }
    }
}`

var updateResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "name": "YourConnectionName",
        "redundant": true,
        "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
        "area": "JPEAST",
        "operationStatus": "Processing",
        "bandwidth": "100M",
        "source": {
            "asn": "65530",
            "primary": {
                "portId": "F010123456789",
                "vlan": 1025
            },
            "secondary": {
                "portId": "F019876543210",
                "vlan": 1057
            }
        },
        "destination": {
            "interconnect": "Tokyo-1",
            "qosType": "guarantee",
            "serviceKey": "6191af11-82f9-4c15-9894-9a69c8f8628a",
            "sharedKey": "a8268f8c96a9",
            "advertisedPublicPrefixes": [
                "100.100.1.1/32",
                "100.100.1.2/32",
                "100.100.1.3/32"
            ],
            "routingRegistryName": "ARIN"
        },
        "primaryConnectedNwAddress": "10.10.0.0/30",
        "secondaryConnectedNwAddress": "10.20.0.0/30",
        "operationId": "8f3a27d111714af3b50830ab8a7f81d5"
    }
}`,
	idConnection1,
)

var connectionUpdated = con.Connection{
	ID:              idConnection1,
	TenantID:        "87e89b8f075a4ee1aa209f6ca6ce242c",
	OperationStatus: "Processing",
	Redundant:       true,
	Name:            "YourConnectionName",
	Bandwidth:       "100M",
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
		AdvertisedPublicPrefixes: []string{
			"100.100.1.1/32",
			"100.100.1.2/32",
			"100.100.1.3/32",
		},
		RoutingRegistryName: "ARIN",
	},
	PrimaryConnectedNetworkAddress:   "10.10.0.0/30",
	SecondaryConnectedNetworkAddress: "10.20.0.0/30",
	OperationID:                      "8f3a27d111714af3b50830ab8a7f81d5",
	Area:                             "JPEAST",
}
