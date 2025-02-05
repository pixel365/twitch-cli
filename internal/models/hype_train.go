// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package models

type ContributionData struct {
	TypeOfContribution           string `json:"type"`
	UserWhoMadeContribution      string `json:"user_id,omitempty"`
	UserNameWhoMadeContribution  string `json:"user_name,omitempty"`
	UserLoginWhoMadeContribution string `json:"user_login,omitempty"`
	TotalContribution            int64  `json:"total"`
}

type HypeTrainEventSubResponse struct {
	Subscription EventsubSubscription   `json:"subscription"`
	Event        HypeTrainEventSubEvent `json:"event"`
}

type HypeTrainEventSubEvent struct {
	Progress                *int64             `json:"progress,omitempty"`
	LastContribution        ContributionData   `json:"last_contribution,omitempty"`
	BroadcasterUserLogin    string             `json:"broadcaster_user_login"`
	BroadcasterUserName     string             `json:"broadcaster_user_name"`
	BroadcasterUserID       string             `json:"broadcaster_user_id"`
	ID                      string             `json:"id"`
	StartedAtTimestamp      string             `json:"started_at,omitempty"`
	ExpiresAtTimestamp      string             `json:"expires_at,omitempty"`
	EndedAtTimestamp        string             `json:"ended_at,omitempty"`
	CooldownEndsAtTimestamp string             `json:"cooldown_ends_at,omitempty"`
	TopContributions        []ContributionData `json:"top_contributions"`
	Level                   int64              `json:"level,omitempty"`
	Total                   int64              `json:"total"`
	Goal                    int64              `json:"goal,omitempty"`
}
