package main

import (
	m3u "github.com/chris102994/go-m3u/pkg/m3u/models"
)

func main() {
	channel1 := m3u.NewChannelBuilder().
		SetDuration(120).
		SetTitle("Channel 1").
		SetURL("http://example.com/channel1").
		SetMetadata(m3u.NewChannelMetadataBuilder().
			SetTvgID("1").
			SetTvgName("Channel 1").
			SetTvgLogo("http://example.com/logo1.png").
			SetGroupTitle("Group 1").
			Build()).
		Build()

	channel2 := m3u.NewChannelBuilder().
		SetDuration(150).
		SetTitle("Channel 2").
		SetURL("http://example.com/channel2").
		SetMetadata(m3u.NewChannelMetadataBuilder().
			SetTvgID("2").
			SetTvgName("Channel 2").
			SetTvgLogo("http://example.com/logo2.png").
			SetGroupTitle("Group 2").
			Build()).
		Build()

	m3uOutput := m3u.NewM3UBuilder().
		AddChannel(channel1).
		AddChannel(channel2).
		AddPlaylistHeader(m3u.NewPlaylistHeaderBuilder().
			SetMetadata(m3u.NewPlaylistHeaderMetadataBuilder().
				SetTvgUrl("http://example.com/tvg").
				SetXTvgUrl("http://example.com/xtvg").
				SetUrlTvg("http://example.com/urltvg").
				Build()).
			Build()).
		Build()

	output, err := m3u.Marshal(m3uOutput)
	if err != nil {
		panic(err)
	}
	println(string(output))

}
