package telegraph

import (
	"fmt"
	"strconv"
	"strings"
)

type AccountField struct{ accountField string }

var (
	FieldAuthorName AccountField = AccountField{"author_name"}
	FieldAuthorURL  AccountField = AccountField{"author_url"}
	FieldAuthURL    AccountField = AccountField{"auth_url"}
	FieldPageCount  AccountField = AccountField{"page_count"}
	FieldShortName  AccountField = AccountField{"short_name"}
)

var stringsAccountFields map[string]AccountField = map[string]AccountField{
	FieldAuthorName.accountField: FieldAuthorName,
	FieldAuthorURL.accountField:  FieldAuthorURL,
	FieldAuthURL.accountField:    FieldAuthURL,
	FieldPageCount.accountField:  FieldPageCount,
	FieldShortName.accountField:  FieldShortName,
}

func (af *AccountField) UnmarshalJSON(v []byte) error {
	unquoted, err := strconv.Unquote(string(v))
	if err != nil {
		return fmt.Errorf("AccountField: UnmarshalJSON: cannot unquote value '%s': %w", string(v), err)
	}

	if accountField, ok := stringsAccountFields[strings.ToLower(unquoted)]; ok {
		*af = accountField
	}

	return nil
}

func (af AccountField) MarshalJSON() ([]byte, error) {
	if af.accountField != "" {
		return []byte(strconv.Quote(af.accountField)), nil
	}

	return nil, nil
}

func (af AccountField) String() string {
	return af.accountField
}

func (af AccountField) GoString() string {
	if af.accountField == "" {
		return "telegraph.AccountField(und)"
	}

	return "telegraph.AccountField(" + af.accountField + ")"
}