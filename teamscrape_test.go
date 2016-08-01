package teamscrape

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestScrapeTwitchTeam(t *testing.T) {
	t.Log("[INFO] Testing case: Blank string")
	if _, _, err := ScrapeTwitchTeam(""); err == nil {
		t.Error("ScrapeTwitchTeam failed to stop a blank string")
	}

	t.Log("[INFO] Testing case: Invalid string")
	if _, _, err := ScrapeTwitchTeam(`/te'a"m`); err == nil {
		t.Error("ScrapeTwitchTeam failed to stop an invalid team name")
	}

	t.Log("[INFO] Testing case: Valid string")
	if _, _, err := ScrapeTwitchTeam("wobblers"); err != nil {
		t.Error("ScrapeTwitchTeam errored on a valid team:\n", err)
	}
}

func TestTwitchAPI(t *testing.T) {
	t.Log("[INFO] Testing if we can prepare and send the API request")

	twitchroot := new(TwitchRoot)

	twitchtest := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.twitch.tv/kraken/", nil)
	if err != nil {
		t.Error("Preparing req failed: ", err)
	}
	req.Header.Add("Accept", apiver)
	req.Header.Add("Client-ID", clientid)

	r, err := twitchtest.Do(req)
	if err != nil {
		t.Error("Doing req failed: ", err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(twitchroot)

	t.Log("[INFO] Testing if we have a valid Client ID (env var: TWITCH_CLIENT_ID)")

	if twitchroot.Identified != true {
		t.Error("We do not have a valid Client ID")
	}

}
