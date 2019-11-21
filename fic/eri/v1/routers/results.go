package routers

import (
	fic "github.com/nttcom/go-fic"
	"github.com/nttcom/go-fic/pagination"
)

type commonResult struct {
	fic.Result
}

// Extract is a function that accepts a result
// and extracts a router resource.
func (r commonResult) Extract() (*Router, error) {
	var router Router
	err := r.ExtractInto(&router)
	return &router, err
}

// Extract interprets any commonResult as a Router, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "router")
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Router.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Router.
type GetResult struct {
	commonResult
}

// // UpdateResult represents the result of an update operation. Call its Extract
// // method to interpret it as a Router.
// type UpdateResult struct {
// 	commonResult
// }

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	fic.ErrResult
}

// RoutingGroup represents each element of routingGroups
// in router resource.
type RoutingGroup struct {
	Name string `json:"name"`
}

// Firewall represents each element of firewalls
// in router resource.
type Firewall struct {
	ID          string `json:"id"`
	IsActivated bool   `json:"isActivated"`
}

// NAT represents each element of nats in router resource.
type NAT struct {
	ID          string `json:"id"`
	IsActivated bool   `json:"isActivated"`
}

// Router represents router resource.
type Router struct {
	ID              string         `json:"id"`
	TenantID        string         `json:"tenantId"`
	Name            string         `json:"name"`
	Area            string         `json:"area"`
	UserIPAddress   string         `json:"userIpAddress"`
	Redundant       *bool          `json:"redundant"`
	OperationStatus string         `json:"operationStatus"`
	Firewalls       []Firewall     `json:"firewalls"`
	NATs            []NAT          `json:"nats"`
	RoutingGroups   []RoutingGroup `json:"routingGroups"`
	OperationID     string         `json:"operationId"`
}

// RouterPage is the page returned by a pager
// when traversing over a collection of routers.
type RouterPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of routers
// have reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r RouterPage) NextPageURL() (string, error) {
	var s struct {
		Links []fic.Link `json:"routers_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return fic.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a RouterPage struct is empty.
func (r RouterPage) IsEmpty() (bool, error) {
	is, err := ExtractRouters(r)
	return len(is) == 0, err
}

// ExtractRouters accepts a Page struct,
// specifically a RouterPage struct, and extracts the elements
// into a slice of Router structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractRouters(r pagination.Page) ([]Router, error) {
	var s []Router
	err := ExtractRoutersInto(r, &s)
	return s, err
}

// ExtractRoutersInto interprets the results of a single page from a List() call,
// producing a slice of Router entities.
func ExtractRoutersInto(r pagination.Page, v interface{}) error {
	return r.(RouterPage).Result.ExtractIntoSlicePtr(v, "routers")
}
