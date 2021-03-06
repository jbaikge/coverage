package rss

import (
	"encoding/xml"
	"github.com/300brand/coverage/feed/parser/charset"
	"github.com/300brand/coverage/feed/parser/decoder"
	"github.com/300brand/coverage/feed/parser/time"
)

type Doc struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	// Category       []string  `xml:"category"`       // Optional. Defines one or more categories for the feed
	// Cloud          string    `xml:"cloud"`          // Optional. Register processes to be notified immediately of updates of the feed
	// Copyright      string    `xml:"copyright"`      // Optional. Notifies about copyrighted material
	// Description    string    `xml:"description"`    // Required. Describes the channel
	// Docs           string    `xml:"docs"`           // Optional. Specifies an URL to the documentation of the format used in the feed
	// Generator      string    `xml:"generator"`      // Optional. Specifies the program used to generate the feed
	// Image          string    `xml:"image"`          // Optional. Allows an image to be displayed when aggregators present a feed
	Item []Item `xml:"item"` // Optional. Stories within the feed
	// Language       string    `xml:"language"`       // Optional. Specifies the language the feed is written in
	// LastBuildDate  time.Time `xml:"lastBuildDate"`  // Optional. Defines the last-modified date of the content of the feed
	// Link           string    `xml:"link"`           // Required. Defines the hyperlink to the channel
	// ManagingEditor string    `xml:"managingEditor"` // Optional. Defines the e-mail address to the editor of the content of the feed
	// PubDate        time.Time `xml:"pubDate"`        // Optional. Defines the last publication date for the content of the feed
	// Rating         string    `xml:"rating"`         // Optional. The PICS rating of the feed
	// SkipDays       string    `xml:"skipDays"`       // Optional. Specifies the days where aggregators should skip updating the feed
	// SkipHours      string    `xml:"skipHours"`      // Optional. Specifies the hours where aggregators should skip updating the feed
	// TextInput      string    `xml:"textInput"`      // Optional. Specifies a text input field that should be displayed with the feed
	Title          string    `xml:"title"`          // Required. Defines the title of the channel
	// Ttl            string    `xml:"ttl"`            // Optional. Specifies the number of minutes the feed can stay cached before refreshing it from the source
	// WebMaster      string    `xml:"webMaster"`      // Optional. Defines the e-mail address to the webmaster of the feed
}

type Item struct {
	// Author      string    `xml:"author"`      // Optional. Specifies the e-mail address to the author of the item
	// Category    string    `xml:"category"`    // Optional. Defines one or more categories the item belongs to
	// Comments    string    `xml:"comments"`    // Optional. Allows an item to link to comments about that item
	// Description string    `xml:"description"` // Required. Describes the item
	// Enclosure   string    `xml:"enclosure"`   // Optional. Allows a media file to be included with the item
	// Guid        string    `xml:"guid"`        // Optional. Defines a unique identifier for the item
	Link    string    `xml:"link"`    // Required. Defines the hyperlink to the item
	PubDate time.Time `xml:"pubDate"` // Optional. Defines the last-publication date for the item
	// Source      string    `xml:"source"`      // Optional. Specifies a third-party source for the item
	Title string `xml:"title"` // Required. Defines the title of the item
}

func init() {
	decoder.RegisterDecoder("rss", &Doc{})
}

func (doc Doc) New() decoder.Decoder {
	return &Doc{}
}

func (doc *Doc) Decode(data []byte) error {
	return charset.TryAll(doc, data)
}
