package telegraph

// PageList represents a list of Telegraph articles belonging to an account.
// Most recently created articles first.
type PageList struct {
	// Requested pages of the target Telegraph account.
	Pages []Page `json:"pages"`

	// Total number of pages belonging to the target Telegraph account.
	TotalCount uint `json:"total_count"`
}