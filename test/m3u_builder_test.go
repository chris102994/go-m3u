package test

import (
	m3u "github.com/chris102994/go-m3u/pkg/m3u/models"
	"testing"
)

func TestM3UBuilder(t *testing.T) {
	var PlaylistHeader = &m3u.PlaylistHeader{
		Metadata: m3u.PlaylistHeaderMetadata{
			TvgUrl:  "http://example.com/tvg",
			XTvgUrl: "http://example.com/xtvg",
			UrlTvg:  "http://example.com/urltvg",
		},
	}

	var Channel = &m3u.Channel{
		Duration: 120,
		Metadata: m3u.ChannelMetadata{
			ChannelID: "12345",
		},
		Title: "Test Channel",
		URL:   "http://example.com/stream",
	}

	m3uBuilder := m3u.NewM3UBuilder().
		AddPlaylistHeader(PlaylistHeader).
		AddChannel(Channel).
		Build()

	if len(m3uBuilder.PlaylistHeaders) != 1 {
		t.Errorf("expected PlaylistHeaders to have length 1, got %d", len(m3uBuilder.PlaylistHeaders))
	} else {
		if m3uBuilder.PlaylistHeaders[0] != *PlaylistHeader {
			t.Errorf("expected PlaylistHeader to be '%v', got '%v'", PlaylistHeader, m3uBuilder.PlaylistHeaders[0])
		}
	}

	if len(m3uBuilder.Channels) != 1 {
		t.Errorf("expected Channels to have length 1, got %d", len(m3uBuilder.Channels))
	} else {
		if m3uBuilder.Channels[0] != *Channel {
			t.Errorf("expected Channel to be '%v', got '%v'", Channel, m3uBuilder.Channels[0])
		}
	}
}
