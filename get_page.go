package telegraph

import (
	gopath "path"
)

type getPage struct {
	// Path to the Telegraph page (in the format Title-12-31, i.e. everything that comes after http://telegra.ph/).
	Path string `json:"path"`

	// If true, content field will be returned in Page object.
	ReturnContent bool `json:"return_content,omitempty"`
}

// GetPage get a Telegraph page. Returns a Page object on success.
func GetPage(path string, returnContent bool) (*Page, error) {
	data, err := makeRequest(gopath.Join("getPage", path), getPage{
		Path:          path,
		ReturnContent: returnContent,
	})
	if err != nil {
		return nil, err
	}

	result := new(Page)
	if err = parser.Unmarshal(data, result); err != nil {
		return nil, err
	}

	return result, nil
}
