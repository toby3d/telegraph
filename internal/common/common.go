package common

const HeaderContentType string = "Content-Type"

const (
	MIMEApplicationJSON            string = "application/json"
	MIMEApplicationJSONCharsetUTF8 string = MIMEApplicationJSON + "; " + charsetUTF8
)

const charsetUTF8 string = "charset=utf-8"