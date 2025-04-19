package models

type ChannelBuilderInterface interface {
	NewChannelBuilder() *ChannelBuilder
	SetDuration(duration int) *ChannelBuilder
	SetMetadata(metadata *ChannelMetadata) *ChannelBuilder
	SetTitle(title string) *ChannelBuilder
	SetURL(url string) *ChannelBuilder
	Build() *Channel
}

type ChannelBuilder struct {
	channel *Channel
}

func NewChannelBuilder() *ChannelBuilder {
	return &ChannelBuilder{
		channel: &Channel{
			Metadata: ChannelMetadata{},
		},
	}
}

func (b *ChannelBuilder) SetDuration(duration int) *ChannelBuilder {
	b.channel.Duration = duration
	return b
}

func (b *ChannelBuilder) SetMetadata(metadata *ChannelMetadata) *ChannelBuilder {
	b.channel.Metadata = *metadata
	return b
}

func (b *ChannelBuilder) SetTitle(title string) *ChannelBuilder {
	b.channel.Title = title
	return b
}

func (b *ChannelBuilder) SetURL(url string) *ChannelBuilder {
	b.channel.URL = url
	return b
}

func (b *ChannelBuilder) Build() *Channel {
	return b.channel
}
