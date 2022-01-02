package emotes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type BttvEmote struct {
	ID        string
	Code      string
	ImageType string
	User      interface{}
}

type BttvJSON struct {
	ID            string
	Bots          []string
	Avatar        string
	ChannelEmotes []BttvEmote
	SharedEmotes  []BttvEmote
}

type FfzEmote struct {
	ID        int
	User      interface{}
	Code      string
	Images    interface{}
	ImageType string
}

type SevenTvEmote struct {
	ID                string
	Name              string
	Owner             interface{}
	Visability        int
	Visability_simple []interface{}
	Mime              string
	Status            int
	Tags              []interface{}
	Width             []interface{}
	Height            []interface{}
	Urls              []interface{}
}

type SubscriberEmote struct {
	Data []struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Images struct {
			URL1X string `json:"url_1x"`
			URL2X string `json:"url_2x"`
			URL4X string `json:"url_4x"`
		} `json:"images"`
		EmoteType  string   `json:"emote_type"`
		EmoteSetID string   `json:"emote_set_id"`
		OwnerID    string   `json:"owner_id"`
		Format     []string `json:"format"`
		Scale      []string `json:"scale"`
		ThemeMode  []string `json:"theme_mode"`
	} `json:"data"`
	Template string `json:"template"`
}

func FetchEmotes(id int) map[string]bool {

	set := make(map[string]bool, 0)

	fetchBTTVGlobalEmotes(set)
	fetchBTTVChannelEmotes(set, id)
	fetchFFZEmotes(set, id)
	fetchSevenTvEmotes(set, id)

	// for emote := range set {
	// 	fmt.Println(emote)
	// }

	return set
}

func fetchFFZEmotes(set map[string]bool, id int) {
	url := fmt.Sprintf("https://api.betterttv.net/3/cached/frankerfacez/users/twitch/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("there was an error reading ffz channel emotes. error: ", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("there was an error reading ffz channel emotes. error: ", err)
		return
	}

	var ffzEmotes []FfzEmote
	err = json.Unmarshal(body, &ffzEmotes)
	if err != nil {
		log.Println("there was an error reading ffz channel emotes. error: ", err)
		return
	}

	for _, emote := range ffzEmotes {
		set[emote.Code] = true
	}
}

func fetchSevenTvEmotes(set map[string]bool, id int) {
	url := fmt.Sprintf("https://api.7tv.app/v2/users/%d/emotes", id)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("There was an error reading 7tv emotes. Error: ", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("there was an error reading 7tv channel emotes. error: ", err)
		return
	}

	var sevenTvEmotes []SevenTvEmote
	err = json.Unmarshal(body, &sevenTvEmotes)
	if err != nil {
		log.Println("there was an error reading 7tv channel emotes. error: ", err)
		return
	}

	for _, emote := range sevenTvEmotes {
		set[emote.Name] = true
	}
}

func fetchBTTVChannelEmotes(set map[string]bool, id int) {
	url := fmt.Sprintf("https://api.betterttv.net/3/cached/users/twitch/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("there was an error reading bttv channel emotes. error: ", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("there was an error reading bttv channel emotes. error: ", err)
		return
	}

	var bttvJSON BttvJSON
	err = json.Unmarshal(body, &bttvJSON)
	if err != nil {
		log.Println("there was an error reading bttv channel emotes. error: ", err)
		return
	}

	for _, emote := range bttvJSON.ChannelEmotes {
		set[emote.Code] = true
	}

	for _, emote := range bttvJSON.SharedEmotes {
		set[emote.Code] = true
	}
}

func fetchBTTVGlobalEmotes(set map[string]bool) {
	resp, err := http.Get("https://api.betterttv.net/3/cached/emotes/global")
	if err != nil {
		log.Println("there was an error reading bttv global emotes. error: ", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("there was an error reading bttv global emotes. error: ", err)
		return
	}

	var bttvEmotes []BttvEmote
	err = json.Unmarshal(body, &bttvEmotes)
	if err != nil {
		log.Println("there was an error reading bttv global emotes. error: ", err)
		return
	}

	for _, emote := range bttvEmotes {
		set[emote.Code] = true
	}
}

func fetchSubscriberEmotes(set map[string]bool, id int) {
	url := fmt.Sprintf("https://api.7tv.app/v2/users/%d/emotes", id)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("There was an error reading 7tv emotes. Error: ", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("there was an error reading 7tv channel emotes. error: ", err)
		return
	}

	var sevenTvEmotes []SevenTvEmote
	err = json.Unmarshal(body, &sevenTvEmotes)
	if err != nil {
		log.Println("there was an error reading 7tv channel emotes. error: ", err)
		return
	}

	for _, emote := range sevenTvEmotes {
		set[emote.Name] = true
	}
}

