package main

import (
	"fmt"
	"os"

	"github.com/godbus/dbus/v5"
	"github.com/shanedabes/polyphon/pkg/dbusutils"
	"github.com/shanedabes/polyphon/pkg/player"
	"github.com/shanedabes/polyphon/pkg/utils"
)

func main() {
	conn, err := dbus.SessionBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to to session bus:", err)
	}
	defer conn.Close()

	players, err := dbusutils.GetPlayerNames(conn.BusObject())
	if err != nil {
		panic(err)
	}

	filters := []string{
		"org.mpris.MediaPlayer2.chromium",
		"org.mpris.MediaPlayer2.mpv",
	}
	players = utils.FilterStrings(players, filters)

	for _, v := range players {
		fmt.Println(v)

		p := player.New(conn, v)

		fmt.Println(p.String())

		status, err := p.Status()
		if err != nil {
			continue
		}

		fmt.Println(status)
	}
}
