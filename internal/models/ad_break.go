// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package models

type AdBreakBeginEventSubEvent struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	RequesterUserID      string `json:"requester_user_id"`
	RequesterUserLogin   string `json:"requester_user_login"`
	RequesterUserName    string `json:"requester_user_name"`
	StartedAt            string `json:"started_at"`
	Duration             int    `json:"duration_seconds"`
	IsAutomatic          bool   `json:"is_automatic"`
}

type AdBreakBeginEventSubResponse struct {
	Subscription EventsubSubscription      `json:"subscription"`
	Event        AdBreakBeginEventSubEvent `json:"event"`
}
