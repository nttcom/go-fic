package testing

import (
	"github.com/nttcom/go-fic/fic/eri/v1/areas"
)

const listResponse = `
{
    "areas": [
        {
            "id": "ff2dfc4597754581a9a97ea32b853bca",
            "name": "JPEAST",
            "number": 1,
            "isPublic": true,
            "enabled": true
        },
        {
            "id": "ff2dfc4597754581a9a97ea32b853bcb",
            "name": "JPWEST",
            "number": 2,
            "isPublic": false,
            "enabled": false
        }
    ]
}`

var area1 = areas.Area{
	ID:       "ff2dfc4597754581a9a97ea32b853bca",
	Name:     "JPEAST",
	Number:   1,
	IsPublic: true,
	Enabled:  true,
}

var area2 = areas.Area{
	ID:       "ff2dfc4597754581a9a97ea32b853bcb",
	Name:     "JPWEST",
	Number:   2,
	IsPublic: false,
	Enabled:  false,
}

var expectedAreasSlice = []areas.Area{
	area1,
	area2,
}
