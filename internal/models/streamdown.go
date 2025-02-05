// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package models

type StreamDownEventSubResponse struct {
	Event        StreamUpEventSubEvent `json:"event"`
	Subscription EventsubSubscription  `json:"subscription"`
}

type StreamDownEventSubEvent struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
}
