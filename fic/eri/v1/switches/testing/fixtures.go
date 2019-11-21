package testing

import (
	"github.com/nttcom/go-fic/fic/eri/v1/switches"
)

const listResponse = `
{
    "switches": [
        {
            "id": "7ae9d49a2b234e818758257ecef8e1e7",
            "switchName": "SwitchName1",
            "area": "JPEAST",
            "location": "NW1",
            "portTypes": [
                {
                    "portType": "1G",
                    "available": true
                },
                {
                    "portType": "10G",
                    "available": false
                }
            ],
            "numOfAvailableVlans": 1488,
            "vlanRanges": [
                {
                    "vlanRange": "1009-1024",
                    "available": true
                },
                {
                    "vlanRange": "1025-1040",
                    "available": false
                }
            ]
        },
        {
            "id": "7ae9d49a2b234e818758257ecef8e1e8",
            "switchName": "SwitchName2",
            "area": "JPWEST",
            "location": "NW2",
            "portTypes": [
                {
                    "portType": "1G",
                    "available": true
                },
                {
                    "portType": "10G",
                    "available": false
                }
            ],
            "numOfAvailableVlans": 1588,
            "vlanRanges": [
                {
                    "vlanRange": "1009-1024",
                    "available": true
                },
                {
                    "vlanRange": "1025-1040",
                    "available": false
                }
            ]
        }    
    ]
}`

var portType = []switches.PortType{
	switches.PortType{
		Type:      "1G",
		Available: true,
	},
	switches.PortType{
		Type:      "10G",
		Available: false,
	},
}

var vlanRanges = []switches.VLANRange{
	switches.VLANRange{
		Range:     "1009-1024",
		Available: true,
	},
	switches.VLANRange{
		Range:     "1025-1040",
		Available: false,
	},
}

var switch1 = switches.Switch{
	ID:                     "7ae9d49a2b234e818758257ecef8e1e7",
	SwitchName:             "SwitchName1",
	Area:                   "JPEAST",
	Location:               "NW1",
	PortTypes:              portType,
	NumberOfAvailableVLANs: 1488,
	VLANRanges:             vlanRanges,
}

var switch2 = switches.Switch{
	ID:                     "7ae9d49a2b234e818758257ecef8e1e8",
	SwitchName:             "SwitchName2",
	Area:                   "JPWEST",
	Location:               "NW2",
	PortTypes:              portType,
	NumberOfAvailableVLANs: 1588,
	VLANRanges:             vlanRanges,
}

var expectedSwitchesSlice = []switches.Switch{
	switch1,
	switch2,
}
