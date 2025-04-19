package models

type PlaylistHeaderBuilderInterface interface {
	NewPlaylistHeaderBuilder() *PlaylistHeaderBuilder
	SetMetadata(metadata *PlaylistHeaderMetadata) *PlaylistHeaderBuilder
	Build() *PlaylistHeader
}

type PlaylistHeaderBuilder struct {
	playlistHeader *PlaylistHeader
}

func NewPlaylistHeaderBuilder() *PlaylistHeaderBuilder {
	return &PlaylistHeaderBuilder{
		playlistHeader: &PlaylistHeader{},
	}
}

func (b *PlaylistHeaderBuilder) SetMetadata(metadata *PlaylistHeaderMetadata) *PlaylistHeaderBuilder {
	b.playlistHeader.Metadata = *metadata
	return b
}

func (b *PlaylistHeaderBuilder) Build() *PlaylistHeader {
	return b.playlistHeader
}
