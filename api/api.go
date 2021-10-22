package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/xandout/assemblyai-client/transcript"
)

const (
	// BaseURLV1 is the base for AssemblyAI v1
	BaseURLV1 = "https://api.assemblyai.com/v1"
	// BaseURLV2 is the base for AssemblyAI v1
	BaseURLV2 = "https://api.assemblyai.com/v2"
)

// AssemblyAIClient is the client
type AssemblyAIClient struct {
	BaseURL    string
	HTTPClient *http.Client
	apiKey     string
}

// NewClient creates a new API client using your API key
func NewClient(apiKey string) *AssemblyAIClient {
	return &AssemblyAIClient{
		BaseURL: BaseURLV2,
		apiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *AssemblyAIClient) sendRequest(req *http.Request, v *transcript.Response) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", c.apiKey)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		if err = json.NewDecoder(res.Body).Decode(&v); err == nil {
			return errors.New("Unable to decode")
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}

// StartTranscript submits the API request to start transcribing a file
func (c *AssemblyAIClient) StartTranscript(tr transcript.Request) (transcript.Response, error) {
	ctr := transcript.Response{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/transcript", c.BaseURL), bytes.NewBuffer(tr.Bytes()))

	if err != nil {
		return ctr, err
	}

	err = c.sendRequest(req, &ctr)
	if err != nil {
		return ctr, err
	}

	return ctr, nil
}

// GetTranscript submits the API request to start transcribing a file
func (c *AssemblyAIClient) GetTranscript(transcriptID string) (transcript.Response, error) {
	ctr := transcript.Response{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/transcript/%s", c.BaseURL, transcriptID), nil)

	if err != nil {
		return ctr, err
	}

	err = c.sendRequest(req, &ctr)
	if err != nil {
		return ctr, err
	}

	return ctr, nil
}
