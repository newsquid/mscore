package mscore

import (
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

/*
PaginationParameters contains parameters used in pagination, a page number and how many items per page
*/
type PaginationParameters struct {
	Page         int
	ItemsPerPage int
}

/*
Offset calculates the current item offset using the page number and items per page
*/
func (p PaginationParameters) Offset() int {
	return (p.Page - 1) * p.ItemsPerPage
}

type PeriodParameters struct {
	From time.Time
	To   time.Time
}

/*
IdParameter contains an id
*/
type IdParameter struct {
	Id int64
}

/*
NameParameter contains a name
Same idea as IdParameter, but a string
*/
type NameParameter struct {
	Name string
}

/*
SearchParameter cotains a search string if success is true
*/
type SearchParameter struct {
	Search  string
	Success bool
}

/*
SearchTerm turns the Search string into the SQL search term: %Search%
*/
func (s SearchParameter) SearchTerm() string {
	return fmt.Sprintf("%%%s%%", s.Search)
}

/*
Pagination retrieves the pagination parameters 'page' and 'items per page' and
validates them. An error is thrown if the parameters are missing or invalid
*/
func Pagination(w http.ResponseWriter, query url.Values, m martini.Context) {
	page, pageErr := strconv.Atoi(query.Get("page"))
	itemsPerPage, itemsPerPageErr := strconv.Atoi(query.Get("items_per_page"))

	if pageErr != nil || itemsPerPageErr != nil || page < 1 || itemsPerPage < 1 {
		m.Map(PaginationParameters{Page: 1, ItemsPerPage: 9})
		return
	}

	m.Map(PaginationParameters{Page: page, ItemsPerPage: itemsPerPage})
}

func DatePeriod(w http.ResponseWriter, query url.Values, m martini.Context) {
	from := query.Get("fromdate")
	to := query.Get("todate")

	if "" == to || "" == to {
		http.Error(w, "fromdate or todate missing", http.StatusBadRequest)
		return
	}

	from_date, from_date_err := time.Parse(time.RFC3339, from)
	to_date, to_date_err := time.Parse(time.RFC3339, to)

	if nil != from_date_err || nil != to_date_err {
		http.Error(w, "fromdate or todate could not be parsed", 422)
	}

	m.Map(PeriodParameters{From: from_date, To: to_date})
}

/*
ResourceId retrieves the 'id' parameter and validates it. An error is thrown if
the parameter is missing or invalid
*/
func ResourceId(w http.ResponseWriter, params martini.Params, m martini.Context) {
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil || id < 1 {
		http.Error(w, "Unprocessable Entity", 422)
	}

	m.Map(IdParameter{Id: id})
}

/*
ResourceName retrieves the 'name' parameter and validates it. An error is thrown if
the parameter is missing or invalid
*/
func ResourceName(name string) func(http.ResponseWriter, martini.Params, martini.Context) {
	return func(w http.ResponseWriter, params martini.Params, m martini.Context) {
		name_value := params[name]

		if "" == name_value {
			http.Error(w, "Unprocessable Entity", 422)
		}

		m.Map(NameParameter{Name: name_value})
	}
}

/*
SearchTerm retrieves the 'search' query parameter and reports whether or not the
parameter is present
*/
func SearchTerm(query url.Values, m martini.Context) {
	p := SearchParameter{}
	p.Search = query.Get("search")
	p.Success = p.Search != ""

	m.Map(p)
}
