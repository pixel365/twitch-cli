// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package models

type PollEventSubResponse struct {
	Subscription EventsubSubscription `json:"subscription"`
	Event        PollEventSubEvent    `json:"event"`
}

type PollEventSubEvent struct {
	ID                   string                      `json:"id"`
	BroadcasterUserID    string                      `json:"broadcaster_user_id"`
	BroadcasterUserLogin string                      `json:"broadcaster_user_login"`
	BroadcasterUserName  string                      `json:"broadcaster_user_name"`
	Title                string                      `json:"title"`
	Status               string                      `json:"status,omitempty"`
	StartedAt            string                      `json:"started_at"`
	EndsAt               string                      `json:"ends_at,omitempty"`
	EndedAt              string                      `json:"ended_at,omitempty"`
	Choices              []PollEventSubEventChoice   `json:"choices"`
	BitsVoting           PollEventSubEventGoodVoting `json:"bits_voting"`
	ChannelPointsVoting  PollEventSubEventGoodVoting `json:"channel_points_voting"`
}

type PollEventSubEventChoice struct {
	BitsVotes          *int   `json:"bits_votes,omitempty"`
	ChannelPointsVotes *int   `json:"channel_points_votes,omitempty"`
	Votes              *int   `json:"votes,omitempty"`
	ID                 string `json:"id"`
	Title              string `json:"title"`
}

type PollEventSubEventGoodVoting struct {
	IsEnabled     bool `json:"is_enabled"`
	AmountPerVote int  `json:"amount_per_vote"`
}
