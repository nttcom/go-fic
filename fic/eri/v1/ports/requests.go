package ports

import (
	fic "github.com/nttcom/go-fic"
	"github.com/nttcom/go-fic/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToPortListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the port attributes you want to see returned.
type ListOpts struct {
}

// ToPortListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToPortListQuery() (string, error) {
	q, err := fic.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over
// a collection of ports.
// It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *fic.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToPortListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return PortPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific port based on its unique ID.
func Get(c *fic.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, id), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToPortCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options used to create a port.
type CreateOpts struct {
	Name          string   `json:"name" required:"true"`
	SwitchName    string   `json:"switchName" required:"true"`
	NumberOfVLANs int      `json:"numOfVlans,omitempty"`
	VLANRanges    []string `json:"vlanRanges,omitempty"`
	PortType      string   `json:"portType" required:"true"`
}

// ToPortCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToPortCreateMap() (map[string]interface{}, error) {
	return fic.BuildRequestBody(opts, "port")
}

// Create accepts a CreateOpts struct and creates a port
// using the values provided.
// This operation does not actually require a request body, i.e. the
// CreateOpts struct argument can be empty.
func Create(c *fic.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToPortCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(createURL(c), b, &r.Body, &fic.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

func Activate(c *fic.ServiceClient, portID string) (r ActivateResult) {
	_, r.Err = c.Post(activateURL(c, portID), map[string]interface{}{}, &r.Body, &fic.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// Delete accepts a unique ID and deletes the port associated with it.
func Delete(c *fic.ServiceClient, portID string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, portID), nil)
	return
}
