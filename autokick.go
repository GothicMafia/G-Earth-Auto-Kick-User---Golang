package main

import (
	"log"
	"strings"
	"xabbo.b7c.io/goearth"
	"xabbo.b7c.io/goearth/in"
	"xabbo.b7c.io/goearth/out"
)

var ext = goearth.NewExt(goearth.ExtInfo{
	Title:       "AutoKick",
	Description: "Automatically kicks specific users from the room and sends a message",
	Version:     "1.2",
	Author:      "Cai",
})

func main() {
	targetUsernames := []string{"Cai", "Charl.", "sinner", "Abuse.", "Botnet", "443"} // Add multiple usernames here

	ext.Intercept(in.Users).With(func(e *goearth.Intercept) {
		username := e.Packet.ReadString() // Read the username

		// Check if the detected username ends with any of the target usernames
		for _, targetUsername := range targetUsernames {
			if strings.HasSuffix(username, targetUsername) {
				log.Printf("Kicking user: %s", username)
				// Send the kick command with only the target username (e.g., "Cai")
				ext.Send(out.KickUser, targetUsername)

				// Send a message to the room notifying that the user has been kicked
				message := targetUsername + " has been kicked from the room."
				ext.Send(out.Shout, message, 0, 0)
				break
			}
		}
	})

	ext.Run()
}
