package routers

import (
	fic "github.com/nttcom/go-fic"
	"github.com/nttcom/go-fic/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToRouterListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the port attributes you want to see returned.
type ListOpts struct {
}

// ToRouterListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToRouterListQuery() (string, error) {
	q, err := fic.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over
// a collection of routers.
// It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *fic.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToRouterListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return RouterPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific router based on its unique ID.
func Get(c *fic.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, id), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToRouterCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options used to create a router.
type CreateOpts struct {
	Name          string `json:"name" required:"true"`
	Area          string `json:"area" required:"true"`
	UserIPAddress string `json:"userIpAddress" required:"true"`
	Redundant     *bool  `json:"redundant" required:"true"`
}

// ToRouterCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToRouterCreateMap() (map[string]interface{}, error) {
	return fic.BuildRequestBody(opts, "router")
}

// Create accepts a CreateOpts struct and creates a router
// using the values provided.
// This operation does not actually require a request body, i.e. the
// CreateOpts struct argument can be empty.
func Create(c *fic.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToRouterCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(createURL(c), b, &r.Body, &fic.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// Delete accepts a unique ID and deletes the router associated with it.
func Delete(c *fic.ServiceClient, routerID string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, routerID), nil)
	return
}
