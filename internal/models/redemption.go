// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package models

type RedemptionEventSubEvent struct {
	ID                   string           `json:"id"`
	BroadcasterUserID    string           `json:"broadcaster_user_id"`
	BroadcasterUserLogin string           `json:"broadcaster_user_login"`
	BroadcasterUserName  string           `json:"broadcaster_user_name"`
	UserID               string           `json:"user_id"`
	UserLogin            string           `json:"user_login"`
	UserName             string           `json:"user_name"`
	UserInput            string           `json:"user_input"`
	Status               string           `json:"status"`
	RedeemedAt           string           `json:"redeemed_at"`
	Reward               RedemptionReward `json:"reward"`
}

type RedemptionReward struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Prompt string `json:"prompt"`
	Cost   int64  `json:"cost"`
}

type RedemptionEventSubResponse struct {
	Event        RedemptionEventSubEvent `json:"event"`
	Subscription EventsubSubscription    `json:"subscription"`
}
