package telegraph

// PageViews represents the number of page views for a Telegraph article.
type PageViews struct {
	// Number of page views for the target page.
	Views uint `json:"views"`
}