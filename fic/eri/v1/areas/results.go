package areas

import (
	"github.com/nttcom/go-fic"
	"github.com/nttcom/go-fic/pagination"
)

type commonResult struct {
	fic.Result
}

// Extract is a function that accepts a result
// and extracts a area resource.
func (r commonResult) Extract() (*Area, error) {
	var a Area
	err := r.ExtractInto(&a)
	return &a, err
}

// Extract interprets any commonResult as an Area, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "area")
}

// PoAreart represents area resource.
type Area struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Number   int    `json:"number"`
	IsPublic bool   `json:"isPublic"`
	Enabled  bool   `json:"enabled"`
}

// AreaPage is the page returned by a pager
// when traversing over a collection of areas.
type AreaPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of areas
// have reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r AreaPage) NextPageURL() (string, error) {
	var s struct {
		Links []fic.Link `json:"areas_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return fic.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a AreaPage struct is empty.
func (r AreaPage) IsEmpty() (bool, error) {
	is, err := ExtractAreas(r)
	return len(is) == 0, err
}

// ExtractAreas accepts a Page struct,
// specifically a AreaPage struct, and extracts the elements
// into a slice of Area structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractAreas(r pagination.Page) ([]Area, error) {
	var s []Area
	err := ExtractAreasInto(r, &s)
	return s, err
}

// ExtractAreasInto interprets the results of a single page from a List() call,
// producing a slice of Area entities.
func ExtractAreasInto(r pagination.Page, v interface{}) error {
	return r.(AreaPage).Result.ExtractIntoSlicePtr(v, "areas")
}
