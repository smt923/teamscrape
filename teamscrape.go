/*
Package teamscrape scrapes and pulls information from a Twitch.tv team page and their APIs

The main ScrapeTwitchTeam function should be pretty fast however combining the two api functions
for multiple people can get pretty slow when you call them both on many people (10+ seconds for a small team)
*/
package teamscrape

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/yhat/scrape"

	"golang.org/x/net/html"
)

var clientid = os.Getenv("TWITCH_CLIENT_ID")

const apiver = `application/vnd.twitchtv.v3+json`

// ScrapeTwitchTeam will take a twitch team url and scrape it for a list of members
func ScrapeTwitchTeam(team string) (members []string, info string, err error) {

	if strings.ContainsAny(team, ` /.'"\`) || len(team) <= 0 {
		return members, info, errors.New("Team name was invalid")
	}

	respinfo, err := http.Get("https://www.twitch.tv/team/" + team)
	if err != nil {
		panic(err)
	}
	rootinfo, err := html.Parse(respinfo.Body)
	if err != nil {
		panic(err)
	}

	infoscrape, ok := scrape.Find(rootinfo, scrape.ById("about"))
	if ok != true {
		info = ""
	} else {
		info = scrape.Text(infoscrape)
	}

	for i := 1; i < 25; i++ {
		page := strconv.Itoa(i)

		respmems, err := http.Get("https://www.twitch.tv/team/" + team + "/live_member_list?page=" + page)
		if err != nil {
			panic(err)
		}
		rootmems, err := html.Parse(respmems.Body)
		if err != nil {
			panic(err)
		}

		mems := scrape.FindAll(rootmems, scrape.ByClass("member_name"))

		if len(mems) <= 0 {
			return members, info, nil
		}

		// We've already checked for something to be here, this should never be empty
		// (but probably don't assume this)

		for _, mem := range mems {
			members = append(members, scrape.Text(mem))
		}
	}
	return members, info, nil
}

// TwitchUserReturner accesses the twitch api and returns a twitchuser object
func TwitchUserReturner(username string) *TwitchUser {
	twitchuser := new(TwitchUser)

	twitchclient := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.twitch.tv/kraken/users/"+username, nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Accept", apiver)
	req.Header.Add("Client-ID", clientid)

	r, err := twitchclient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(twitchuser)

	return twitchuser
}

// TwitchStreamReturner accesses the twitch api and returns a twitchstream object
func TwitchStreamReturner(username string) *TwitchStreamOnline {
	twitchstream := new(TwitchStreamOnline)

	twitchclient := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.twitch.tv/kraken/streams/"+username, nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Accept", apiver)
	req.Header.Add("Client-ID", clientid)

	r, err := twitchclient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(twitchstream)

	return twitchstream
}
