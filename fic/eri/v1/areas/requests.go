package areas

import (
	fic "github.com/nttcom/go-fic"
	"github.com/nttcom/go-fic/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToAreaListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the area attributes you want to see returned.
type ListOpts struct {
}

// ToAreaListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToAreaListQuery() (string, error) {
	q, err := fic.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over
// a collection of areas.
// It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *fic.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToAreaListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return AreaPage{pagination.LinkedPageBase{PageResult: r}}
	})
}
