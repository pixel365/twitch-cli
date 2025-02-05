// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package models

// channel.shoutout.create

type ShoutoutCreateEventSubResponse struct {
	Event        ShoutoutCreateEventSubEvent `json:"event"`
	Subscription EventsubSubscription        `json:"subscription"`
}

type ShoutoutCreateEventSubEvent struct {
	ModeratorUserID        string `json:"moderator_user_id"`
	BroadcasterUserName    string `json:"broadcaster_user_name"`
	BroadcasterUserLogin   string `json:"broadcaster_user_login"`
	ToBroadcasterUserID    string `json:"to_broadcaster_user_id"`
	ToBroadcasterUserName  string `json:"to_broadcaster_user_name"`
	ToBroadcasterUserLogin string `json:"to_broadcaster_user_login"`
	BroadcasterUserID      string `json:"broadcaster_user_id"`
	ModeratorUserName      string `json:"moderator_user_name"`
	ModeratorUserLogin     string `json:"moderator_user_login"`
	StartedAt              string `json:"started_at"`
	CooldownEndsAt         string `json:"cooldown_ends_at"`
	TargetCooldownEndsAt   string `json:"target_cooldown_ends_at"`
	ViewerCount            int    `json:"viewer_count"`
}

// channel.shoutout.receive

type ShoutoutReceivedEventSubResponse struct {
	Event        ShoutoutReceivedEventSubEvent `json:"event"`
	Subscription EventsubSubscription          `json:"subscription"`
}

type ShoutoutReceivedEventSubEvent struct {
	BroadcasterUserID        string `json:"broadcaster_user_id"`
	BroadcasterUserName      string `json:"broadcaster_user_name"`
	BroadcasterUserLogin     string `json:"broadcaster_user_login"`
	FromBroadcasterUserID    string `json:"from_broadcaster_user_id"`
	FromBroadcasterUserName  string `json:"from_broadcaster_user_name"`
	FromBroadcasterUserLogin string `json:"from_broadcaster_user_login"`
	StartedAt                string `json:"started_at"`
	ViewerCount              int    `json:"viewer_count"`
}
