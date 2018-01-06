package globals

type FieldOption struct {
	Key   string
	Value string
}

type Field struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Caption         string `json:"caption"`
	Placeholder     string `json:"placeholder"`
	Multiple        bool `json:"multiple"`
	Class           string `json:"class"`
	Type            string `json:"type"`
	Values          []string `json:"values"`
	Options         []FieldOption `json:"options"`
	Title           string `json:"title"`
	Required        bool `json:"required"`
	OnClick         string `json:"onclick"`
	OnChange        string `json:"onchange"`
	OndDblClick     string `json:"onddblclick"`
	OnFocus         string `json:"onfocus"`
	OnFocusOut      string `json:"onfocusout"`
	OnHover         string `json:"onhover"`
	OnKeyDown       string `json:"onkeydown"`
	OnKeyPress      string `json:"onkeypress"`
	OnKeyUp         string `json:"onkeyup"`
	OnMouseDown     string `json:"onmousedown"`
	OnMouseEnter    string `json:"onmouseenter"`
	OnMouseLeave    string `json:"onmouseleave"`
	OnMouseMove     string `json:"onmousemove"`
	OnMouseOut      string `json:"onmouseout"`
	OnMouseOver     string `json:"onmouseover"`
	OnMouseUp       string `json:"onmouseup"`
	AllowExtensions string `json:"allow_extensions"`
	InfoblockID     int64 `json:"infoblock_id"`
	InfoblockCode   string `json:"infoblock_code"`
	Disabled        bool `json:"disabled"`
	Exclude         string `json:"exclude"`
	Sort            int `json:"sort"`
	Checked         bool `json:"checked"`
	Autocomplete    bool `json:"autocomplete"`
	EditorMode      string `json:"editor_mode"`
	EditorTheme     string `json:"editor_theme"`
	CountryID       int64 `json:"country_id"`
	RegionID        int64 `json:"region_id"`
}

type Fieldser interface {
	StringsToField(params []string) Field
	Build(field Field, db DataBaser, out Outputer, options Optionser, files Fileser) string
}