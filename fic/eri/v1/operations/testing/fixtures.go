package testing

import (
	"github.com/nttcom/go-fic/fic/eri/v1/operations"
)

const listResponse = `
{
    "operations": [
        {
            "id": "cc43d0f05df24b1aabdea46456d46e39",
            "tenantId": "9d6a6682f0044660beae9701624e5dc4",
            "resourceId": "F0n0123456789",
            "resourceName": "YourResourceName",
            "resourceType": "Router",
            "requestType": "Delete",
            "status": "Completed",
            "receptionTime": "2020-07-24T06:16:46Z",
            "commitTime": "2020-07-24T06:16:49Z",
            "error": "",
            "requestBody": ""
        }
    ]
}
`

var operation1 = operations.Operation{
	ID:            "cc43d0f05df24b1aabdea46456d46e39",
	TenantID:      "9d6a6682f0044660beae9701624e5dc4",
	ResourceID:    "F0n0123456789",
	ResourceName:  "YourResourceName",
	ResourceType:  "Router",
	RequestType:   "Delete",
	Status:        "Completed",
	ReceptionTime: "2020-07-24T06:16:46Z",
	CommitTime:    "2020-07-24T06:16:49Z",
	Error:         "",
	RequestBody:   "",
}

var expectedOperationsSlice = []operations.Operation{
	operation1,
}

const getResponse = `
{
    "operation": {
        "id": "cc43d0f05df24b1aabdea46456d46e39",
        "tenantId": "9d6a6682f0044660beae9701624e5dc4",
        "resourceId": "F0n0123456789",
        "resourceName": "YourResourceName",
        "resourceType": "Router",
        "requestType": "Delete",
        "status": "Completed",
        "receptionTime": "2020-07-24T06:16:46Z",
        "commitTime": "2020-07-24T06:16:49Z",
        "error": "",
        "requestBody": ""

    }
}
`
