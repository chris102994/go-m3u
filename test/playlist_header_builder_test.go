package test

import (
	m3u "github.com/chris102994/go-m3u/pkg/m3u/models"
	"testing"
)

func TestPlaylistHeaderBuilder(t *testing.T) {
	var metadata = &m3u.PlaylistHeaderMetadata{
		TvgUrl:  "http://example.com/tvg",
		XTvgUrl: "http://example.com/xtvg",
		UrlTvg:  "http://example.com/urltvg",
	}

	playlistHeader := m3u.NewPlaylistHeaderBuilder().
		SetMetadata(metadata).
		Build()

	if playlistHeader.Metadata.TvgUrl != metadata.TvgUrl {
		t.Errorf("expected TvgUrl to be '%s', got '%s'", metadata.TvgUrl, playlistHeader.Metadata.TvgUrl)
	}
}
