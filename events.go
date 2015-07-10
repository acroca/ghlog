package main

import (
	"fmt"
	"time"
)

// Event is used to send event lists
type Event interface {
	GetEventBody() string
}

// IssueEvent are events of type IssueEvent
type IssueEvent struct {
	Type      string    `json:"-"`
	Org       string    `json:"-"`
	CreatedAt time.Time `json:"-"`
	Action    string
	Issue     struct {
		HtmlUrl string `json:"html_url"`
		Title   string
	}
}

// GetEventBody returns the body string for the event
func (e *IssueEvent) GetEventBody() string {
	base := fmt.Sprintf("Type: %v; Org: %v; %v ago",
		e.Type, e.Org, time.Now().Sub(e.CreatedAt))
	return fmt.Sprintf("%v\n%v - %v\nURL: %v",
		base, e.Action, e.Issue.Title, e.Issue.HtmlUrl)
}

// PullRequestEvent are events of type PullRequestEvent
type PullRequestEvent struct {
	Type        string    `json:"-"`
	Org         string    `json:"-"`
	CreatedAt   time.Time `json:"-"`
	Action      string
	PullRequest struct {
		HtmlUrl string `json:"html_url"`
		Title   string
	} `json:"pull_request"`
}

// GetEventBody returns the body string for the event
func (e *PullRequestEvent) GetEventBody() string {
	base := fmt.Sprintf("Type: %v; Org: %v; %v ago",
		e.Type, e.Org, time.Now().Sub(e.CreatedAt))
	return fmt.Sprintf("%v\n%v - %v\nURL: %v",
		base, e.Action, e.PullRequest.Title, e.PullRequest.HtmlUrl)
}

// IssueCommentEvent are events of type IssueCommentEvent
type IssueCommentEvent struct {
	Type      string    `json:"-"`
	Org       string    `json:"-"`
	CreatedAt time.Time `json:"-"`
	Action    string
	Comment   struct {
		HtmlUrl string `json:"html_url"`
	}
	Issue struct {
		Title string
	}
}

// GetEventBody returns the body string for the event
func (e *IssueCommentEvent) GetEventBody() string {
	base := fmt.Sprintf("Type: %v; Org: %v; %v ago",
		e.Type, e.Org, time.Now().Sub(e.CreatedAt))
	return fmt.Sprintf("%v\n%v - %v\nURL: %v",
		base, e.Action, e.Issue.Title, e.Comment.HtmlUrl)
}

// PushEvent are events of type PushEvent
type PushEvent struct {
	Type      string    `json:"-"`
	Org       string    `json:"-"`
	CreatedAt time.Time `json:"-"`
	Repo      string    `json:"-"`
	Commits   []struct {
		Message string
	}
}

// GetEventBody returns the body string for the event
func (e *PushEvent) GetEventBody() string {
	base := fmt.Sprintf("Type: %v; Org: %v; %v ago",
		e.Type, e.Org, time.Now().Sub(e.CreatedAt))
	commits := ""
	for _, commit := range e.Commits {
		commits += " * " + commit.Message + "\n"
	}

	return fmt.Sprintf("%v\nRepo: %v\n%v",
		base, e.Repo, commits)
}
