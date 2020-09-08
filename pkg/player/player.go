package player

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/godbus/dbus/v5"
)

const (
	mprisPath = "/org/mpris/MediaPlayer2"

	member           = "org.mpris.MediaPlayer2.Player"
	playMessage      = member + ".Play"
	pauseMessage     = member + ".Pause"
	playPauseMessage = member + ".PlayPause"
	previousMessage  = member + ".Previous"
	nextMessage      = member + ".Next"
	metadataMessage  = member + ".Metadata"
	statusMessage    = member + ".PlaybackStatus"
)

type Metadata struct {
	Artist      []string `mpris:"xesam:artist"`
	Title       string   `mpris:"xesam:title"`
	Album       string   `mpris:"xesam:album"`
	AlbumArtist []string `mpris:"xesam:albumArtist"`
	AutoRating  float64  `mpris:"xesam:autoRating"`
	DiskNumber  int32    `mpris:"xesam:discNumber"`
	TrackNumber int32    `mpris:"xesam:trackNumber"`
	URL         string   `mpris:"xesam:url"`
	TrackID     string   `mpris:"mpris:trackid"`
	Length      uint64   `mpris:"mpris:length"`
}

// parseMetadata returns a parsed Metadata struct.
func parseMetadata(variant dbus.Variant) *Metadata {
	metadataMap := variant.Value().(map[string]dbus.Variant)
	metadataStruct := new(Metadata)

	valueOf := reflect.ValueOf(metadataStruct).Elem()
	typeOf := reflect.TypeOf(metadataStruct).Elem()

	for key, val := range metadataMap {
		for i := 0; i < typeOf.NumField(); i++ {
			field := typeOf.Field(i)
			if field.Tag.Get("mpris") == key {
				field := valueOf.Field(i)
				field.Set(reflect.ValueOf(val.Value()))
			}
		}
	}

	return metadataStruct
}

type Player struct {
	name string
	conn *dbus.Conn
	obj  dbus.BusObject
}

func New(conn *dbus.Conn, name string) Player {
	p := Player{
		name: name,
		conn: conn,
		obj:  conn.Object(name, mprisPath),
	}

	return p
}

func (p Player) Metadata() (*Metadata, error) {
	prop, err := p.obj.GetProperty(metadataMessage)
	if err != nil {
		return nil, err
	}

	metadata := parseMetadata(prop)

	return metadata, nil
}

func (p Player) String() string {
	metadata, err := p.Metadata()
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}

	artists := strings.Join(metadata.Artist, ", ")
	title := metadata.Title

	return fmt.Sprintf("%s - %s", artists, title)
}

func (p Player) Play() error {
	c := p.obj.Call(playMessage, 0)
	if c.Err != nil {
		return c.Err
	}

	return nil
}

func (p Player) Pause() error {
	c := p.obj.Call(pauseMessage, 0)
	if c.Err != nil {
		return c.Err
	}

	return nil
}

func (p Player) PlayPause() error {
	c := p.obj.Call(playPauseMessage, 0)
	if c.Err != nil {
		return c.Err
	}

	return nil
}

func (p Player) Next() error {
	c := p.obj.Call(nextMessage, 0)
	if c.Err != nil {
		return c.Err
	}

	return nil
}

func (p Player) Previous() error {
	c := p.obj.Call(previousMessage, 0)
	if c.Err != nil {
		return c.Err
	}

	return nil
}

func (p Player) Status() (string, error) {
	prop, err := p.obj.GetProperty(statusMessage)
	if err != nil {
		return "Unknown", err
	}

	return strings.Trim(prop.String(), "\""), nil
}
