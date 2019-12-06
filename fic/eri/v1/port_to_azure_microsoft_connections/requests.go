package port_to_azure_microsoft_connections

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

// Primary represents primary port for connection.
type Primary struct {
	PortID string `json:"portId" required:"true"`
	VLAN   int    `json:"vlan" required:"true"`
}

// Secondary represents secondary port for connection.
type Secondary struct {
	PortID string `json:"portId" required:"true"`
	VLAN   int    `json:"vlan" required:"true"`
}

// Destination represents destination parameter for connection.
type Destination struct {
	Interconnect             string   `json:"interconnect" required:"true"`
	ServiceKey               string   `json:"serviceKey" required:"true"`
	SharedKey                string   `json:"sharedKey"`
	QosType                  string   `json:"qosType" required:"true"`
	AdvertisedPublicPrefixes []string `json:"advertisedPublicPrefixes" required:"true"`
	RoutingRegistryName      string   `json:"routingRegistryName"`
}

// Source represents source parameter for connection.
type Source struct {
	Primary   Primary   `json:"primary" required:"true"`
	Secondary Secondary `json:"secondary" required:"true"`
	ASN       string    `json:"asn" required:"true"`
}

// CreateOpts represents options used to create a connection.
type CreateOpts struct {
	Name                             string      `json:"name" required:"true"`
	Source                           Source      `json:"source" required:"true"`
	Destination                      Destination `json:"destination" required:"true"`
	Bandwidth                        string      `json:"bandwidth" required:"true"`
	PrimaryConnectedNetworkAddress   string      `json:"primaryConnectedNwAddress" required:"true"`
	SecondaryConnectedNetworkAddress string      `json:"secondaryConnectedNwAddress" required:"true"`
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

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToUpdateMap() (map[string]interface{}, error)
}

// DestinationForUpdate represents Destination parameter in case of Updating.
type DestinationForUpdate struct {
	AdvertisedPublicPrefixes []string `json:"advertisedPublicPrefixes" required:"true"`
	RoutingRegistryName      string   `json:"routingRegistryName"`
}

// UpdateOpts represents options used to update a connection.
type UpdateOpts struct {
	Destination DestinationForUpdate `json:"destination" required:"true"`
}

// ToUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToUpdateMap() (map[string]interface{}, error) {
	return fic.BuildRequestBody(opts, "connection")
}

// Update accepts a UpdateOpts struct and update a connection
// using the values provided.
func Update(c *fic.ServiceClient, connectionID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Patch(updateURL(c, connectionID), b, &r.Body, &fic.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
