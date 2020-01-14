package telegraph

import "errors"

// All types used in the Telegraph API responses are represented as JSON-objects. Optional fields may be not returned
// when irrelevant.
type (
	// Account represents a Telegraph account.
	Account struct {
		// Only returned by the createAccount and revokeAccessToken method. Access token of the Telegraph
		// account.
		AccessToken string `json:"access_token"`

		// URL to authorize a browser on telegra.ph and connect it to a Telegraph account. This URL is valid
		// for only one use and for 5 minutes only.
		AuthURL string `json:"auth_url,omitempty"`

		// Account name, helps users with several accounts remember which they are currently using. Displayed
		// to the user above the "Edit/Publish" button on Telegra.ph, other users don't see this name.
		ShortName string `json:"short_name"`

		// Default author name used when creating new articles.
		AuthorName string `json:"author_name"`

		// Profile link, opened when users click on the author's name below the title. Can be any link, not
		// necessarily to a Telegram profile or channel.
		AuthorURL string `json:"author_url"`

		// Number of pages belonging to the Telegraph account.
		PageCount int `json:"page_count,omitempty"`
	}

	// PageList represents a list of Telegraph articles belonging to an account. Most recently created articles
	// first.
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

		// Name of the author, displayed below the title.
		AuthorName string `json:"author_name,omitempty"`

		// Profile link, opened when users click on the author's name below the title. Can be any link, not
		// necessarily to a Telegram profile or channel.
		AuthorURL string `json:"author_url,omitempty"`

		// Image URL of the page.
		ImageURL string `json:"image_url,omitempty"`

		// Content of the page.
		Content []Node `json:"content,omitempty"`

		// Number of page views for the page.
		Views int `json:"views"`

		// Only returned if access_token passed. True, if the target Telegraph account can edit the page.
		CanEdit bool `json:"can_edit,omitempty"`
	}

	// PageViews represents the number of page views for a Telegraph article.
	PageViews struct {
		// Number of page views for the target page.
		Views int `json:"views"`
	}

	// Node is abstract object represents a DOM Node. It can be a String which represents a DOM text node or a
	// NodeElement object.
	Node interface{}

	// NodeElement represents a DOM element node.
	NodeElement struct {
		// Name of the DOM element. Available tags: a, aside, b, blockquote, br, code, em, figcaption, figure,
		// h3, h4, hr, i, iframe, img, li, ol, p, pre, s, strong, u, ul, video.
		Tag string `json:"tag"`

		// Attributes of the DOM element. Key of object represents name of attribute, value represents value
		// of attribute. Available attributes: href, src.
		Attrs map[string]string `json:"attrs,omitempty"`

		// List of child nodes for the DOM element.
		Children []Node `json:"children,omitempty"`
	}
)

const (
	// FieldShortName used as GetAccountInfo argument for getting account name.
	FieldShortName string = "short_name"

	// FieldAuthorName used as GetAccountInfo argument for getting author name.
	FieldAuthorName string = "author_name"

	// FieldAuthorURL used as GetAccountInfo argument for getting profile link.
	FieldAuthorURL string = "author_url"

	// FieldAuthURL used as GetAccountInfo argument for getting URL to authorize a browser on telegra.ph.
	FieldAuthURL string = "auth_url"

	// FieldPageCount used as GetAccountInfo argument for getting number of pages belonging to the Telegraph
	// account.
	FieldPageCount string = "page_count"
)

var (
	// ErrInvalidDataType is returned when ContentFormat function are passed a data argument of invalid type.
	ErrInvalidDataType = errors.New("invalid data type")

	// ErrNoInputData is returned when any method get nil argument.
	ErrNoInputData = errors.New("no input data")
)
