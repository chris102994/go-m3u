package test

import (
	m3u "github.com/chris102994/go-m3u/pkg/m3u/models"
	"testing"
)

func TestChannelBuilder(t *testing.T) {
	var duration = 120
	var metadata = &m3u.ChannelMetadata{
		ChannelID: "12345",
	}
	var title = "Test Channel"
	var url = "http://example.com/stream"

	channel := m3u.NewChannelBuilder().
		SetDuration(duration).
		SetMetadata(metadata).
		SetTitle(title).
		SetURL(url).
		Build()

	if channel.Duration != duration {
		t.Errorf("expected Duration to be '%d', got '%d'", duration, channel.Duration)
	}

	if channel.Metadata.ChannelID != metadata.ChannelID {
		t.Errorf("expected ChannelID to be '%s', got '%s'", metadata.ChannelID, channel.Metadata.ChannelID)
	}

	if channel.Title != title {
		t.Errorf("expected Title to be '%s', got '%s'", title, channel.Title)
	}

	if channel.URL != url {
		t.Errorf("expected URL to be '%s', got '%s'", url, channel.URL)
	}
}
