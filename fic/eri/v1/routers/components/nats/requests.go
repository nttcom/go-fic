package nats

import (
	"github.com/nttcom/go-fic"
	"github.com/nttcom/go-fic/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToNATListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the nat attributes you want to see returned.
type ListOpts struct {
}

// ToNATListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToNATListQuery() (string, error) {
	q, err := fic.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over
// a collection of nats.
// It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *fic.ServiceClient, routerID string, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c, routerID)
	if opts != nil {
		query, err := opts.ToNATListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return NATPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific nat based on its unique ID.
func Get(c *fic.ServiceClient, routerID, natID string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, routerID, natID), &r.Body, nil)
	return
}

// ActivateOptsBuilder allows extensions to add additional parameters to the
// Activate request.
type ActivateOptsBuilder interface {
	ToActivateMap() (map[string]interface{}, error)
}

// GlobalIPAddressSet represents each element of
// GlobalIPAddressSets options for nat.
type GlobalIPAddressSet struct {
	Name              string `json:"name" required:"true"`
	Type              string `json:"type" required:"true"`
	NumberOfAddresses int    `json:"numOfAddresses" required:"true"`
}

// ActivateOpts represents options used to activate a nat.
type ActivateOpts struct {
	UserIPAddresses     []string             `json:"userIpAddresses,omitempty"`
	GlobalIPAddressSets []GlobalIPAddressSet `json:"globalIpAddressSets,omitempty"`
}

// ToActivateMap builds a request body from ActivateOpts.
func (opts ActivateOpts) ToActivateMap() (map[string]interface{}, error) {
	return fic.BuildRequestBody(opts, "nat")
}

// Activate accepts a ActivateOpts struct and activate a nat
// using the values provided.
func Activate(c *fic.ServiceClient, routerID, natID string, opts ActivateOptsBuilder) (r ActivateResult) {
	b, err := opts.ToActivateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(actionURL(c, routerID, natID, "activate"), b, &r.Body, &fic.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToUpdateMap() (map[string]interface{}, error)
}

// EntryInSourceNATRule represents Entry option in SourceNATRule.
type EntryInSourceNAPTRule struct {
	Then []string `json:"then" required:"true"`
}

// SourceNAPTRule represents SourceNATRule.
type SourceNAPTRule struct {
	From    []string                `json:"from" required:"true"`
	To      string                  `json:"to" required:"true"`
	Entries []EntryInSourceNAPTRule `json:"entries,omitempty"`
}

// DestinationNATRule represents each element of DestinationNATRule.
type DestinationNATRule struct {
	From    string                      `json:"from" required:"true"`
	To      string                      `json:"to" required:"true"`
	Entries []EntryInDestinationNATRule `json:"entries,omitempty"`
}

// EntryInDestinationNATRule represents Entry of DestinationNATRule.
type EntryInDestinationNATRule struct {
	Match Match  `json:"match" required:"true"`
	Then  string `json:"then" required:"true"`
}

// Match represents Match parameter in element of Entry of DestinationNATRule.
type Match struct {
	DestinationAddress string `json:"destinationAddress" required:"true"`
}

// UpdateOpts represents options used to activate a nat.
type UpdateOpts struct {
	SourceNAPTRules     []SourceNAPTRule     `json:"sourceNaptRules" required:"true"`
	DestinationNATRules []DestinationNATRule `json:"destinationNatRules" required:"true"`
}

// ToUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToUpdateMap() (map[string]interface{}, error) {
	return fic.BuildRequestBody(opts, "nat")
}

// Update accepts a UpdateOpts struct and update a nat
// using the values provided.
func Update(c *fic.ServiceClient, routerID, natID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(updateURL(c, routerID, natID), b, &r.Body, &fic.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// Deactivate deactivate a nat.
func Deactivate(c *fic.ServiceClient, routerID, natID string) (r DeactivateResult) {
	_, r.Err = c.Post(actionURL(c, routerID, natID, "deactivate"), map[string]interface{}{}, &r.Body, &fic.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
