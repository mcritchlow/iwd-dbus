package main

import (
	"fmt"

	diagnostic "git.sr.ht/~mcritchlow/iwd-dbus/iwd"
	dbus "github.com/godbus/dbus/v5"
	iwd "github.com/shibumi/iwd"
)

func main() {
	conn, err := dbus.ConnectSystemBus()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	iwdConnection := iwd.New(conn)
	for _, v := range iwdConnection.Stations {
		for _, n := range iwdConnection.Networks {
			if n.Connected {
				fmt.Printf("Connected Network is %s\n", n.Name)
			}
		}
		diagnostic.Frequency(conn, v)
	}
}
