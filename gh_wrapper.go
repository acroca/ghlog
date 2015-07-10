package main

import (
	"encoding/json"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// GhWrapper contains the logic to interact with Github API.
type GhWrapper struct {
	client *github.Client
	user   *User
}

// User models a github user.
type User struct {
	Login string
}

// NewGhWrapper build a new wrapper for the given token
func NewGhWrapper(token string) *GhWrapper {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)

	wrapper := &GhWrapper{
		client: client,
	}
	return wrapper
}

// GetUser fetches and returns the current user
func (w *GhWrapper) GetUser() *User {
	if w.user != nil {
		return w.user
	}
	user, _, err := w.client.Users.Get("")
	if err != nil {
		panic(err)
	}
	w.user = &User{
		Login: *user.Login,
	}
	return w.user
}

// GetEvents fetches and returns the last events
func (w *GhWrapper) GetEvents() []Event {
	events, _, err := w.client.Activity.ListEventsPerformedByUser(
		w.GetUser().Login, false, &github.ListOptions{})
	if err != nil {
		panic(err)
	}
	res := make([]Event, len(events))

	for idx, event := range events {
		switch *event.Type {
		case "IssuesEvent":
			res[idx] = &IssueEvent{
				Type:      *event.Type,
				Org:       *event.Org.Login,
				CreatedAt: *event.CreatedAt,
			}
			json.Unmarshal(*event.RawPayload, &res[idx])
		case "PullRequestEvent":
			res[idx] = &PullRequestEvent{
				Type:      *event.Type,
				Org:       *event.Org.Login,
				CreatedAt: *event.CreatedAt,
			}
			json.Unmarshal(*event.RawPayload, &res[idx])
		case "IssueCommentEvent":
			res[idx] = &IssueCommentEvent{
				Type:      *event.Type,
				Org:       *event.Org.Login,
				CreatedAt: *event.CreatedAt,
			}
			json.Unmarshal(*event.RawPayload, &res[idx])
		case "PushEvent":
			res[idx] = &PushEvent{
				Type:      *event.Type,
				Org:       *event.Org.Login,
				CreatedAt: *event.CreatedAt,
			}
			json.Unmarshal(*event.RawPayload, &res[idx])
		case "CreateEvent":
			// Ignore
		case "DeleteEvent":
			// Ignore
		default:
			// raw, _ := event.RawPayload.MarshalJSON()
			// fmt.Println("RAW JSON:", string(raw))
		}
	}

	return res
}
