package testing

import (
	"fmt"

	con "github.com/nttcom/go-fic/fic/eri/v1/port_to_port_connections"
	"github.com/nttcom/go-fic/fic/eri/v1/ports"
)

const idConnection1 = "F010123456789"

var listResponse = fmt.Sprintf(`
{
    "connections": [
        {
            "id": "%s",
            "name": "YourConnectionName",
            "redundant": false,
            "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
            "area": "JPEAST",
            "operationStatus": "Completed",
            "source": {
                "portId": "F010123456789",
                "vlan": 101
            },
            "destination": {
                "portId": "F019876543210",
                "vlan": 102
            },
            "bandwidth": "100M"
        }
    ]
}
`,
	idConnection1,
)

var connection1 = con.Connection{
	ID:              idConnection1,
	Name:            "YourConnectionName",
	Redundant:       false,
	TenantID:        "87e89b8f075a4ee1aa209f6ca6ce242c",
	Area:            "JPEAST",
	OperationStatus: "Completed",
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

var expectedConnectionsSlice = []con.Connection{
	connection1,
}

var getResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "name": "YourConnectionName",
        "redundant": false,
        "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
        "area": "JPEAST",
        "operationStatus": "Completed",
        "source": {
            "portId": "F010123456789",
            "vlan": 101
        },
        "destination": {
            "portId": "F019876543210",
            "vlan": 102
        },
        "bandwidth": "100M"
    }
}
`,
	idConnection1,
)

const createRequestWithVLAN = `
{
    "connection": {
        "name": "YourConnectionName",
        "source": {
            "portId": "F010123456789",
            "vlan": 101
        },
        "destination": {
            "portId": "F019876543210",
            "vlan": 102
        },
        "bandwidth": "100M"
    }
}`

var createResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "name": "YourConnectionName",
        "redundant": false,
        "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
        "area": "JPEAST",
        "operationStatus": "Processing",
        "source": {
            "portId": "F010123456789",
            "vlan": 101
        },
        "destination": {
            "portId": "F019876543210",
            "vlan": 102
        },
        "bandwidth": "100M",
        "operationId": "a76e81f621894eadab6b98e60486d678"
    }
}
`,
	idConnection1,
)

var connectionCreated = con.Connection{
	ID:              idConnection1,
	Name:            "YourConnectionName",
	Redundant:       false,
	TenantID:        "87e89b8f075a4ee1aa209f6ca6ce242c",
	Area:            "JPEAST",
	OperationStatus: "Processing",
	Source: con.Source{
		PortID: "F010123456789",
		VLAN:   101,
	},
	Destination: con.Destination{
		PortID: "F019876543210",
		VLAN:   102,
	},
	Bandwidth:   "100M",
	OperationID: "a76e81f621894eadab6b98e60486d678",
}

const createRequestWithVLANRanges = `
    {
    	"port": {
            "name": "YourPortName",
            "switchName": "SwitchName",
            "vlanRanges": ["113-128"],
            "portType": "1G"
        }
    }
`

const createResponseWithVLANRanges = `
{
    "port": {
        "id": "F010123456789",
        "name": "YourPortName",
        "operationStatus": "Processing",
        "isActivated": false,
        "vlanRanges": [
            "113-128"
        ],
        "tenantId": "0dc3a0ff7cbc4f49b7ec70ac6af57104",
        "switchName": "SwitchName",
        "portType": "1G",
        "location": "NW1",
        "area": "JPEAST",
        "vlans": [
            {
                "vid": 32,
                "status": "unused"
            }
        ],
        "operationId": "80a7b22906be4badb2ba8c096b17c572"
    }
}`

var connectionCreatedWithVLANRanges = ports.Port{
	ID:              idConnection1,
	Name:            "YourPortName",
	OperationStatus: "Processing",
	IsActivated:     false,
	VLANRanges: []string{
		"113-128",
	},
	TenantID:   "0dc3a0ff7cbc4f49b7ec70ac6af57104",
	SwitchName: "SwitchName",
	PortType:   "1G",
	Location:   "NW1",
	Area:       "JPEAST",
	VLANs: []ports.VLAN{
		{
			VID:    32,
			Status: "unused",
		},
	},
	OperationID: "80a7b22906be4badb2ba8c096b17c572",
}
