package main

import (
	"fmt"
	"log"
	"os"

	dgo "github.com/bwmarrin/discordgo"
)

var token string

func init() {
	token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		fmt.Println("No token found")
	}
	fmt.Println(token)
}

func main() {

	dg, err := dgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}
	_ = dg
	err = dg.Open()
	if err != nil {
		panic(err)
	}
	defer dg.Close()
	var input string
	fmt.Scanln(&input)

}
