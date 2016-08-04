[![Build Status](https://travis-ci.org/smt923/teamscrape.svg?branch=master)](https://travis-ci.org/smt923/teamscrape)

# teamscrape
Scrape information from twitch.tv teams

# Installation
go get -u github.com/smt923/teamscrape

import "github.com/smt923/teamscrape"

# Instructions

twitch client ID is stored in the environment variable "TWITCH_CLIENT_ID", I don't think it's needed for any of the api calls (defintely not needed for the main scraping function) but if the api calls fail you can get a key from [here]( https://www.twitch.tv/settings/connections) and then set the environment variable

### ScrapeTwitchTeam
```go
members, info, err := teamscrape.ScrapeTwitchTeam("teamname")
```

will return a range-able list of members, the information in the team bio (if any) and an error, if any

### Twitch API returners
```go
user   := teamscrape.TwitchUserReturner("username")
stream := teamscrape.TwitchStreamReturner("username")
```
these will return a simple way to access some basic helpers for common twitch apis (user and stream objects)
this is slow, and mainly added as a helper for me to access some common information that might go towards a twitch team, if you need to access other parts of the api you can do a manual request with the information from ScrapeTwitchTeam

