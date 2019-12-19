package switches

import (
	"github.com/nttcom/go-fic"
	"github.com/nttcom/go-fic/pagination"
)

type commonResult struct {
	fic.Result
}

// Extract is a function that accepts a result
// and extracts a switch resource.
func (r commonResult) Extract() (*Switch, error) {
	var s Switch
	err := r.ExtractInto(&s)
	return &s, err
}

// Extract interprets any commonResult as a Switch, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "switch")
}

type PortType struct {
	Type      string `json:"portType"`
	Available bool   `json:"available"`
}

type VLANRange struct {
	Range     string `json:"vlanRange"`
	Available bool   `json:"available"`
}

// Switch represents switch resource.
type Switch struct {
	ID                     string      `json:"id"`
	SwitchName             string      `json:"switchName"`
	Area                   string      `json:"area"`
	Location               string      `json:"location"`
	PortTypes              []PortType  `json:"portTypes"`
	NumberOfAvailableVLANs int         `json:"numOfAvailableVlans"`
	VLANRanges             []VLANRange `json:"vlanRanges"`
}

// SwitchPage is the page returned by a pager
// when traversing over a collection of switches.
type SwitchPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of switches
// have reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r SwitchPage) NextPageURL() (string, error) {
	var s struct {
		Links []fic.Link `json:"switches_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return fic.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a SwitchPage struct is empty.
func (r SwitchPage) IsEmpty() (bool, error) {
	is, err := ExtractSwitches(r)
	return len(is) == 0, err
}

// ExtractSwitches accepts a Page struct,
// specifically a SwitchPage struct, and extracts the elements
// into a slice of Switch structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractSwitches(r pagination.Page) ([]Switch, error) {
	var s []Switch
	err := ExtractSwitchesInto(r, &s)
	return s, err
}

// ExtractSwitchesInto interprets the results of a single page from a List() call,
// producing a slice of Switch entities.
func ExtractSwitchesInto(r pagination.Page, v interface{}) error {
	return r.(SwitchPage).Result.ExtractIntoSlicePtr(v, "switches")
}
