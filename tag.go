package telegraph

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/net/html/atom"
)

type Tag struct{ tag atom.Atom }

var (
	A          Tag = Tag{atom.A}          // "a"
	Aside      Tag = Tag{atom.Aside}      // "aside"
	B          Tag = Tag{atom.B}          // "b"
	Blockquote Tag = Tag{atom.Blockquote} // "blockquote"
	Br         Tag = Tag{atom.Br}         // "br"
	Code       Tag = Tag{atom.Code}       // "code"
	Em         Tag = Tag{atom.Em}         // "em"
	Figcaption Tag = Tag{atom.Figcaption} // "figcaption"
	Figure     Tag = Tag{atom.Figure}     // "figure"
	H3         Tag = Tag{atom.H3}         // "h3"
	H4         Tag = Tag{atom.H4}         // "h4"
	Hr         Tag = Tag{atom.Hr}         // "hr"
	I          Tag = Tag{atom.I}          // "i"
	Iframe     Tag = Tag{atom.Iframe}     // "iframe"
	Img        Tag = Tag{atom.Img}        // "img"
	Li         Tag = Tag{atom.Li}         // "li"
	Ol         Tag = Tag{atom.Ol}         // "ol"
	P          Tag = Tag{atom.P}          // "p"
	Pre        Tag = Tag{atom.Pre}        // "pre"
	S          Tag = Tag{atom.S}          // "s"
	Strong     Tag = Tag{atom.Strong}     // "strong"
	U          Tag = Tag{atom.U}          // "u"
	Ul         Tag = Tag{atom.Ul}         // "ul"
	Video      Tag = Tag{atom.Video}      // "video"
)

var ErrTag error = errors.New("unsupported Tag")

func NewTag(t atom.Atom) (Tag, error) {
	switch t {
	default:
		return Tag{}, fmt.Errorf("Tag: %w: want 'a', 'aside', 'b', 'blockquote', 'br', 'code', 'em', "+
			"'figcaption', 'figure', 'h3', 'h4', 'hr', 'i', 'iframe', 'img', 'li', 'ol', 'p', 'pre', "+
			"'s', 'strong', 'u', 'ul' or 'video', got '%s'", ErrTag, t.String())
	case atom.A, atom.Aside, atom.B, atom.Blockquote, atom.Br, atom.Code, atom.Em, atom.Figcaption, atom.Figure,
		atom.H3, atom.H4, atom.Hr, atom.I, atom.Iframe, atom.Img, atom.Li, atom.Ol, atom.P, atom.Pre, atom.S,
		atom.Strong, atom.U, atom.Ul, atom.Video:
	}

	return Tag{tag: t}, nil
}

func (t *Tag) UnmarshalJSON(v []byte) error {
	unquoted, err := strconv.Unquote(string(v))
	if err != nil {
		return fmt.Errorf("Tag: UnmarshalJSON: cannot unquote value '%s': %w", string(v), err)
	}

	parsed, err := NewTag(atom.Lookup([]byte(strings.ToLower(unquoted))))
	if err != nil {
		return fmt.Errorf("Tag: UnmarshalJSON: cannot parse value '%s': %w", string(v), err)
	}

	*t = parsed

	return nil
}

func (t Tag) MarshalJSON() ([]byte, error) {
	if t.tag != 0 {
		return []byte(strconv.Quote(t.tag.String())), nil
	}

	return nil, nil
}

func (t Tag) Atom() atom.Atom {
	return t.tag
}

func (t Tag) String() string {
	return t.tag.String()
}

func (t Tag) GoString() string {
	return "telegraph.Tag(" + t.String() + ")"
}