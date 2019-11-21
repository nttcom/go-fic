package operations

import (
	fic "github.com/nttcom/go-fic"
	"github.com/nttcom/go-fic/pagination"
)

type commonResult struct {
	fic.Result
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Connection.
type GetResult struct {
	commonResult
}

// Extract is a function that accepts a result
// and extracts a operation resource.
func (r commonResult) Extract() (*Operation, error) {
	var o Operation
	err := r.ExtractInto(&o)
	return &o, err
}

// Extract interprets any commonResult as an Operation, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "operation")
}

// Operation represents operation resource.
type Operation struct {
	ID            string `json:"id"`
	TenantID      string `json:"tenantId"`
	ResourceID    string `json:"resourceId"`
	ResourceName  string `json:"resourceName"`
	ResourceType  string `json:"resourceType"`
	RequestType   string `json:"requestType"`
	Status        string `json:"status"`
	ReceptionTime string `json:"receptionTime"`
	CommitTime    string `json:"commitTime"`
	Error         string `json:"error"`
	RequestBody   string `json:"requestBody"`
}

// OperationPage is the page returned by a pager
// when traversing over a collection of operations.
type OperationPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of operations
// have reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r OperationPage) NextPageURL() (string, error) {
	var s struct {
		Links []fic.Link `json:"operations_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return fic.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a OperationPage struct is empty.
func (r OperationPage) IsEmpty() (bool, error) {
	is, err := ExtractOperations(r)
	return len(is) == 0, err
}

// ExtractOperations accepts a Page struct,
// specifically a OperationPage struct, and extracts the elements
// into a slice of Operation structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractOperations(r pagination.Page) ([]Operation, error) {
	var s []Operation
	err := ExtractOperationsInto(r, &s)
	return s, err
}

// ExtractOperationsInto interprets the results of a single page from a List() call,
// producing a slice of Operation entities.
func ExtractOperationsInto(r pagination.Page, v interface{}) error {
	return r.(OperationPage).Result.ExtractIntoSlicePtr(v, "operations")
}
