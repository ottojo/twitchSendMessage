package main

import (
	"bufio"
	"flag"
	"github.com/gempir/go-twitch-irc"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var channel string
var oauth string
var user string

func init() {
	flag.StringVar(&oauth, "o", "oauth:XXX...", "oauth token")
	flag.StringVar(&channel, "c", "teamspatzenhirn", "Target Channel name")
	flag.StringVar(&user, "u", "teamspatzenhirn", "Twitch User Name")
	flag.Parse()
}

func main() {
	client := twitch.NewClient(user, oauth)
	client.Join(channel)

	client.OnConnect(func() {
		reader := bufio.NewReader(os.Stdin)
		for true {
			text, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			text = strings.TrimSpace(text)
			client.Say(channel, text)
			time.Sleep(200 * time.Millisecond)
		}
		err := client.Disconnect()
		if err != nil {
			log.Println(err)
		}
	})

	_ = client.Connect()
}
