package nats

import (
	fic "github.com/nttcom/go-fic"
	"github.com/nttcom/go-fic/pagination"
)

type commonResult struct {
	fic.Result
}

// Extract is a function that accepts a result
// and extracts a nat resource.
func (r commonResult) Extract() (*NAT, error) {
	var n NAT
	err := r.ExtractInto(&n)
	return &n, err
}

// Extract interprets any commonResult as a NAT, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "nat")
}

// ActivateResult represents the result of a activate operation. Call its Extract
// method to interpret it as a NAT.
type ActivateResult struct {
	commonResult
}

// UpdateResult represents the result of a update operation.
// Call its Extract method to interpret it as a NAT.
type UpdateResult struct {
	commonResult
}

// GetResult represents the result of a get operation.
// Call its Extract method to interpret it as a NAT.
type GetResult struct {
	commonResult
}

// DeactivateResult represents the result of a deactivate operation.
// Call its Extract method to interpret it as a NAT.
type DeactivateResult struct {
	commonResult
}

// GlobalIpAddressSets represents
// global ip address sets resource.
type GlobalIpAddressSets struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	Type              string   `json:"type"`
	NATComponentID    string   `json:"natComponentId"`
	OperationStatus   string   `json:"operationStatus"`
	TenantID          string   `json:"tenantId"`
	NumberOfAddresses int      `json:"numOfAddresses"`
	Addresses         []string `json:"addresses"`
}

// Match represents Match parameter in destination address.
// type Match struct {
// 	DestinationAddress string `json:"destinationAddress"`
// }

// EntryOfSourceNAPTRule represents Entries parameter in source NAPT rule.
// type EntryOfSourceNAPTRule struct {
// 	// Match Match    `json:"match"`
// 	Then []string `json:"then"`
// }

// EntryOfDestinationNATRule represents each element of Entries parameter
// in destination NAT rule.
// type EntryOfDestinationNATRule struct {
// 	Match Match  `json:"match"`
// 	Then  string `json:"then"`
// }

// 11
// NAT represents nat resource.
type NAT struct {
	ID                  string                `json:"id"`
	TenantID            string                `json:"tenantId"`
	Redundant           bool                  `json:"redundant"`
	IsActivated         bool                  `json:"isActivated"`
	OperationStatus     string                `json:"operationStatus"`
	SourceNAPTRules     []SourceNAPTRule      `json:"sourceNaptRules"`
	DestinationNATRules []DestinationNATRule  `json:"destinationNatRules"`
	UserIPAddresses     []string              `json:"userIpAddresses"`
	GlobalIPAddressSets []GlobalIpAddressSets `json:"globalIpAddressSets"`
	OperationID         string                `json:"operationId"`
}

// NATPage is the page returned by a pager
// when traversing over a collection of nats.
type NATPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of nats
// have reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r NATPage) NextPageURL() (string, error) {
	var s struct {
		Links []fic.Link `json:"nats_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return fic.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a NATPage struct is empty.
func (r NATPage) IsEmpty() (bool, error) {
	is, err := ExtractNATs(r)
	return len(is) == 0, err
}

// ExtractNATs accepts a Page struct,
// specifically a NATPage struct, and extracts the elements
// into a slice of Port structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractNATs(r pagination.Page) ([]NAT, error) {
	var s []NAT
	err := ExtractNATsInto(r, &s)
	return s, err
}

// ExtractNATsInto interprets the results of a single page from a List() call,
// producing a slice of Port entities.
func ExtractNATsInto(r pagination.Page, v interface{}) error {
	return r.(NATPage).Result.ExtractIntoSlicePtr(v, "nats")
}
