// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package models

type CharityEventSubEventAmount struct {
	Currency      string `json:"currency"`
	Value         int    `json:"value"`
	DecimalPlaces int    `json:"decimal_places"`
}

type CharityEventSubEvent struct {
	Amount               *CharityEventSubEventAmount `json:"amount,omitempty"`
	StoppedAt            *string                     `json:"stopped_at,omitempty"`
	StartedAt            *string                     `json:"started_at,omitempty"`
	TargetAmount         *CharityEventSubEventAmount `json:"target_amount,omitempty"`
	CampaignID           *string                     `json:"campaign_id,omitempty"`
	UserID               *string                     `json:"user_id,omitempty"`
	UserName             *string                     `json:"user_name,omitempty"`
	UserLogin            *string                     `json:"user_login,omitempty"`
	CurrentAmount        *CharityEventSubEventAmount `json:"current_amount,omitempty"`
	BroadcasterUserLogin string                      `json:"broadcaster_user_login"`
	CharityLogo          string                      `json:"charity_logo"`
	CharityWebsite       string                      `json:"charity_website,omitempty"`
	CharityDescription   string                      `json:"charity_description,omitempty"`
	CharityName          string                      `json:"charity_name"`
	BroadcasterUserName  string                      `json:"broadcaster_user_name"`
	BroadcasterUserID    string                      `json:"broadcaster_user_id"`
	ID                   string                      `json:"id,omitempty"`
}

type CharityEventSubResponse struct {
	Event        CharityEventSubEvent `json:"event"`
	Subscription EventsubSubscription `json:"subscription"`
}
