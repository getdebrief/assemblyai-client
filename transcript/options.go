package transcript

// Option sets options on the Request
type Option func(*Request)

// WithSpeakerLabels sets speaker_labels
func WithSpeakerLabels() Option {
	return func(h *Request) {
		h.SpeakerLabels = true
	}
}

// WithAudioURL sets audio_url
func WithAudioURL(url string) Option {
	return func(h *Request) {
		h.AudioURL = url
	}
}

// WithAcousticModel sets acoustic_model
func WithAcousticModel(acousticModel string) Option {
	return func(h *Request) {
		h.AcousticModel = acousticModel
	}
}

// WithLanguageModel sets language_model
func WithLanguageModel(languageModel string) Option {
	return func(h *Request) {
		h.LanguageModel = languageModel
	}
}

// WithFormatText sets format_text
func WithFormatText() Option {
	return func(h *Request) {
		h.FormatText = true
	}
}

// WithPunctuate sets punctuate
func WithPunctuate() Option {
	return func(h *Request) {
		h.Punctuate = true
	}
}

// WithDualChannel sets dual_channel
func WithDualChannel() Option {
	return func(h *Request) {
		h.DualChannel = true
	}
}

// WithWebhookURL sets webhook_url
func WithWebhookURL(webhookURL string) Option {
	return func(h *Request) {
		h.WebhookURL = webhookURL
	}
}

// WithAudioStartFrom sets audio_start_from
func WithAudioStartFrom(startFrom int64) Option {
	return func(h *Request) {
		h.AudioStartFrom = startFrom
	}
}

// WithAudioEndAt sets audio_end_at
func WithAudioEndAt(endAt int64) Option {
	return func(h *Request) {
		h.AudioEndAt = endAt
	}
}

func WithWordBoost(words []string) Option {
	return func(h *Request) {
		h.WordBoost = words
	}
}

func WithBoostParam(boost BoostParamType) Option {
	return func(h *Request) {
		h.BoostParam = boost
	}
}

func WithAutoHighlights() Option {
	return func(h *Request) {
		h.AutoHighlights = true
	}
}
