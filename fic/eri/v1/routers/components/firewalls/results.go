package firewalls

import (
	fic "github.com/nttcom/go-fic"
	"github.com/nttcom/go-fic/pagination"
)

type commonResult struct {
	fic.Result
}

// Extract is a function that accepts a result
// and extracts a firewall resource.
func (r commonResult) Extract() (*Firewall, error) {
	var f Firewall
	err := r.ExtractInto(&f)
	return &f, err
}

// Extract interprets any commonResult as a Firewall, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "firewall")
}

// ActivateResult represents the result of a activate operation.
// Call its Extract method to interpret it as a Firewall.
type ActivateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Firewall.
type GetResult struct {
	commonResult
}

// DeativateResult represents the result of a activate operation.
// Call its Extract method to interpret it as a Firewall.
type DeactivateResult struct {
	commonResult
}

// UpdateResult represents the result of a update operation.
type UpdateResult struct {
	commonResult
}

// // CustomApplication represents the result of a update operation.
// type CustomApplication struct {
// 	Name            string `json:"name"`
// 	Protocol        string `json:"protocol"`
// 	DestinationPort string `json:"destinationPort"`
// }

// // ApplicationSet represents ApplicationSet parameter in firewall resource.
// type ApplicationSet struct {
// 	Name         string   `json:"name"`
// 	Applications []string `json:"applications"`
// }

// // AddressSet represents AddressSet parameter of RoutingGroupSetting.
// type AddressSet struct {
// 	Name      string   `json:"name"`
// 	Addresses []string `json:"addresses"`
// }

// // RoutingGroupSetting represents RoutingGroupSetting parameter of firewall.
// type RoutingGroupSetting struct {
// 	GroupName   string       `json:"groupName"`
// 	AddressSets []AddressSet `json:"addressSets"`
// }

// // Match represents match parameter of entry.
// type Match struct {
// 	SourceAddressSets      []string `json:"sourceAddressSets"`
// 	DestinationAddressSets []string `json:"destinationAddressSets"`
// 	Application            string   `json:"application"`
// }

// Entry represents each parameter of entry parmeter of rule.
// type Entry struct {
// 	Name   string `json:"name"`
// 	Match  Match  `json:"match"`
// 	Action string `json:"action"`
// }

// Rule represents each parameter of rules property in firwall resource.
// type Rule struct {
// 	From    string  `json:"from"`
// 	To      string  `json:"to"`
// 	Entries []Entry `json:"entries"`
// }

// Firewall represents firwall resource.
type Firewall struct {
	ID                   string                `json:"id"`
	TenantID             string                `json:"tenantId"`
	Redundant            bool                  `json:"redundant"`
	IsActivated          bool                  `json:"isActivated"`
	OperationStatus      string                `json:"operationStatus"`
	CustomApplications   []CustomApplication   `json:"customApplications"`
	ApplicationSets      []ApplicationSet      `json:"applicationSets"`
	RoutingGroupSettings []RoutingGroupSetting `json:"routingGroupSettings"`
	Rules                []Rule                `json:"rules"`
	UserIPAddresses      []string              `json:"userIpAddresses"`
	OperationID          string                `json:"operationId"`
}

// FirewallPage is the page returned by a pager
// when traversing over a collection of ports.
type FirewallPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of firewalls
// have reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r FirewallPage) NextPageURL() (string, error) {
	var s struct {
		Links []fic.Link `json:"firewalls_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return fic.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a FirewallPage struct is empty.
func (r FirewallPage) IsEmpty() (bool, error) {
	is, err := ExtractFirewalls(r)
	return len(is) == 0, err
}

// ExtractFirewalls accepts a Page struct,
// specifically a FirewallPage struct, and extracts the elements
// into a slice of Firewall structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractFirewalls(r pagination.Page) ([]Firewall, error) {
	var s []Firewall
	err := ExtractFirewallsInto(r, &s)
	return s, err
}

// ExtractFirewallsInto interprets the results of a single page from a List() call,
// producing a slice of Firewall entities.
func ExtractFirewallsInto(r pagination.Page, v interface{}) error {
	return r.(FirewallPage).Result.ExtractIntoSlicePtr(v, "firewalls")
}
