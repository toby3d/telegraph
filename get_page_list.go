package telegraph

type getPageList struct {
	// Access token of the Telegraph account.
	AccessToken string `json:"access_token"`

	// Sequential number of the first page to be returned.
	Offset int `json:"offset,omitempty"`

	// Limits the number of pages to be retrieved.
	Limit int `json:"limit,omitempty"`
}

// GetPageList get a list of pages belonging to a Telegraph account. Returns a PageList object, sorted by most
// recently created pages first.
func (a *Account) GetPageList(offset, limit int) (*PageList, error) {
	data, err := makeRequest("getPageList", getPageList{
		AccessToken: a.AccessToken,
		Offset:      offset,
		Limit:       limit,
	})
	if err != nil {
		return nil, err
	}

	result := new(PageList)
	if err = parser.Unmarshal(data, result); err != nil {
		return nil, err
	}

	return result, nil
}
