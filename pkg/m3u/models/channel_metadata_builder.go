package models

type ChannelMetadataBuilderInterface interface {
	NewChannelMetadataBuilder() *ChannelMetadataBuilder
	SetCensored(censored string) *ChannelMetadataBuilder
	SetChannelID(channelID string) *ChannelMetadataBuilder
	SetGroupID(groupID string) *ChannelMetadataBuilder
	SetGroupTitle(groupTitle string) *ChannelMetadataBuilder
	SetTvgCountry(tvgCountry string) *ChannelMetadataBuilder
	SetTvgID(tvgID string) *ChannelMetadataBuilder
	SetTvgLanguage(tvgLanguage string) *ChannelMetadataBuilder
	SetTvgLogo(tvgLogo string) *ChannelMetadataBuilder
	SetTvgName(tvgName string) *ChannelMetadataBuilder
	SetTvgShift(tvgShift string) *ChannelMetadataBuilder
	Build() *ChannelMetadata
}

type ChannelMetadataBuilder struct {
	channelMetadata *ChannelMetadata
}

func NewChannelMetadataBuilder() *ChannelMetadataBuilder {
	return &ChannelMetadataBuilder{
		channelMetadata: &ChannelMetadata{},
	}
}

func (b *ChannelMetadataBuilder) SetCensored(censored string) *ChannelMetadataBuilder {
	b.channelMetadata.Censored = censored
	return b
}

func (b *ChannelMetadataBuilder) SetChannelID(channelID string) *ChannelMetadataBuilder {
	b.channelMetadata.ChannelID = channelID
	return b
}

func (b *ChannelMetadataBuilder) SetGroupID(groupID string) *ChannelMetadataBuilder {
	b.channelMetadata.GroupID = groupID
	return b
}

func (b *ChannelMetadataBuilder) SetGroupTitle(groupTitle string) *ChannelMetadataBuilder {
	b.channelMetadata.GroupTitle = groupTitle
	return b
}

func (b *ChannelMetadataBuilder) SetTvgCountry(tvgCountry string) *ChannelMetadataBuilder {
	b.channelMetadata.TvgCountry = tvgCountry
	return b
}

func (b *ChannelMetadataBuilder) SetTvgID(tvgID string) *ChannelMetadataBuilder {
	b.channelMetadata.TvgID = tvgID
	return b
}

func (b *ChannelMetadataBuilder) SetTvgLanguage(tvgLanguage string) *ChannelMetadataBuilder {
	b.channelMetadata.TvgLanguage = tvgLanguage
	return b
}

func (b *ChannelMetadataBuilder) SetTvgLogo(tvgLogo string) *ChannelMetadataBuilder {
	b.channelMetadata.TvgLogo = tvgLogo
	return b
}

func (b *ChannelMetadataBuilder) SetTvgName(tvgName string) *ChannelMetadataBuilder {
	b.channelMetadata.TvgName = tvgName
	return b
}

func (b *ChannelMetadataBuilder) SetTvgShift(tvgShift string) *ChannelMetadataBuilder {
	b.channelMetadata.TvgShift = tvgShift
	return b
}

func (b *ChannelMetadataBuilder) Build() *ChannelMetadata {
	return b.channelMetadata
}
