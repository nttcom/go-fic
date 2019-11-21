package testing

import (
	"fmt"

	con "github.com/nttcom/go-fic/fic/eri/v1/port_to_aws_connections"
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
                "qosType": "guarantee",
                "awsAccountId": "123456789012"
            }
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
		QosType:      "guarantee",
		AWSAccountID: "123456789012",
	},
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
            "qosType": "guarantee",
            "awsAccountId": "123456789012"
        }
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
            "vlan": 1009
        },
        "destination": {
            "interconnect": "@Tokyo-CC2-1",
            "awsAccountId": "123456789012",
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
            "qosType": "guarantee",
            "awsAccountId": "123456789012"
        },
        "operationId": "75f4cad2afcc436ebe3039275cdf0a7e"
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
		QosType:      "guarantee",
		AWSAccountID: "123456789012",
	},
	OperationID: "75f4cad2afcc436ebe3039275cdf0a7e",
}
