package testing

import (
	"fmt"

	con "github.com/nttcom/go-fic/fic/eri/v1/port_to_gcp_connections"
)

const idConnection1 = "F010123456789"

var listResponse = fmt.Sprintf(`
{
    "connections": [
        {
            "id": "%s",
            "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
            "operationStatus": "Completed",
            "redundant": false,
            "name": "YourConnectionName",
            "bandwidth": "100M",
            "source": {
                "portId": "F010123456789",
                "vlan": 1009
            },
            "destination": {
                "interconnect": "@Tokyo-CC2-1",
                "pairingKey": "40e744d1-8076-4add-86aa-c17cf958f9cb/asia-northeast1/2",
                "qosType": "guarantee"
            },
            "primaryConnectedNwAddress": "169.254.51.168/29"
        }
    ]
}
`,
	idConnection1,
)

var connection1 = con.Connection{
	ID:              idConnection1,
	TenantID:        "87e89b8f075a4ee1aa209f6ca6ce242c",
	OperationStatus: "Completed",
	Redundant:       false,
	Name:            "YourConnectionName",
	Bandwidth:       "100M",
	Source: con.Source{
		PortID: "F010123456789",
		VLAN:   1009,
	},
	Destination: con.Destination{
		Interconnect: "@Tokyo-CC2-1",
		PairingKey:   "40e744d1-8076-4add-86aa-c17cf958f9cb/asia-northeast1/2",
		QosType:      "guarantee",
	},
	PrimaryConnectedNetworkAddress: "169.254.51.168/29",
}

var expectedConnectionsSlice = []con.Connection{
	connection1,
}

var getResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
        "operationStatus": "Completed",
        "redundant": false,
        "name": "YourConnectionName",
        "bandwidth": "100M",
        "source": {
            "portId": "F010123456789",
            "vlan": 1009
        },
        "destination": {
            "interconnect": "@Tokyo-CC2-1",
            "pairingKey": "40e744d1-8076-4add-86aa-c17cf958f9cb/asia-northeast1/2",
            "qosType": "guarantee"
        },
        "primaryConnectedNwAddress": "169.254.51.168/29"
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
            "portId": "F010123456789",
            "vlan": 1009,
            "asn": "65530"
        },
        "destination": {
            "interconnect": "@Tokyo-CC2-1",
            "pairingKey": "40e744d1-8076-4add-86aa-c17cf958f9cb/asia-northeast1/2",
            "qosType": "guarantee"
        },
        "bandwidth": "100M"
    }
}
`

var createResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
        "operationStatus": "Processing",
        "redundant": false,
        "name": "YourConnectionName",
        "bandwidth": "100M",
        "source": {
            "portId": "F010123456789",
            "vlan": 1009
        },
        "destination": {
            "interconnect": "@Tokyo-CC2-1",
            "pairingKey": "40e744d1-8076-4add-86aa-c17cf958f9cb/asia-northeast1/2",
            "qosType": "guarantee"
        },
        "operationId": "da816305380e44abad08602a6a8ce66f"
    }
}
`,
	idConnection1,
)

var connectionCreated = con.Connection{
	ID:              idConnection1,
	TenantID:        "87e89b8f075a4ee1aa209f6ca6ce242c",
	OperationStatus: "Processing",
	Redundant:       false,
	Name:            "YourConnectionName",
	Bandwidth:       "100M",
	Source: con.Source{
		PortID: "F010123456789",
		VLAN:   1009,
	},
	Destination: con.Destination{
		Interconnect: "@Tokyo-CC2-1",
		PairingKey:   "40e744d1-8076-4add-86aa-c17cf958f9cb/asia-northeast1/2",
		QosType:      "guarantee",
	},
	OperationID: "da816305380e44abad08602a6a8ce66f",
}
