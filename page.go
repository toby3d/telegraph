package telegraph

// Page represents a page on Telegraph.
type Page struct {
	// URL of the page.
	URL *URL `json:"url,omitempty"`

	// Optional. Profile link, opened when users click on the author's name
	// below the title. Can be any link, not necessarily to a Telegram
	// profile or channel.
	AuthorURL *URL `json:"author_url,omitempty"`

	// Optional. Image URL of the page.
	ImageURL *URL `json:"image_url,omitempty"`

	// Path to the page.
	Path string `json:"path,omitempty"`

	// Title of the page.
	Title string `json:"title,omitempty"`

	// Description of the page.
	Description string `json:"description,omitempty"`

	// Optional. Name of the author, displayed below the title.
	AuthorName string `json:"author_name,omitempty"`

	// Optional. [Content] of the page.
	Content []Node `json:"content,omitempty"`

	// Number of page views for the page.
	Views uint `json:"views,omitempty"`

	// Optional. Only returned if access_token passed. True, if the target
	// Telegraph account can edit the page.
	CanEdit bool `json:"can_edit,omitempty"`
}