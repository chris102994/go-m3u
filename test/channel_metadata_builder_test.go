package test

import (
	m3u "github.com/chris102994/go-m3u/pkg/m3u/models"
	"testing"
)

func TestChannelMetadataBuilder(t *testing.T) {
	var censor = "Censored"
	var channelID = "12345"
	var groupID = "67890"
	var groupTitle = "Test Group"
	var tvgCountry = "US"
	var tvgID = "tvg123"
	var tvgLanguage = "en"
	var tvgLogo = "http://example.com/logo.png"
	var tvgName = "Test TV"
	var tvgShift = "0"

	channelMetadata := m3u.NewChannelMetadataBuilder().
		SetCensored(censor).
		SetChannelID(channelID).
		SetGroupID(groupID).
		SetGroupTitle(groupTitle).
		SetTvgCountry(tvgCountry).
		SetTvgID(tvgID).
		SetTvgLanguage(tvgLanguage).
		SetTvgLogo(tvgLogo).
		SetTvgName(tvgName).
		SetTvgShift(tvgShift).
		Build()

	if channelMetadata.Censored != censor {
		t.Errorf("expected Censored to be '%s', got '%s'", censor, channelMetadata.Censored)
	}

	if channelMetadata.ChannelID != channelID {
		t.Errorf("expected ChannelID to be '%s', got '%s'", channelID, channelMetadata.ChannelID)
	}

	if channelMetadata.GroupID != groupID {
		t.Errorf("expected GroupID to be '%s', got '%s'", groupID, channelMetadata.GroupID)
	}

	if channelMetadata.GroupTitle != groupTitle {
		t.Errorf("expected GroupTitle to be '%s', got '%s'", groupTitle, channelMetadata.GroupTitle)
	}

	if channelMetadata.TvgCountry != tvgCountry {
		t.Errorf("expected TvgCountry to be '%s', got '%s'", tvgCountry, channelMetadata.TvgCountry)
	}

	if channelMetadata.TvgID != tvgID {
		t.Errorf("expected TvgID to be '%s', got '%s'", tvgID, channelMetadata.TvgID)
	}

	if channelMetadata.TvgLanguage != tvgLanguage {
		t.Errorf("expected TvgLanguage to be '%s', got '%s'", tvgLanguage, channelMetadata.TvgLanguage)
	}

	if channelMetadata.TvgLogo != tvgLogo {
		t.Errorf("expected TvgLogo to be '%s', got '%s'", tvgLogo, channelMetadata.TvgLogo)
	}

	if channelMetadata.TvgName != tvgName {
		t.Errorf("expected TvgName to be '%s', got '%s'", tvgName, channelMetadata.TvgName)
	}

	if channelMetadata.TvgShift != tvgShift {
		t.Errorf("expected TvgShift to be '%s', got '%s'", tvgShift, channelMetadata.TvgShift)
	}
}
