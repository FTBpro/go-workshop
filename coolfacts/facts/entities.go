package facts

type MF struct {
	Id            string   `json:"id"`
	Url           string   `json:"url"`
	FactId        string   `json:"factId"`
	Headline      string   `json:"headline"`
	ShortHeadline string   `json:"shortHeadline"`
	FactText      string   `json:"fact"`
	FullStoryUrl  string   `json:"fullStoryUrl"`
	Tags          []string `json:"tags"`
	PrimaryImage  string   `json:"primaryImage"`
	ImageCredit   string   `json:"imageCredit"`
}

type Fact struct {
	Image       string
	Url         string
	Description string
}