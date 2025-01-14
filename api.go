package assemblyai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	// BaseURLV1 is the base for AssemblyAI v1
	BaseURLV1 = "https://api.assemblyai.com/v1"
	// BaseURLV2 is the base for AssemblyAI v1
	BaseURLV2 = "https://api.assemblyai.com/v2"
)

// AssemblyAIClient is the client
type assemblyAIClient struct {
	BaseURL    string
	HTTPClient *http.Client
	apiKey     string
}

type AssemblyAIClient interface {
	StartTranscript(tr Request) (Response, error)
	GetTranscript(transcriptID string) (Response, error)
	UploadFile(filepath string) (string, error)
}

type UploadResponse struct {
	UploadUrl string `json:"upload_url"`
}

// NewClient creates a new API client using your API key
func NewClient(apiKey string) AssemblyAIClient {
	return &assemblyAIClient{
		BaseURL: BaseURLV2,
		apiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *assemblyAIClient) sendRequest(req *http.Request, v interface{}) error {
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
			return fmt.Errorf("unable to decode %d", res.StatusCode)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}

// StartTranscript submits the API request to start transcribing a file
func (c *assemblyAIClient) StartTranscript(tr Request) (Response, error) {
	ctr := Response{}
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
func (c *assemblyAIClient) GetTranscript(transcriptID string) (Response, error) {
	ctr := Response{}
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

// UploadFile uploads a local file to Assembly.ai
func (c *assemblyAIClient) UploadFile(filepath string) (string, error) {
	resp := UploadResponse{}
	fileBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	// pass it into where nil is rn
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/upload", c.BaseURL), bytes.NewBuffer(fileBytes))
	if err != nil {
		return "", err
	}

	err = c.sendRequest(req, &resp)
	if err != nil {
		return "", err
	}

	return resp.UploadUrl, nil
}
