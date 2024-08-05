package telegraph

type editAccountInfo struct {
	// Access token of the Telegraph account.
	AccessToken string `json:"access_token"`

	// New account name.
	ShortName string `json:"short_name,omitempty"`

	// New default author name used when creating new articles.
	AuthorName string `json:"author_name,omitempty"`

	// New default profile link, opened when users click on the author's name below the title. Can be any link,
	// not necessarily to a Telegram profile or channel.
	AuthorURL string `json:"author_url,omitempty"`
}

// EditAccountInfo update information about a Telegraph account. Pass only the parameters that you want to edit. On
// success, returns an Account object with the default fields.
func (a *Account) EditAccountInfo(update Account) (*Account, error) {
	data, err := makeRequest("editAccountInfo", editAccountInfo{
		AccessToken: a.AccessToken,
		ShortName:   update.ShortName,
		AuthorName:  update.AuthorName,
		AuthorURL:   update.AuthorURL,
	})
	if err != nil {
		return nil, err
	}

	result := new(Account)
	if err = parser.Unmarshal(data, result); err != nil {
		return nil, err
	}

	return result, nil
}
