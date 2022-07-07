package main

import (
	"fmt"

	diagnostic "git.sr.ht/~mcritchlow/iwd-dbus/iwd"
	dbus "github.com/godbus/dbus/v5"
	iwd "github.com/shibumi/iwd"
)

const objectIwd = "net.connman.iwd"

func main() {
	conn, err := dbus.ConnectSystemBus()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// var s []string
	// err = conn.BusObject().Call("org.freedesktop.DBus.ListNames", 0).Store(&s)
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, "Failed to get list of owned names:", err)
	// 	os.Exit(1)
	// }
	//
	// fmt.Println("Currently owned names on the session bus:")
	// for _, v := range s {
	// 	fmt.Println(v)
	// }

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
