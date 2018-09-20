package scrapbox

type Line struct {
	ID      string `json:"id"`
	Text    string `json:"text"`
	UserID  string `json:"userId"`
	Created int    `json:"created"`
	Updated int    `json:"updated"`
}
