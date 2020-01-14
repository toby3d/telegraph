package telegraph

type createAccount struct {
	// Account name, helps users with several accounts remember which they are currently using. Displayed to the
	// user above the "Edit/Publish" button on Telegra.ph, other users don't see this name.
	ShortName string `json:"short_name"`

	// Default author name used when creating new articles.
	AuthorName string `json:"author_name,omitempty"`

	// Default profile link, opened when users click on the author's name below the title. Can be any link, not
	// necessarily to a Telegram profile or channel.
	AuthorURL string `json:"author_url,omitempty"`
}

// CreateAccount create a new Telegraph account. Most users only need one account, but this can be useful for channel
// administrators who would like to keep individual author names and profile links for each of their channels. On
// success, returns an Account object with the regular fields and an additional access_token field.
func CreateAccount(account Account) (*Account, error) {
	data, err := makeRequest("createAccount", createAccount{
		ShortName:  account.ShortName,
		AuthorName: account.AuthorName,
		AuthorURL:  account.AuthorURL,
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
