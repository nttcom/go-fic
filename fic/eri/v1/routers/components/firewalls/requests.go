package firewalls

import (
	fic "github.com/nttcom/go-fic"
	"github.com/nttcom/go-fic/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToFirewallListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the firewall attributes you want to see returned.
type ListOpts struct {
}

// ToFirewallListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToFirewallListQuery() (string, error) {
	q, err := fic.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over
// a collection of firewalls.
// It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *fic.ServiceClient, routerID string, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c, routerID)
	if opts != nil {
		query, err := opts.ToFirewallListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return FirewallPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific firewall based on its unique ID.
func Get(c *fic.ServiceClient, routerID, firewallID string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, routerID, firewallID), &r.Body, nil)
	return
}

// ActivateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type ActivateOptsBuilder interface {
	ToFirewallCreateMap() (map[string]interface{}, error)
}

// ActivateOpts represents options used to create a firewall.
type ActivateOpts struct {
	UserIPAddresses []string `json:"userIpAddresses" required:"true"`
}

// ToFirewallCreateMap builds a request body from ActivateOpts.
func (opts ActivateOpts) ToFirewallCreateMap() (map[string]interface{}, error) {
	return fic.BuildRequestBody(opts, "firewall")
}

// Create accepts a ActivateOpts struct and creates a firwall
// using the values provided.
// This operation does not actually require a request body, i.e. the
// ActivateOpts struct argument can be empty.
func Activate(c *fic.ServiceClient, routerID, firewallID string, opts ActivateOptsBuilder) (r ActivateResult) {
	b, err := opts.ToFirewallCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(actionURL(c, routerID, firewallID, "activate"), b, &r.Body, &fic.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// Deactivate deactivate a firewall.
func Deactivate(c *fic.ServiceClient, routerID, natID string) (r DeactivateResult) {
	_, r.Err = c.Post(actionURL(c, routerID, natID, "deactivate"), map[string]interface{}{}, &r.Body, &fic.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToUpdateMap() (map[string]interface{}, error)
}

// Match represents request parameter to update firewall.
type Match struct {
	SourceAddressSets      []string `json:"sourceAddressSets" required:"true"`
	DestinationAddressSets []string `json:"destinationAddressSets" required:"true"`
	Application            string   `json:"application" required:"true"`
}

// Entry represents request parameter to update firewall.
type Entry struct {
	Name   string `json:"name" required:"true"`
	Match  Match  `json:"match" required:"true"`
	Action string `json:"action" required:"true"`
}

// Rule represents request parameter to update firewall.
type Rule struct {
	From    string  `json:"from" required:"true"`
	To      string  `json:"to" required:"true"`
	Entries []Entry `json:"entries,omitempty"`
}

// CustomApplication represents request parameter to update firewall.
type CustomApplication struct {
	Name            string `json:"name" required:"true"`
	Protocol        string `json:"protocol" required:"true"`
	DestinationPort string `json:"destinationPort,omitempty"`
}

// ApplicationSet represents request parameter to update firewall.
type ApplicationSet struct {
	Name         string   `json:"name" required:"true"`
	Applications []string `json:"applications" required:"true"`
}

// AddressSet represents request parameter to update firewall.
type AddressSet struct {
	Name      string   `json:"name" required:"true"`
	Addresses []string `json:"addresses" required:"true"`
}

// RoutingGroupSetting represents request parameter to update firewall.
type RoutingGroupSetting struct {
	GroupName   string       `json:"groupName" required:"true"`
	AddressSets []AddressSet `json:"addressSets" required:"true"`
}

// UpdateOpts represents options used to activate a firewall.
type UpdateOpts struct {
	Rules                []Rule                `json:"rules,omitempty"`
	CustomApplications   []CustomApplication   `json:"customApplications,omitemtpry"`
	ApplicationSets      []ApplicationSet      `json:"applicationSets,omitemtpry"`
	RoutingGroupSettings []RoutingGroupSetting `json:"routingGroupSettings,omitempty"`
}

// ToUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToUpdateMap() (map[string]interface{}, error) {
	return fic.BuildRequestBody(opts, "firewall")
}

// Update accepts a UpdateOpts struct and update a firewall
// using the values provided.
func Update(c *fic.ServiceClient, routerID, firewallID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(updateURL(c, routerID, firewallID), b, &r.Body, &fic.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
