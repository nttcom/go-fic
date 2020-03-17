package nat_global_ip_address_sets

import (
	"github.com/nttcom/go-fic"
	"github.com/nttcom/go-fic/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToGlobalIPAddressSetListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the global ip address set attributes you want to see returned.
type ListOpts struct {
}

// ToGlobalIPAddressSetListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToGlobalIPAddressSetListQuery() (string, error) {
	q, err := fic.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over
// a collection of global ip address sets.
// It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *fic.ServiceClient, routerID, natID string, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c, routerID, natID)
	if opts != nil {
		query, err := opts.ToGlobalIPAddressSetListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return GlobalIPAddressSetPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific global ip address set based on its unique ID.
func Get(c *fic.ServiceClient, routerID, natID, globalIPAddressSetID string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, routerID, natID, globalIPAddressSetID), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options used to activate a global ip address set.
type CreateOpts struct {
	Name              string `json:"name" required:"true"`
	Type              string `json:"type" required:"true"`
	NumberOfAddresses int    `json:"numOfAddresses" required:"true"`
}

// ToCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToCreateMap() (map[string]interface{}, error) {
	return fic.BuildRequestBody(opts, "globalIpAddressSet")
}

// Create accepts a CreateOpts struct and create a global ip address set
// using the values provided.
func Create(c *fic.ServiceClient, routerID, natID string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(createURL(c, routerID, natID), b, &r.Body, &fic.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return
}

// Delete accepts a unique ID and deletes the global ip address set associated with it.
func Delete(c *fic.ServiceClient, routerID, natID, globalIPAddressSetID string) (r DeleteResult) {
	_, r.Err = c.DeleteWithJSONResponse(deleteURL(c, routerID, natID, globalIPAddressSetID), &r.Body, &fic.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
