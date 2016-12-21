package telegraph

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"strconv"
)

/*
// CreatePage create a new Telegraph page. On success, returns a Page object.
func (account *Account) CreatePage(page *Page, returnContent bool) (*Page, error) {
	var args fasthttp.Args

	args.Add("access_token", account.AccessToken)
	// Required. Access token of the Telegraph account.

	args.Add("title", page.Title)
	// Required. Page title.

	args.Add("author_name", page.AuthorName)
	// Author name, displayed below the article's title.

	args.Add("author_url", page.AuthorURL)
	// Profile link, opened when users click on the author's name below the
	// title. Can be any link, not necessarily to a Telegram profile or channel.

	args.Add("content", page.Content)
	// Required. Content of the page.

	args.Add("return_content", strconv.FormatBool(returnContent))
	// If true, a content field will be returned in the Page object.

	url := fmt.Sprintf(APIEndpoint, "createPage")
	body, err := request(nil, url, &args)
	if err != nil {
		return nil, err
	}

	var resp Page
	if err := json.Unmarshal(body.Result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
*/

/*
// EditPage edit an existing Telegraph page. On success, returns a Page object.
func (account *Account) EditPage(update *Page, returnContent bool) (*Page, error) {
	var args fasthttp.Args

	args.Add("access_token", account.AccessToken)
	// Required. Access token of the Telegraph account.

	args.Add("title", update.Title)
	// Required. Page title.

	args.Add("content", update.Content.Data)
	// Required. Content of the page.

	args.Add("author_name", update.AuthorName)
	// Author name, displayed below the article's title.

	args.Add("author_url", update.AuthorURL)
	// Profile link, opened when users click on the author's name below the
	// title. Can be any link, not necessarily to a Telegram profile or channel.

	args.Add("return_content", strconv.FormatBool(returnContent))
	// If true, a content field will be returned in the Page object.

	url := fmt.Sprintf(PathEndpoint, "editPage", update.Path)
	body, err := request(nil, url, &args)
	if err != nil {
		return nil, err
	}

	var resp Page
	if err := json.Unmarshal(body.Result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
*/

// GetPage get a Telegraph page. Returns a Page object on success.
func GetPage(path string, returnContent bool) (*Page, error) {
	var args fasthttp.Args

	args.Add("return_content", strconv.FormatBool(returnContent))
	// If true, content field will be returned in Page object.

	url := fmt.Sprintf(PathEndpoint, "getPage", path)
	body, err := request(nil, url, &args)
	if err != nil {
		return nil, err
	}

	var resp Page
	if err := json.Unmarshal(body.Result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetPageList get a list of pages belonging to a Telegraph account. Returns a
// PageList object, sorted by most recently created pages first.
func (account *Account) GetPageList(offset int, limit int) (*PageList, error) {
	var args fasthttp.Args

	args.Add("access_token", account.AccessToken)
	// Required. Access token of the Telegraph account.

	args.Add("offset", strconv.Itoa(offset))
	// Sequential number of the first page to be returned.

	args.Add("limit", strconv.Itoa(limit))
	// Limits the number of pages to be retrieved.

	url := fmt.Sprintf(APIEndpoint, "getPageList")
	body, err := request(nil, url, &args)
	if err != nil {
		return nil, err
	}

	var resp PageList
	if err := json.Unmarshal(body.Result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetViews get the number of views for a Telegraph article. By default, the
// total number of page views will be returned. Returns a PageViews object on
// success.
func GetViews(path string, year int, month int, day int, hour int) (*PageViews, error) {
	var args fasthttp.Args

	if year >= 2000 && year <= 2100 {
		args.Add("year", strconv.Itoa(year))
		// Required if month is passed. If passed, the number of page views for
		// the requested year will be returned.
	}

	if month > 0 {
		args.Add("month", strconv.Itoa(month))
		// Required if day is passed. If passed, the number of page views for
		// the requested month will be returned.
	}

	if day > 0 {
		args.Add("day", strconv.Itoa(day))
		// Required if hour is passed. If passed, the number of page views for
		// the requested day will be returned.
	}

	if hour > -1 {
		args.Add("hour", strconv.Itoa(hour))
		// If passed, the number of page views for the requested hour will be
		// returned.
	}

	url := fmt.Sprintf(PathEndpoint, "getViews", path)
	body, err := request(nil, url, &args)
	if err != nil {
		return nil, err
	}

	var resp PageViews
	if err := json.Unmarshal(body.Result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
