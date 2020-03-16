package nat_global_ip_address_sets

import (
	"github.com/nttcom/go-fic"
	"github.com/nttcom/go-fic/pagination"
)

type commonResult struct {
	fic.Result
}

// Extract is a function that accepts a result
// and extracts a global ip address set resource.
func (r commonResult) Extract() (*GlobalIPAddressSet, error) {
	var g GlobalIPAddressSet
	err := r.ExtractInto(&g)
	return &g, err
}

// Extract interprets any commonResult as a GlobalIPAddressSet, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "globalIpAddressSet")
}

// CreateResult represents the result of a activate operation.
// Call its Extract method to interpret it as a GlobalIPAddressSet.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation.
// Call its Extract method to interpret it as a GlobalIPAddressSet.
type GetResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation.
type DeleteResult struct {
	fic.ErrResult
}

// GlobalIPAddressSet represents global ip address set resource.
type GlobalIPAddressSet struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	Type              string   `json:"type"`
	NATComponentID    string   `json:"natComponentId"`
	OperationStatus   string   `json:"operationStatus"`
	TenantID          string   `json:"tenantId"`
	NumberOfAddresses int      `json:"numOfAddresses"`
	Addresses         []string `json:"addresses"`
	OperationID       string   `json:"operationId"`
}

// GlobalIPAddressSetPage is the page returned by a pager
// when traversing over a collection of global ip address sets.
type GlobalIPAddressSetPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of global ip address sets
// have reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r GlobalIPAddressSetPage) NextPageURL() (string, error) {
	var s struct {
		Links []fic.Link `json:"global_ip_address_set_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return fic.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a GlobalIPAddressSetPage struct is empty.
func (r GlobalIPAddressSetPage) IsEmpty() (bool, error) {
	is, err := ExtractGlobalIPAddressSets(r)
	return len(is) == 0, err
}

// ExtractGlobalIPAddressSets accepts a Page struct,
// specifically a GlobalIPAddressSetPage struct, and extracts the elements
// into a slice of Port structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractGlobalIPAddressSets(r pagination.Page) ([]GlobalIPAddressSet, error) {
	var s []GlobalIPAddressSet
	err := ExtractGlobalIPAddressSetsInto(r, &s)
	return s, err
}

// ExtractGlobalIPAddressSetsInto interprets the results of a single page from a List() call,
// producing a slice of GlobalIPAddressSetPage entities.
func ExtractGlobalIPAddressSetsInto(r pagination.Page, v interface{}) error {
	return r.(GlobalIPAddressSetPage).Result.ExtractIntoSlicePtr(v, "globalIpAddressSets")
}
