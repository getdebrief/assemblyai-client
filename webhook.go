package assemblyai

// WebhookResponse is the data assembly calls the webhook with
type WebhookResponse struct {
	TranscriptID string `json:"transcript_id,omitempty"`
	Status       string `json:"status,omitempty"`
}
