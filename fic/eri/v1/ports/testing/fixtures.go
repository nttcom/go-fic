package testing

import (
	"fmt"

	"github.com/nttcom/go-fic/fic/eri/v1/ports"
)

const idPort1 = "F010123456789"
const idPort2 = "F010123456780"

var listResponse = fmt.Sprintf(`
{
    "ports": [
        {
            "id": "%s",
            "name": "YourPortName1",
            "operationStatus": "Completed",
            "isActivated": true,
            "vlanRanges": [
                "113-128",
                "129-144"
            ],
            "tenantId": "0dc3a0ff7cbc4f49b7ec70ac6af57104",
            "switchName": "SwitchName1",
            "portType": "1G",
            "location": "NW1",
            "area": "JPEAST",
            "vlans": [
                {
                    "vid": 113,
                    "status": "unused"
                },
                {
                    "vid": 114,
                    "status": "used"
                }
            ]
        },
        {
            "id": "%s",
            "name": "YourPortName2",
            "operationStatus": "Completed",
            "isActivated": false,
            "vlanRanges": [
                "113-128",
                "129-144"
            ],
            "tenantId": "0dc3a0ff7cbc4f49b7ec70ac6af57104",
            "switchName": "SwitchName2",
            "portType": "1G",
            "location": "NW2",
            "area": "JPWEST",
            "vlans": [
                {
                    "vid": 113,
                    "status": "unused"
                },
                {
                    "vid": 114,
                    "status": "used"
                }
            ]
        }
    ]
}
`,
	idPort1,
	idPort2,
)

var port1 = ports.Port{
	ID:              idPort1,
	Name:            "YourPortName1",
	OperationStatus: "Completed",
	IsActivated:     true,
	VLANRanges: []string{
		"113-128",
		"129-144",
	},
	TenantID:   "0dc3a0ff7cbc4f49b7ec70ac6af57104",
	SwitchName: "SwitchName1",
	PortType:   "1G",
	Location:   "NW1",
	Area:       "JPEAST",
	VLANs: []ports.VLAN{
		ports.VLAN{
			VID:    113,
			Status: "unused",
		},
		ports.VLAN{
			VID:    114,
			Status: "used",
		},
	},
}

var port2 = ports.Port{
	ID:              idPort2,
	Name:            "YourPortName2",
	OperationStatus: "Completed",
	IsActivated:     false,
	VLANRanges: []string{
		"113-128",
		"129-144",
	},
	TenantID:   "0dc3a0ff7cbc4f49b7ec70ac6af57104",
	SwitchName: "SwitchName2",
	PortType:   "1G",
	Location:   "NW2",
	Area:       "JPWEST",
	VLANs: []ports.VLAN{
		ports.VLAN{
			VID:    113,
			Status: "unused",
		},
		ports.VLAN{
			VID:    114,
			Status: "used",
		},
	},
}

var expectedPortsSlice = []ports.Port{
	port1,
	port2,
}

var getResponse = fmt.Sprintf(`
{
    "port":
        {
            "id": "%s",
            "name": "YourPortName1",
            "operationStatus": "Completed",
            "isActivated": true,
            "vlanRanges": [
                "113-128",
                "129-144"
            ],
            "tenantId": "0dc3a0ff7cbc4f49b7ec70ac6af57104",
            "switchName": "SwitchName1",
            "portType": "1G",
            "location": "NW1",
            "area": "JPEAST",
            "vlans": [
                {
                    "vid": 113,
                    "status": "unused"
                },
                {
                    "vid": 114,
                    "status": "used"
                }
            ]
        }
}
`,
	idPort1,
)

const createRequestWithVLAN = `
    {
    	"port": {
            "name": "YourPortName",
            "switchName": "SwitchName",
            "numOfVlans": 32,
            "portType": "1G"
        }
    }
`

const createResponseWithVLAN = `
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

const activateRequest = `{}`

var port1CreatedWithVLAN = ports.Port{
	ID:              idPort1,
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
		ports.VLAN{
			VID:    32,
			Status: "unused",
		},
	},
	OperationID: "80a7b22906be4badb2ba8c096b17c572",
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

var port1CreatedWithVLANRanges = ports.Port{
	ID:              idPort1,
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
		ports.VLAN{
			VID:    32,
			Status: "unused",
		},
	},
	OperationID: "80a7b22906be4badb2ba8c096b17c572",
}
