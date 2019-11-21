package router_paired_to_aws_connections

import (
	fic "github.com/nttcom/go-fic"
	"github.com/nttcom/go-fic/pagination"
)

type commonResult struct {
	fic.Result
}

// Extract is a function that accepts a result
// and extracts a connection resource.
func (r commonResult) Extract() (*Connection, error) {
	var c Connection
	err := r.ExtractInto(&c)
	return &c, err
}

// Extract interprets any commonResult as a Connection, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "connection")
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Connection.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Connection.
type GetResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	fic.ErrResult
}

// UpdateResult represents the result of a update operation.
type UpdateResult struct {
	commonResult
}

// Connection represents connection resource.
type Connection struct {
	ID                               string      `json:"id"`
	TenantID                         string      `json:"tenantId"`
	OperationStatus                  string      `json:"operationStatus"`
	Redundant                        bool        `json:"redundant"`
	Name                             string      `json:"name"`
	Bandwidth                        string      `json:"bandwidth"`
	Source                           Source      `json:"source"`
	Destination                      Destination `json:"destination"`
	PrimaryConnectedNetworkAddress   string      `json:"primaryConnectedNwAddress"`
	SecondaryConnectedNetworkAddress string      `json:"secondaryConnectedNwAddress"`
	OperationID                      string      `json:"operationId"`
}

// ConnectionPage is the page returned by a pager
// when traversing over a collection of connections.
type ConnectionPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of connections
// have reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r ConnectionPage) NextPageURL() (string, error) {
	var s struct {
		Links []fic.Link `json:"connections_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return fic.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a ConnectionPage struct is empty.
func (r ConnectionPage) IsEmpty() (bool, error) {
	is, err := ExtractConnections(r)
	return len(is) == 0, err
}

// ExtractConnections accepts a Page struct,
// specifically a ConnectionPage struct, and extracts the elements
// into a slice of Connection structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractConnections(r pagination.Page) ([]Connection, error) {
	var s []Connection
	err := ExtractConnectionsInto(r, &s)
	return s, err
}

// ExtractConnectionsInto interprets the results of a single page from a List() call,
// producing a slice of Connection entities.
func ExtractConnectionsInto(r pagination.Page, v interface{}) error {
	return r.(ConnectionPage).Result.ExtractIntoSlicePtr(v, "connections")
}
