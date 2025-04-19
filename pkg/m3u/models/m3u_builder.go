package models

type M3UBuilderInterface interface {
	NewM3UBuilder() *M3UBuilder
	AddPlaylistHeader(header *PlaylistHeader) *M3UBuilder
	AddChannel(channel *Channel) *M3UBuilder
	Build() *M3U
}

type M3UBuilder struct {
	m3u *M3U
}

func NewM3UBuilder() *M3UBuilder {
	return &M3UBuilder{
		m3u: &M3U{
			PlaylistHeaders: make([]PlaylistHeader, 0),
			Channels:        make([]Channel, 0),
		},
	}
}

func (b *M3UBuilder) AddPlaylistHeader(header *PlaylistHeader) *M3UBuilder {
	b.m3u.PlaylistHeaders = append(b.m3u.PlaylistHeaders, *header)
	return b
}

func (b *M3UBuilder) AddChannel(channel *Channel) *M3UBuilder {
	b.m3u.Channels = append(b.m3u.Channels, *channel)
	return b
}

func (b *M3UBuilder) Build() *M3U {
	return b.m3u
}
