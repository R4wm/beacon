package main

import (
	"flag"
	"log"

	"github.com/r4wm/beacon"
)

func main() {
	// Args
	var hostname = flag.String("hostname", "", "hostname or ip of where to send beacon")
	flag.Parse()

	// Get public IP addr
	ip := beacon.GetOutboundIP()
	if ip == "" {
		log.Fatal("No ip recieved from beacon.GetOutboundIP()")
	}

	// Get Geo
	city, err := beacon.GetGeo(ip)
	if err != nil {
		log.Fatal(err)
	}

	// Send beacon to base
	err = beacon.SendBeacon(city, *hostname)
	if err != nil {
		log.Fatal(err)
	}
}
