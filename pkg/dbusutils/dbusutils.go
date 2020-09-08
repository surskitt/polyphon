package dbusutils

import (
	"strings"

	"github.com/godbus/dbus/v5"
)

const (
	mprisPrefix = "org.mpris.MediaPlayer2."
)

func GetPlayerNames(bo dbus.BusObject) (players []string, err error) {
	names := []string{}

	err = bo.Call("org.freedesktop.DBus.ListNames", 0).Store(&names)
	if err != nil {
		return
	}

	for _, name := range names {
		if strings.HasPrefix(name, mprisPrefix) {
			players = append(players, name)
		}
	}

	return
}
