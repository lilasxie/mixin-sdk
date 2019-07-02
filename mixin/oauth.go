package mixin

// DApp mixin
type DApp struct {
	Type         string   `json:"type"`
	AppID        string   `json:"app_id"`
	AppNumber    string   `json:"app_number"`
	RedirectURI  string   `json:"redirect_uri"`
	HomeURI      string   `json:"home_uri"`
	Name         string   `json:"name"`
	IconURL      string   `json:"icon_url"`
	Description  string   `json:"description"`
	Capabilities []string `json:"capabilities"`
	AppSecret    string   `json:"app_secret"`
	CreatorID    string   `json:"creator_id"`
}

// Authorization authorization info
type Authorization struct {
	Type              string   `json:"type"`
	AuthorizationID   string   `json:"authorization_id"`
	AuthorizationCode string   `json:"authorization_code"`
	Scopes            []string `json:"scopes"`
	App               *DApp    `json:"app"`
}
