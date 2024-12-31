package models

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const StructTag = "M3U"

// ChannelMetadata represents the channel metadata attributes in an M3U entry.
type ChannelMetadata struct {
	Censored    string `M3U:"censored"`
	ChannelID   string `M3U:"channel-id"`
	GroupID     string `M3U:"group-id"`
	GroupTitle  string `M3U:"group-title"`
	TvgCountry  string `M3U:"tvg-country"`
	TvgID       string `M3U:"tvg-id"`
	TvgLanguage string `M3U:"tvg-language"`
	TvgLogo     string `M3U:"tvg-logo"`
	TvgName     string `M3U:"tvg-name"`
	TvgShift    string `M3U:"tvg-shift"`
}

// Channel represents the metadata and URL for a channel in the M3U playlist.
type Channel struct {
	Duration int
	Metadata ChannelMetadata
	Title    string
	URL      string
}

// PlaylistHeader represents the header information for a playlist.
type PlaylistHeader struct {
	Metadata PlaylistHeaderMetadata
}

// PlaylistHeaderMetadata represents the metadata attributes for a playlist header.
type PlaylistHeaderMetadata struct {
	TvgUrl  string `M3U:"tvg-url"`
	XTvgUrl string `M3U:"x-tvg-url"`
	UrlTvg  string `M3U:"url-tvg"`
}

// M3U represents the entire M3U file, with a list of channels.
type M3U struct {
	PlaylistHeaders []PlaylistHeader
	Channels        []Channel
}

// Unmarshal parses the M3U data into the M3U struct.
func Unmarshal(data []byte, m3u *M3U) error {
	// Convert input to a string for easier parsing
	rawData := string(data)

	// Regex pattern to match #EXTINF (channel) lines and Url/Path
	channelPattern := `#EXTINF:(?P<duration>-?\d+)(?P<metadata>.*?)?,\s*(?P<title>[^,\r\n]+)\s*(?P<url>[^\r\n]+)`
	channelRegex := regexp.MustCompile(channelPattern)

	// Regex pattern to match #EXTM3U (playlist) header lines
	playlistHeaderPattern := `#EXTM3U\s+(.*?)\s*(\n|$)`
	playlistHeaderRegex := regexp.MustCompile(playlistHeaderPattern)

	// Find all matches
	channelMatches := channelRegex.FindAllStringSubmatch(rawData, -1)
	playlistHeaderMatches := playlistHeaderRegex.FindAllStringSubmatch(rawData, -1)

	for _, match := range playlistHeaderMatches {
		playlistMetadata := match[1]
		headerMetadataStruct := PlaylistHeaderMetadata{}
		if playlistMetadata != "" {
			parseMetadata(playlistMetadata, &headerMetadataStruct)
		}
		m3u.PlaylistHeaders = append(m3u.PlaylistHeaders, PlaylistHeader{
			Metadata: headerMetadataStruct,
		})
	}

	// Parse each match into the Channel struct
	for _, match := range channelMatches {
		duration := match[1]
		channelMetadata := match[2]
		title := match[3]
		url := match[4]

		// Convert duration to an integer
		var durationInt int
		_, err := fmt.Sscanf(duration, "%d", &durationInt)
		if err != nil {
			return err
		}

		// Parse the metadata into a structured Metadata object
		metadataStruct := ChannelMetadata{}
		parseMetadata(channelMetadata, &metadataStruct)

		// Add the channel to the list
		m3u.Channels = append(m3u.Channels, Channel{
			Duration: durationInt,
			Metadata: metadataStruct,
			Title:    title,
			URL:      url,
		})
	}
	return nil
}

// parseMetadata parses the metadata attributes into a Metadata struct.
func parseMetadata[T any](metadata string, metadataStruct *T) {
	// Split metadata by spaces, and then key-value by '='
	metadata = strings.TrimSpace(metadata)
	if len(metadata) > 0 {
		parts := strings.Fields(metadata)
		for _, part := range parts {
			kv := strings.SplitN(part, "=", 2)
			if len(kv) == 2 {
				// Remove quotes from values if present
				value := strings.Trim(kv[1], `"`)
				setM3UMetadataField(metadataStruct, kv[0], value)
			}
		}
	}
}

// Marshal converts the M3U struct into M3U data (as a byte slice).
func Marshal(m3u *M3U) ([]byte, error) {
	var result strings.Builder

	for _, header := range m3u.PlaylistHeaders {
		result.WriteString("#EXTM3U")
		metadataFields := getM3UMetadataFields(header.Metadata)
		for key, value := range metadataFields {
			result.WriteString(fmt.Sprintf(" %s=\"%s\"", key, value))
		}
		result.WriteString("\n")
	}

	// Write each channel entry
	for _, channel := range m3u.Channels {
		// Write #EXTINF line
		result.WriteString(fmt.Sprintf("#EXTINF:%d", channel.Duration))

		// Write metadata attributes dynamically using reflection
		metadataFields := getM3UMetadataFields(channel.Metadata)
		for key, value := range metadataFields {
			result.WriteString(fmt.Sprintf(" %s=\"%s\"", key, value))
		}

		// Write the title and URL
		result.WriteString(fmt.Sprintf(", %s\n", channel.Title))
		result.WriteString(fmt.Sprintf("%s\n", channel.URL))
	}

	return []byte(result.String()), nil
}

// setM3UMetadataField sets a field in the Metadata struct dynamically based on the custom tag.
func setM3UMetadataField(metadata any, key, value string) {
	// Use reflection to find the field that matches the M3U tag
	// Look through all fields in the Metadata struct
	v := reflect.ValueOf(metadata).Elem()
	t := reflect.TypeOf(metadata).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(StructTag)
		if tag == key {
			// Set the field value dynamically
			v.Field(i).SetString(value)
			break
		}
	}
}

// getM3UMetadataFields uses reflection to dynamically get metadata key-value pairs based on struct tags.
func getM3UMetadataFields(metadata any) map[string]string {
	metadataFields := make(map[string]string)

	// Use reflection to read struct fields and their custom tags
	v := reflect.ValueOf(metadata)
	t := reflect.TypeOf(metadata)

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(StructTag)
		value := v.Field(i).String()

		// Only include fields that have non-empty values
		if tag != "" && value != "" {
			metadataFields[tag] = value
		}
	}
	return metadataFields
}
