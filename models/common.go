package models

import (
	"github.com/koolalex/mdblog/config"
	"time"
)

type JsonTime time.Time

type Tag string

type Markdown struct {
	Meta
	Description string   `json:"description"`
	Date        JsonTime `json:"date"` //
	Path        string   `json:"path"`
}

type Meta struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	Tags     []Tag  `json:"tags"`
	Author   string `json:"author"`
}

type MarkdownDetails struct {
	Markdown
	Body string
}

type MarkdownList []Markdown

type MarkdownPagination struct {
	Markdowns   MarkdownList
	Total       int
	CurrentPage int
	PageNumber  []int
}

type Categories []Category

type Category struct {
	Name             string
	Path             string
	Number           int
	MarkdownFileList MarkdownList
}

/*JsonTime*/
func (t *JsonTime) UnmarshalJSON(b []byte) error {
	date, err := time.ParseInLocation(`"`+config.Cfg.TimeLayout+`"`, string(b), time.Local)
	if err != nil {
		return nil
	}
	*t = JsonTime(date)
	return nil
}

func (t JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(t.Format(`"` + config.Cfg.TimeLayout + `"`)), nil
}

func (t JsonTime) Format(layout string) string {
	return time.Time(t).Format(layout)
}

/*MarkdownList*/
func (m MarkdownList) Len() int { return len(m) }

func (m MarkdownList) Less(i, j int) bool { return time.Time(m[i].Date).After(time.Time(m[j].Date)) }

func (m MarkdownList) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
