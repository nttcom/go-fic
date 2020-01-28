package testing

import (
	"fmt"

	"github.com/nttcom/go-fic/fic/eri/v1/routers"
)

const idRouter1 = "F020123456789"
const idRouter2 = "F020123456780"

var bTrue bool = true
var bFalse bool = false

var listResponse = fmt.Sprintf(`
{
    "routers": [
        {
            "id": "%s",
            "tenantId": "9d6a6682f0044660beae9701624e5dc4",
            "name": "YourRouterName1",
            "area": "JPEAST",
            "userIpAddress": "10.100.0.0/27",
            "redundant": true,
            "operationStatus": "Completed",
            "firewalls": [
                {
                    "id": "F040123456789",
                    "isActivated": false
                }
            ],
            "nats": [
                {
                    "id": "F050123456789",
                    "isActivated": false
                }
            ],
            "routingGroups": [
                {
                    "name": "group_1"
                },
                {
                    "name": "group_2"
                },
                {
                    "name": "group_3"
                },
                {
                    "name": "group_4"
                },
                {
                    "name": "group_5"
                },
                {
                    "name": "group_6"
                },
                {
                    "name": "group_7"
                },
                {
                    "name": "group_8"
                }
            ]
        },
        {
            "id": "%s",
            "tenantId": "9d6a6682f0044660beae9701624e5dc4",
            "name": "YourRouterName2",
            "area": "JPEAST",
            "userIpAddress": "10.100.1.0/27",
            "redundant": false,
            "operationStatus": "Completed",
            "firewalls": [
                {
                    "id": "F040123456789",
                    "isActivated": false
                }
            ],
            "nats": [
                {
                    "id": "F050123456789",
                    "isActivated": false
                }
            ],
            "routingGroups": [
                {
                    "name": "group_1"
                },
                {
                    "name": "group_2"
                },
                {
                    "name": "group_3"
                },
                {
                    "name": "group_4"
                },
                {
                    "name": "group_5"
                },
                {
                    "name": "group_6"
                },
                {
                    "name": "group_7"
                },
                {
                    "name": "group_8"
                }
            ]
        }
    ]
}
`,
	idRouter1,
	idRouter2,
)

var firewalls = []routers.Firewall{
	{
		ID:          "F040123456789",
		IsActivated: false,
	},
}

var nats = []routers.NAT{
	{
		ID:          "F050123456789",
		IsActivated: false,
	},
}

var routingGroups = []routers.RoutingGroup{
	{Name: "group_1"},
	{Name: "group_2"},
	{Name: "group_3"},
	{Name: "group_4"},
	{Name: "group_5"},
	{Name: "group_6"},
	{Name: "group_7"},
	{Name: "group_8"},
}

var router1 = routers.Router{
	ID:              idRouter1,
	TenantID:        "9d6a6682f0044660beae9701624e5dc4",
	Name:            "YourRouterName1",
	Area:            "JPEAST",
	UserIPAddress:   "10.100.0.0/27",
	Redundant:       &bTrue,
	OperationStatus: "Completed",
	Firewalls:       firewalls,
	NATs:            nats,
	RoutingGroups:   routingGroups,
}

var router2 = routers.Router{
	ID:              idRouter2,
	TenantID:        "9d6a6682f0044660beae9701624e5dc4",
	Name:            "YourRouterName2",
	Area:            "JPEAST",
	UserIPAddress:   "10.100.1.0/27",
	Redundant:       &bFalse,
	OperationStatus: "Completed",
	Firewalls:       firewalls,
	NATs:            nats,
	RoutingGroups:   routingGroups,
}

var expectedRoutersSlice = []routers.Router{
	router1,
	router2,
}

var getResponse = fmt.Sprintf(`
{
    "router": {
        "id": "%s",
        "tenantId": "9d6a6682f0044660beae9701624e5dc4",
        "name": "YourRouterName1",
        "area": "JPEAST",
        "userIpAddress": "10.100.0.0/27",
        "redundant": true,
        "operationStatus": "Completed",
        "firewalls": [
            {
                "id": "F040123456789",
                "isActivated": false
            }
        ],
        "nats": [
            {
                "id": "F050123456789",
                "isActivated": false
            }
        ],
        "routingGroups": [
            {
                "name": "group_1"
            },
            {
                "name": "group_2"
            },
            {
                "name": "group_3"
            },
            {
                "name": "group_4"
            },
            {
                "name": "group_5"
            },
            {
                "name": "group_6"
            },
            {
                "name": "group_7"
            },
            {
                "name": "group_8"
            }
        ]
    }
}
`,
	idRouter1,
)

const createRequest = `
{
    "router": {
        "name": "YourRouterName",
        "area": "JPEAST",
        "userIpAddress": "10.100.0.0/27",
        "redundant": true
    }
}`

const createResponse = `
{
    "router": {
        "id": "F020123456789",
        "tenantId": "9d6a6682f0044660beae9701624e5dc4",
        "name": "YourRouterName",
        "area": "JPEAST",
        "userIpAddress": "10.100.0.0/27",
        "redundant": true,
        "operationStatus": "Processing",
        "firewalls": [
            {
                "id": "F040123456789",
                "isActivated": false
            }
        ],
        "nats": [
            {
                "id": "F050123456789",
                "isActivated": false
            }
        ],
        "routingGroups": [
            {
                "name": "group_1"
            },
            {
                "name": "group_2"
            },
            {
                "name": "group_3"
            },
            {
                "name": "group_4"
            },
            {
                "name": "group_5"
            },
            {
                "name": "group_6"
            },
            {
                "name": "group_7"
            },
            {
                "name": "group_8"
            }
        ],
        "operationId": "64555c9f7de344e994c2cce2af90c102"
    }
}`

var router1Created = routers.Router{
	ID:              idRouter1,
	TenantID:        "9d6a6682f0044660beae9701624e5dc4",
	Name:            "YourRouterName",
	Area:            "JPEAST",
	UserIPAddress:   "10.100.0.0/27",
	Redundant:       &bTrue,
	OperationStatus: "Processing",
	Firewalls:       firewalls,
	NATs:            nats,
	RoutingGroups:   routingGroups,
	OperationID:     "64555c9f7de344e994c2cce2af90c102",
}
