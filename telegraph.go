package telegraph

import (
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
	"golang.org/x/net/html"
)

// Telegraph constants
const (
	// APIEndpoint should be presented in this for all queries to the Telegraph
	// API must be served over HTTPS.
	APIEndpoint = "https://api.telegra.ph/%s"

	// PathEndpoint used if a path parameter is present.
	PathEndpoint = "https://api.telegra.ph/%s/%s"
)

type (
	// Account represents a Telegraph account.
	Account struct {
		// Optional. Only returned by the createAccount and revokeAccessToken
		// method. Access token of the Telegraph account.
		AccessToken string `json:"access_token"`

		// Optional. URL to authorize a browser on telegra.ph and connect it to
		// a Telegraph account. This URL is valid for only one use and for 5
		// minutes only.
		AuthURL string `json:"auth_url"`

		// Account name, helps users with several accounts remember which they
		// are currently using. Displayed to the user above the "Edit/Publish"
		// button on Telegra.ph, other users don't see this name.
		ShortName string `json:"short_name"`

		// Default author name used when creating new articles.
		AuthorName string `json:"author_name"`

		// Profile link, opened when users click on the author's name below the
		// title. Can be any link, not necessarily to a Telegram profile or
		// channel.
		AuthorURL string `json:"author_url"`

		// Optional. Number of pages belonging to the Telegraph account.
		PageCount int `json:"page_count"`
	}

	// PageList represents a list of Telegraph articles belonging to an account.
	// Most recently created articles first.
	PageList struct {
		// Total number of pages belonging to the target Telegraph account.
		TotalCount int `json:"total_count"`

		// Requested pages of the target Telegraph account.
		Pages []Page `json:"pages"`
	}

	// Page represents a page on Telegraph.
	Page struct {
		// Path to the page.
		Path string `json:"path"`

		// URL of the page.
		URL string `json:"url"`

		// Title of the page.
		Title string `json:"title"`

		// Description of the page.
		Description string `json:"description"`

		// Optional. Name of the author, displayed below the title.
		AuthorName string `json:"author_name"`

		// Optional. Profile link, opened when users click on the author's name
		// below the title. Can be any link, not necessarily to a Telegram
		// profile or channel.
		AuthorURL string `json:"author_url"`

		// Optional. Image URL of the page.
		ImageURL string `json:"image_url"`

		// Optional. Content of the page.
		Content []Node `json:"content"`

		// Number of page views for the page.
		Views int `json:"views"`

		// Optional. Only returned if access_token passed. True, if the target
		// Telegraph account can edit the page.
		CanEdit bool `json:"can_edit"`
	}

	// PageViews represents the number of page views for a Telegraph article.
	PageViews struct {
		// Number of page views for the target page.
		Views int `json:"views"`
	}

	// Node is abstract object represents a DOM Node. It can be a String which
	// represents a DOM text node or a NodeElement object.
	Node interface{}

	// NodeElement represents a DOM element node.
	NodeElement struct {
		// Name of the DOM element. Available tags: a, aside, b, blockquote, br,
		// code, em, figcaption, figure, h3, h4, hr, i, iframe, img, li, ol, p,
		// pre, s, strong, u, ul, video.
		Tag string `json:"tag"`

		// Optional. Attributes of the DOM element. Key of object represents
		// name of attribute, value represents value of attribute. Available
		// attributes: href, src.
		Attrs []html.Attribute `json:"attrs"`

		// Optional. List of child nodes for the DOM element.
		Children []Node `json:"children"`
	}

	// Response represents a response from the Telegram API with the result
	// stored raw. If ok equals true, the request was successful, and the result
	// of the query can be found in the result field. In case of an unsuccessful
	// request, ok equals false, and the error is explained in the error field
	// (e.g. SHORT_NAME_REQUIRED).
	Response struct {
		Ok     bool            `json:"ok"`
		Error  string          `json:"error"`
		Result json.RawMessage `json:"result"`
	}
)

func request(dst []byte, url string, args *fasthttp.Args) (*Response, error) {
	/*
		if args != nil {
			url += "?" + args.String()
		}

		var req fasthttp.Request
		req.Header.SetUserAgent("telegraph")
		req.Header.SetMethod("POST")
		req.Header.SetContentType("application/json; charset=utf-8")
		req.SetRequestURI(url)
		if dst != nil {
			req.SetBody(dst)
		}

		var res fasthttp.Response
		err := fasthttp.Do(&req, &res)
		if err != nil {
			return nil, err
		}
	*/

	_, res, err := fasthttp.Post(dst, url, args)
	if err != nil {
		return nil, err
	}

	var tResp Response
	if err := json.Unmarshal(res, &tResp); err != nil {
		return nil, err
	}

	if !tResp.Ok {
		return nil, errors.New(tResp.Error)
	}

	return &tResp, nil
}
