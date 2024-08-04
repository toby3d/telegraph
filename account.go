package telegraph

// Account represents a Telegraph account.
type Account struct {
	// Profile link, opened when users click on the author's name below the
	// title. Can be any link, not necessarily to a Telegram profile or
	// channel.
	AuthorURL URL `json:"author_url"`

	// Optional. URL to authorize a browser on [telegra.ph] and connect it to
	// a Telegraph account. This URL is valid for only one use and for 5
	// minutes only.
	//
	// [telegra.ph]: https://telegra.ph/
	AuthURL *URL `json:"auth_url,omitempty"`

	ShortName ShortName `json:"short_name"`

	// Default author name used when creating new articles.
	AuthorName AuthorName `json:"author_name"`

	// Optional. Only returned by the [createAccount] and
	// [revokeAccessToken] method. Access token of the Telegraph account.
	AccessToken string `json:"access_token,omitempty"`

	// Optional. Number of pages belonging to the Telegraph account.
	PageCount uint `json:"page_count,omitempty"`
}