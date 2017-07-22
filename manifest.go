package globals

type Author struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Site string `json:"site"`
}
type Manifest struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	Version string `json:"version"`
	ReleaseDate string `json:"release_date"`
	Description string `json:"description"`
	Author Author `json:"author"`
	HomePage string `json:"home_page"`
	MinimumCmsVersion string `json:"minimum_cms_version"`
}
