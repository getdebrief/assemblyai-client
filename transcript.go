package assemblyai

import (
	"bytes"
	"encoding/json"
	"log"
)

type BoostParamType string

const (
	BoostParamLow     BoostParamType = "low"
	BoostParamDefault BoostParamType = "default"
	BoostParamHigh    BoostParamType = "high"
)

// Request is the JSON body to pass to the API
type Request struct {
	AcousticModel   string         `json:"acoustic_model,omitempty"`
	AudioEndAt      int64          `json:"audio_end_at,omitempty"`
	AudioStartFrom  int64          `json:"audio_start_from,omitempty"`
	AudioURL        string         `json:"audio_url,omitempty"`
	DualChannel     bool           `json:"dual_channel,omitempty"`
	FormatText      bool           `json:"format_text,omitempty"`
	LanguageModel   string         `json:"language_model,omitempty"`
	Punctuate       bool           `json:"punctuate,omitempty"`
	SpeakerLabels   bool           `json:"speaker_labels,omitempty"`
	WebhookURL      string         `json:"webhook_url,omitempty"`
	AutoHighlights  bool           `json:"auto_highlights,omitempty"`
	WordBoost       []string       `json:"word_boost,omitempty"`
	BoostParam      BoostParamType `json:"boost_param,omitempty"`
	IABCategories   bool           `json:"iab_categories,omitempty"`
	EntityDetection bool           `json:"entity_detection,omitempty"`
	AutoChapters    bool           `json:"auto_chapters,omitempty"`
}

type Timestamp struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

type IABStatus string

const (
	IABSuccess     IABStatus = "success"
	IABUnavailable IABStatus = "unavailable"
)

type IABLabel struct {
	Relevance float64 `json:"relevance"`
	Label     string  `json:"label"`
}

type IABResult struct {
	Text      string     `json:"text"`
	Timestamp Timestamp  `json:"timestamp"`
	Labels    []IABLabel `json:"labels"`
}

type IABCatResponse struct {
	Status  IABStatus   `json:"status"`
	Results []IABResult `json:"results"`
	Summary interface{} `json:"summary"` // Each category is its own item, so typing this is infeasable.
}

type AutoHighlight struct {
	Count      int         `json:"count"`
	Rank       float64     `json:"rank"`
	Text       string      `json:"text"`
	Timestamps []Timestamp `json:"timestamps"`
}

type AutoHighlightsResponse struct {
	Status  string          `json:"status"`
	Results []AutoHighlight `json:"results"`
}

type Word struct {
	Text       string  `json:"text"`
	Start      int64   `json:"start"`
	End        int64   `json:"end"`
	Confidence float64 `json:"confidence"`
	Speaker    string  `json:"speaker"`
}

type Utterance struct {
	Start      int64   `json:"start"`
	End        int64   `json:"end"`
	Confidence float64 `json:"confidence"`
	Speaker    string  `json:"speaker"`
	Text       string  `json:"text"`
	Words      []Word  `json:"words"`
}

type EntityResponse struct {
	Type    string `json:"entity_type"`
	Text    string `json:"text"`
	StartMS int    `json:"start"`
	EndMS   int    `json:"end"`
}

type AutoChapterResponse struct {
	Summary  string `json:"summary"`
	Headline string `json:"headline"`
	StartMS  int    `json:"start"`
	EndMS    int    `json:"end"`
}

// Response is the API response
type Response struct {
	AcousticModel        string                 `json:"acoustic_model,omitempty"`
	AudioDuration        float64                `json:"audio_duration,omitempty"`
	AudioURL             string                 `json:"audio_url,omitempty"`
	Confidence           float64                `json:"confidence,omitempty"`
	DualChannel          bool                   `json:"dual_channel,omitempty"`
	Error                string                 `json:"error,omitempty"`
	FormatText           bool                   `json:"format_text,omitempty"`
	ID                   string                 `json:"id,omitempty"`
	LanguageModel        string                 `json:"language_model,omitempty"`
	Punctuate            bool                   `json:"punctuate,omitempty"`
	Status               string                 `json:"status,omitempty"`
	Text                 string                 `json:"text,omitempty"`
	Utterances           []Utterance            `json:"utterances,omitempty"`
	WebhookStatusCode    int64                  `json:"webhook_status_code,omitempty"`
	WebhookURL           string                 `json:"webhook_url,omitempty"`
	Words                []Word                 `json:"words,omitempty"`
	AutoHighlightsResult AutoHighlightsResponse `json:"auto_highlights_result,omitempty"`
	IABCategoriesResult  IABCatResponse         `json:"iab_categories_result,omitempty"`
	Entities             []EntityResponse       `json:"entities,omitempty"`
	Chapters             []AutoChapterResponse  `json:"chapters,omitempty"`
}

// Reader returns a bytes.Reader from Request
func (t *Request) Reader() *bytes.Reader {
	return bytes.NewReader(t.Bytes())
}

// Reader returns a bytes.Reader from Response
func (t *Response) Reader() *bytes.Reader {
	return bytes.NewReader(t.Bytes())
}

// Bytes returns the Bytes from Response
func (t *Response) Bytes() []byte {
	b, err := json.Marshal(t)
	if err != nil {
		log.Print(err)
	}
	return b

}

// Bytes returns the Bytes from Request
func (t *Request) Bytes() []byte {
	b, err := json.Marshal(t)
	if err != nil {
		log.Print(err)
	}
	return b

}

// NewRequest creates a new transcript request
func NewRequest(opts ...Option) *Request {
	tr := &Request{}
	for _, opt := range opts {
		// Call the option giving the instantiated
		// *Request as the argument
		opt(tr)
	}

	return tr
}
