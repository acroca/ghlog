package main

import (
	"fmt"
	"time"
)

// Event is used to send event lists
type Event interface {
	GetEventBody() string
}

// BaseEvent is the generic data structure for events
type BaseEvent struct {
	Type      string
	Org       string
	CreatedAt time.Time
}

// GetEventBody returns the body string for the event
func (e *BaseEvent) GetEventBody() string {
	return fmt.Sprintf("Type: %v; Org: %v; %.1f hours ago",
		e.Type, e.Org, time.Now().Sub(e.CreatedAt).Hours())
}

// IssueEvent are events of type IssueEvent
type IssueEvent struct {
	Base   *BaseEvent
	Action string
	Issue  struct {
		HtmlUrl string `json:"html_url"`
		Title   string
	}
}

// GetEventBody returns the body string for the event
func (e *IssueEvent) GetEventBody() string {
	return fmt.Sprintf("%v\n%v - %v\nURL: %v",
		e.Base.GetEventBody(), e.Action, e.Issue.Title, e.Issue.HtmlUrl)
}

// PullRequestEvent are events of type PullRequestEvent
type PullRequestEvent struct {
	Base        *BaseEvent
	Action      string
	PullRequest struct {
		HtmlUrl string `json:"html_url"`
		Title   string
	} `json:"pull_request"`
}

// GetEventBody returns the body string for the event
func (e *PullRequestEvent) GetEventBody() string {
	return fmt.Sprintf("%v\n%v - %v\nURL: %v",
		e.Base.GetEventBody(), e.Action, e.PullRequest.Title, e.PullRequest.HtmlUrl)
}

// IssueCommentEvent are events of type IssueCommentEvent
type IssueCommentEvent struct {
	Base    *BaseEvent
	Action  string
	Comment struct {
		HtmlUrl string `json:"html_url"`
	}
	Issue struct {
		Title string
	}
}

// GetEventBody returns the body string for the event
func (e *IssueCommentEvent) GetEventBody() string {
	return fmt.Sprintf("%v\n%v - %v\nURL: %v",
		e.Base.GetEventBody(), e.Action, e.Issue.Title, e.Comment.HtmlUrl)
}

// PushEvent are events of type PushEvent
type PushEvent struct {
	Base    *BaseEvent
	Repo    string `json:"-"`
	Commits []struct {
		Message string
	}
}

// GetEventBody returns the body string for the event
func (e *PushEvent) GetEventBody() string {
	commits := ""
	for _, commit := range e.Commits {
		commits += " * " + commit.Message + "\n"
	}

	return fmt.Sprintf("%v\nRepo: %v\n%v",
		e.Base.GetEventBody(), e.Repo, commits)
}
