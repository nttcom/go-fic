package ports

import (
	"github.com/nttcom/go-fic"
	"github.com/nttcom/go-fic/pagination"
)

type commonResult struct {
	fic.Result
}

// Extract is a function that accepts a result
// and extracts a port resource.
func (r commonResult) Extract() (*Port, error) {
	var p Port
	err := r.ExtractInto(&p)
	return &p, err
}

// Extract interprets any commonResult as a Port, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "port")
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Port.
type CreateResult struct {
	commonResult
}

// ActivateResult represents the result of a activate operation.
// Call its Extract method to interpret it as a Port.
type ActivateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Port.
type GetResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	fic.ErrResult
}

type VLAN struct {
	VID    int    `json:"vid"`
	Status string `json:"status"`
}

// Port represents port resource.
type Port struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	OperationStatus string   `json:"operationStatus"`
	IsActivated     bool     `json:"isActivated"`
	VLANRanges      []string `json:"vlanRanges"`
	TenantID        string   `json:"tenantId"`
	SwitchName      string   `json:"switchName"`
	PortType        string   `json:"portType"`
	Location        string   `json:"location"`
	Area            string   `json:"area"`
	VLANs           []VLAN   `json:"vlans"`
	OperationID     string   `json:"operationId"`
}

// PortPage is the page returned by a pager
// when traversing over a collection of ports.
type PortPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of ports
// have reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r PortPage) NextPageURL() (string, error) {
	var s struct {
		Links []fic.Link `json:"ports_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return fic.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a PortPage struct is empty.
func (r PortPage) IsEmpty() (bool, error) {
	is, err := ExtractPorts(r)
	return len(is) == 0, err
}

// ExtractPorts accepts a Page struct,
// specifically a PortPage struct, and extracts the elements
// into a slice of Port structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractPorts(r pagination.Page) ([]Port, error) {
	var s []Port
	err := ExtractPortsInto(r, &s)
	return s, err
}

// ExtractPortsInto interprets the results of a single page from a List() call,
// producing a slice of Port entities.
func ExtractPortsInto(r pagination.Page, v interface{}) error {
	return r.(PortPage).Result.ExtractIntoSlicePtr(v, "ports")
}
