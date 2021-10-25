package assemblyai

// Body is the data assembly calls the webhook with
type Body struct {
	TranscriptID string `json:"transcript_id,omitempty"`
	Status       string `json:"status,omitempty"`
}
