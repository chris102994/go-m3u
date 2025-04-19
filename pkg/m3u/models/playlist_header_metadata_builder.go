package models

type PlaylistHeaderMetadataBuilderInterface interface {
	NewPlaylistHeaderMetadataBuilder() *PlaylistHeaderMetadataBuilder
	SetTvgUrl(tvgUrl string) *PlaylistHeaderMetadataBuilder
	SetXTvgUrl(xTvgUrl string) *PlaylistHeaderMetadataBuilder
	SetUrlTvg(urlTvg string) *PlaylistHeaderMetadataBuilder
}

type PlaylistHeaderMetadataBuilder struct {
	playlistHeaderMetadata *PlaylistHeaderMetadata
}

func NewPlaylistHeaderMetadataBuilder() *PlaylistHeaderMetadataBuilder {
	return &PlaylistHeaderMetadataBuilder{
		playlistHeaderMetadata: &PlaylistHeaderMetadata{},
	}
}

func (b *PlaylistHeaderMetadataBuilder) SetTvgUrl(tvgUrl string) *PlaylistHeaderMetadataBuilder {
	b.playlistHeaderMetadata.TvgUrl = tvgUrl
	return b
}

func (b *PlaylistHeaderMetadataBuilder) SetXTvgUrl(xTvgUrl string) *PlaylistHeaderMetadataBuilder {
	b.playlistHeaderMetadata.XTvgUrl = xTvgUrl
	return b
}

func (b *PlaylistHeaderMetadataBuilder) SetUrlTvg(urlTvg string) *PlaylistHeaderMetadataBuilder {
	b.playlistHeaderMetadata.UrlTvg = urlTvg
	return b
}

func (b *PlaylistHeaderMetadataBuilder) Build() *PlaylistHeaderMetadata {
	return b.playlistHeaderMetadata
}
