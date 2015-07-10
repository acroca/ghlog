package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var client *github.Client
var currentUser *github.User

type IssueEvent struct {
	Action string
	Issue  struct {
		HtmlUrl string `json:"html_url"`
		Title   string
	}
}

type PullRequestEvent struct {
	Action      string
	PullRequest struct {
		HtmlUrl string `json:"html_url"`
		Title   string
	} `json:"pull_request"`
}

type IssueCommentEvent struct {
	Action  string
	Comment struct {
		HtmlUrl string `json:"html_url"`
	}
	Issue struct {
		Title string
	}
}

type PushEvent struct {
	Commits []struct {
		Message string
	}
}

func main() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GH_TOKEN")},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client = github.NewClient(tc)
	currentUser = fetchCurrentUser()

	events := fetchEvents()

	for _, event := range events {
		fmt.Println("-----------------------------------------")
		fmt.Println(
			"Type:", *event.Type,
			"; Org:", *event.Org.Login,
			"; ", time.Now().Sub(*event.CreatedAt), "ago.")

		switch *event.Type {
		case "IssuesEvent":
			e := &IssueEvent{}
			json.Unmarshal(*event.RawPayload, &e)
			fmt.Println(e.Action, " - ", e.Issue.Title, "\nURL:", e.Issue.HtmlUrl)
		case "PullRequestEvent":
			e := &PullRequestEvent{}
			json.Unmarshal(*event.RawPayload, &e)
			fmt.Println(e.Action, " - ", e.PullRequest.Title, "\nURL:", e.PullRequest.HtmlUrl)
		case "IssueCommentEvent":
			e := &IssueCommentEvent{}
			json.Unmarshal(*event.RawPayload, &e)
			fmt.Println(e.Action, " - ", e.Issue.Title, "\nURL:", e.Comment.HtmlUrl)
		case "PushEvent":
			e := &PushEvent{}
			json.Unmarshal(*event.RawPayload, &e)
			if event.Repo != nil {
				fmt.Println("Repo:", *event.Repo.Name)
			}
			for _, commit := range e.Commits {
				fmt.Println(" *", commit.Message)
			}
		case "CreateEvent":
			// Ignore
		case "DeleteEvent":
			// Ignore
		default:
			raw, _ := event.RawPayload.MarshalJSON()
			fmt.Println("RAW JSON:", string(raw))
		}

	}
}

func fetchCurrentUser() *github.User {
	user, _, err := client.Users.Get("")
	if err != nil {
		panic(err)
	}
	return user
}

func fetchEvents() []github.Event {
	events, _, err := client.Activity.ListEventsPerformedByUser(
		*currentUser.Login, false, &github.ListOptions{})
	if err != nil {
		panic(err)
	}
	return events
}
