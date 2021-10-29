# AssemblyAI Golang Client (WIP)

This is a golang client for the [AssemblyAI API](https://docs.assemblyai.com/overview/getting-started).

> This library is new and mostly built for my own purposes.  Use at your own risk or submit a PR to improve :)

## Version history

### v0.0.6
* fixed bug where `AutoHighlight` objects did not grab timestamps

### v0.0.5
* fixed bug where `iab_categories` timestamp object was of the wrong type
* renamed webhook response type from `Body` to `WebhookResponse` for clarity

### v0.0.4
* reorganized to put everything under one package for ease of import/naming, and easier use alongside other transcription services.

### v0.0.3
* Typed all untyped responses
* added functionality for `word_boost`, `iab_categories` and `auto_highlights`

### v0.0.2
* Renamed `Transcript` to `Request`
* Added `Reader` method on `Request` and `Response` to aid downstream users

### v0.0.1
* Initial version

## Features

* Start transcription with `audio_url`
* Get transcript as JSON

* API options covered
    - `speaker_labels`
    - `audio_url`
    - `acoustic_model`
    - `language_model`
    - `format_text`
    - `punctuate`
    - `dual_channel`
    - `webhook_url`
    - `audio_start_from`
    - `audio_end_at`
    - `word_boost`
    - `boost_param`
    - `auto_highlights`
    - `iab_categories`

## Usage

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/getdebrief/assemblyai-client/api"
	"github.com/getdebrief/assemblyai-client/transcript"
)

func main() {
	api := api.NewClient("API KEY")

	audioURL, err := api.UploadFile("FILEPATH")
	if err != nil {
		log.Fatal(err)
	}

	tr := transcript.NewRequest(
		transcript.WithAudioURL(audioURL),
		transcript.WithSpeakerLabels(),
		transcript.WithAutoHighlights(),
		transcript.WithWordBoost([]string{"Word"}),
		transcript.WithBoostParam(transcript.BoostParamHigh),
	)

	started, err := api.StartTranscript(*tr)
	if err != nil {
		log.Fatal(err)
	}

	var finished transcript.Response
	processing := true
	for processing {
		time.Sleep(30 * time.Second)
		finished, err = api.GetTranscript(started.ID)
		if err != nil {
			log.Fatal(err)
		} else if finished.Status != "processing" {
			processing = false
		}
	}

	fmt.Println(string(finished.Bytes()))
}

```

```json
{
	"acoustic_model": "assemblyai_default",
	"audio_duration": 12.225,
	"audio_url": "https://you-audio-file-host.com/your-file.mp3",
	"confidence": 0.913333333333333,
	"format_text": true,
	"id": "long-uuid",
	"language_model": "assemblyai_default",
	"punctuate": true,
	"status": "completed",
	"text": "Quick brown Fox jumped over the fence.",
	"utterances": [{
		"confidence": 0.9328571428571429,
		"end": 6141,
		"speaker": "A",
		"start": 2241,
		"text": "Quick brown Fox jumped over the fence.",
		"words": [{
			"confidence": 0.97,
			"end": 2970,
			"speaker": "A",
			"start": 2610,
			"text": "Quick"
		}, {
			"confidence": 0.92,
			"end": 3390,
			"speaker": "A",
			"start": 3060,
			"text": "brown"
		}, {
			"confidence": 0.98,
			"end": 3840,
			"speaker": "A",
			"start": 3420,
			"text": "Fox"
		}, {
			"confidence": 0.94,
			"end": 4530,
			"speaker": "A",
			"start": 3810,
			"text": "jumped"
		}, {
			"confidence": 0.92,
			"end": 5010,
			"speaker": "A",
			"start": 4500,
			"text": "over"
		}, {
			"confidence": 0.96,
			"end": 5250,
			"speaker": "A",
			"start": 4950,
			"text": "the"
		}, {
			"confidence": 0.84,
			"end": 5520,
			"speaker": "A",
			"start": 5190,
			"text": "fence."
		}]
	}],
	"words": [{
		"confidence": 0.97,
		"end": 2970,
		"speaker": "A",
		"start": 2610,
		"text": "Quick"
	}, {
		"confidence": 0.92,
		"end": 3390,
		"speaker": "A",
		"start": 3060,
		"text": "brown"
	}, {
		"confidence": 0.98,
		"end": 3840,
		"speaker": "A",
		"start": 3420,
		"text": "Fox"
	}, {
		"confidence": 0.94,
		"end": 4530,
		"speaker": "A",
		"start": 3810,
		"text": "jumped"
	}, {
		"confidence": 0.92,
		"end": 5010,
		"speaker": "A",
		"start": 4500,
		"text": "over"
	}, {
		"confidence": 0.96,
		"end": 5250,
		"speaker": "A",
		"start": 4950,
		"text": "the"
	}, {
		"confidence": 0.84,
		"end": 5520,
		"speaker": "A",
		"start": 5190,
		"text": "fence."
	}]
}
```