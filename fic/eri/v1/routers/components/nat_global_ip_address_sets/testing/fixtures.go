package testing

import (
	"fmt"

	"github.com/nttcom/go-fic/fic/eri/v1/routers/components/nat_global_ip_address_sets"
)

const idRouter = "R000000000001"
const idNAT = "N000000000001"

const idGIP1 = "G000000000001"
const idGIP2 = "G000000000002"

var listResponse = fmt.Sprintf(`
{
    "globalIpAddressSets": [
        {
            "id": "%s",
            "name": "src-set-01",
            "type": "sourceNapt",
            "natComponentId": "F050123456789",
            "operationStatus": "Processing",
            "tenantId": "080f290761484afabbec22938adc6a2e",
            "numOfAddresses": 5,
            "addresses": [
                "100.131.66.79",
                "100.131.66.80",
                "100.131.66.81",
                "100.131.66.82",
                "100.131.66.83"
            ]
        },
        {
            "id": "%s",
            "name": "dst-set-01",
            "type": "destinationNat",
            "natComponentId": "F050123456789",
            "operationStatus": "Processing",
            "tenantId": "080f290761484afabbec22938adc6a2e",
            "numOfAddresses": 1,
            "addresses": [
                "100.131.65.2"
            ]
        }
    ]
}
`,
	idGIP1,
	idGIP2,
)

var gip1 = nat_global_ip_address_sets.GlobalIPAddressSet{
	ID:                idGIP1,
	Name:              "src-set-01",
	Type:              "sourceNapt",
	NATComponentID:    "F050123456789",
	OperationStatus:   "Processing",
	TenantID:          "080f290761484afabbec22938adc6a2e",
	NumberOfAddresses: 5,
	Addresses: []string{
		"100.131.66.79",
		"100.131.66.80",
		"100.131.66.81",
		"100.131.66.82",
		"100.131.66.83",
	},
}

var gip2 = nat_global_ip_address_sets.GlobalIPAddressSet{
	ID:                idGIP2,
	Name:              "dst-set-01",
	Type:              "destinationNat",
	NATComponentID:    "F050123456789",
	OperationStatus:   "Processing",
	TenantID:          "080f290761484afabbec22938adc6a2e",
	NumberOfAddresses: 1,
	Addresses: []string{
		"100.131.65.2",
	},
}

var expectedGlobalIPAddressSetSlice = []nat_global_ip_address_sets.GlobalIPAddressSet{
	gip1,
	gip2,
}

var getResponse = fmt.Sprintf(`
{
    "globalIpAddressSet": {
        "id": "%s",
        "name": "src-set-01",
        "type": "sourceNapt",
        "natComponentId": "F050123456789",
        "operationStatus": "Processing",
        "tenantId": "080f290761484afabbec22938adc6a2e",
        "numOfAddresses": 5,
        "addresses": [
            "100.131.66.79",
            "100.131.66.80",
            "100.131.66.81",
            "100.131.66.82",
            "100.131.66.83"
        ]
    }
}
`,
	idGIP1,
)

const createRequest = `
{
    "globalIpAddressSet": {
        "name": "src-set-02",
        "type": "sourceNapt",
        "numOfAddresses": 5
    }
}`

const createResponse = `
{
    "globalIpAddressSet": {
        "id": "2d9ae7b27152408f94caf5442ca9b73b",
        "name": "src-set-02",
        "type": "sourceNapt",
        "natComponentId": "F050123456789",
        "operationStatus": "Processing",
        "tenantId": "080f290761484afabbec22938adc6a2e",
        "numOfAddresses": 5,
        "addresses": [
            "100.131.66.12",
            "100.131.66.13",
            "100.131.66.14",
            "100.131.66.15",
            "100.131.66.16"
        ],
        "operationId": "b70562414a5242d8ab556afae976ca47"
    }
}
`

var gip1Created = nat_global_ip_address_sets.GlobalIPAddressSet{
	ID:                "2d9ae7b27152408f94caf5442ca9b73b",
	Name:              "src-set-02",
	Type:              "sourceNapt",
	NATComponentID:    "F050123456789",
	OperationStatus:   "Processing",
	TenantID:          "080f290761484afabbec22938adc6a2e",
	NumberOfAddresses: 5,
	Addresses: []string{
		"100.131.66.12",
		"100.131.66.13",
		"100.131.66.14",
		"100.131.66.15",
		"100.131.66.16",
	},
	OperationID: "b70562414a5242d8ab556afae976ca47",
}
