package port_to_port_connections

import (
	fic "github.com/nttcom/go-fic"
	"github.com/nttcom/go-fic/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToConnectionListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the connection attributes you want to see returned.
type ListOpts struct {
}

// ToConnectionListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToConnectionListQuery() (string, error) {
	q, err := fic.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over
// a collection of connections.
// It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *fic.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToConnectionListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return ConnectionPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific connection based on its unique ID.
func Get(c *fic.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, id), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToConnectionCreateMap() (map[string]interface{}, error)
}

type Source struct {
	PortID string `json:"portId" required:"true"`
	VLAN   int    `json:"vlan" required:"true"`
}

type Destination struct {
	PortID string `json:"portId" required:"true"`
	VLAN   int    `json:"vlan" required:"true"`
}

// CreateOpts represents options used to create a connection.
type CreateOpts struct {
	Name        string      `json:"name" required:"true"`
	Source      Source      `json:"source" required:"true"`
	Destination Destination `json:"destination,omitempty"`
	Bandwidth   string      `json:"bandwidth" required:"true"`
}

// ToConnectionCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToConnectionCreateMap() (map[string]interface{}, error) {
	return fic.BuildRequestBody(opts, "connection")
}

// Create accepts a CreateOpts struct and creates a connection
// using the values provided.
// This operation does not actually require a request body, i.e. the
// CreateOpts struct argument can be empty.
func Create(c *fic.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToConnectionCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(createURL(c), b, &r.Body, &fic.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// Delete accepts a unique ID and deletes the connection associated with it.
func Delete(c *fic.ServiceClient, connectionID string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, connectionID), nil)
	return
}
