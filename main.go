package main

import (
	"log"
	"time"

	"github.com/hugolgst/rich-go/client"
	"github.com/thediveo/osrelease"
)

// Read os-release with osrelease and map
func ReadOsRelaese() (release map[string]string) {

	release = make(map[string]string)

	osrelease := osrelease.New()
	for key, value := range osrelease {
		release[key] = value
	}

	return

}

func main() {

	// Connect to DISCORD
	const APP_ID = "980223541335191622"

	err := client.Login(APP_ID)
	if err != nil {
		log.Panic("Could not connect to discord.", err)
	}

	// READ os-release
	OsRelease := ReadOsRelaese()

	// Get time
	now := time.Now()

	// Set discord activity
	for {

		err = client.SetActivity(client.Activity{
			Details:    OsRelease["PRETTY_NAME"],
			LargeImage: "tux",
			LargeText:  "Linux",
			SmallImage: OsRelease["ID"],
			SmallText:  OsRelease["PRETTY_NAME"],
			Timestamps: &client.Timestamps{
				Start: &now,
			},
		})

		if err != nil {
			panic(err)
		}

		time.Sleep(time.Duration(time.Second * 100))
	}

}
