package iwd

import (
	"fmt"
	"os"

	"github.com/godbus/dbus/v5"
	"github.com/shibumi/iwd"
)

const (
	objectIwd                     = "net.connman.iwd"
	objectStation                 = "net.connman.iwd.Station"
	objectStationDiagnostic       = "net.connman.iwd.StationDiagnostic"
	objectStationDiagnosticMethod = "net.connman.iwd.StationDiagnostic.GetDiagnostics"
)

// StationDiagnostic refers to net.connman.iwd.StationDiagnostic
type StationDiagnostic struct {
	Path             dbus.ObjectPath
	ConnectedNetwork dbus.ObjectPath
	Scanning         bool
	State            string
}

func Frequency(conn *dbus.Conn, station iwd.Station) error {
	var objects map[string]dbus.Variant
	err := conn.Object(objectIwd, station.Path).Call(objectStationDiagnosticMethod, 0).Store(&objects)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to introspect StationDiagnostic", err)
		os.Exit(1)
	}
	// output looks like:
	// k ConnectedBss v "9c:a5:70:68:45:c6"
	//
	// k RxBitrate v @u 6500
	//
	// k RxMode v "802.11ac"
	//
	// k RxMCS v @y 0x7
	//
	// k RSSI v @n -54
	//
	// k AverageRSSI v @n -53
	//
	// k TxMCS v @y 0x6
	//
	// k Security v "WPA2-Personal"
	//
	// k TxMode v "802.11ac"
	//
	// k TxBitrate v @u 5266
	//
	// k Frequency v @u 5580
	frequency := objects["Frequency"].Value().(uint32)
	fmt.Printf("Frequency is %d\n", frequency)
	if frequency > 5000 {
		fmt.Println("Good news, iwd is using 5GHz")
	} else {
		fmt.Println("Poop")
		// TODO we want to check for Frequency == 5580, or > 5000. If not, restart? put in a loop until it is?
		// This should be in a different func so Diagnostic can be called (and renamed..) in an easy loop
		// 1. disconnect
		// 2. reconnect
		// 3. check frequency again
	}
	return nil
}
