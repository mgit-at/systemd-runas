package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/coreos/go-systemd/dbus"
)

func main() {
	conn, err := dbus.NewSystemdConnection()
	if err != nil {
		log.Fatal("connecting to systemd: ", err)
	}
	defer conn.Close()

	log.Printf("cmds: (%d) %v\n", len(os.Args), os.Args)
	if len(os.Args) != 3 {
		log.Fatal("invalid number of arguments: systemd-runas <slice> <pid>")
	}

	slice := os.Args[1]
	pid, err := strconv.ParseUint(os.Args[2], 10, 32)
	if err != nil {
		log.Fatal("parsing pid: ", err)
	}
	pid32 := uint32(pid)
	props := []dbus.Property{
		dbus.PropSlice(slice + ".slice"),
		dbus.PropPids(pid32),
	}
	target := fmt.Sprintf("%s-%d.scope", slice, pid32)

	// Start the unit
	_, err = conn.StartTransientUnit(target, "replace", props, nil)
	if err != nil {
		log.Fatal(err)
	}
}
