package test

import (
	m3u "github.com/chris102994/go-m3u/pkg/m3u/models"
	"testing"
)

func TestPlaylistHeaderMetadataBuilder(t *testing.T) {
	var tvgURL = "http://example.com/tvg"
	var xTvgURL = "http://example.com/xtvg"
	var urlTvg = "http://example.com/urltvg"

	playlistHeaderMetadata := m3u.NewPlaylistHeaderMetadataBuilder().
		SetTvgUrl(tvgURL).
		SetXTvgUrl(xTvgURL).
		SetUrlTvg(urlTvg).
		Build()

	if playlistHeaderMetadata.TvgUrl != tvgURL {
		t.Errorf("expected TvgUrl to be '%s', got '%s'", tvgURL, playlistHeaderMetadata.TvgUrl)
	}

	if playlistHeaderMetadata.XTvgUrl != xTvgURL {
		t.Errorf("expected XTvgUrl to be '%s', got '%s'", xTvgURL, playlistHeaderMetadata.XTvgUrl)
	}

	if playlistHeaderMetadata.UrlTvg != urlTvg {
		t.Errorf("expected UrlTvg to be '%s', got '%s'", urlTvg, playlistHeaderMetadata.UrlTvg)
	}
}
