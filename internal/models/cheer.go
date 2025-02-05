// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package models

type CheerEventSubEvent struct {
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Message              string `json:"message"`
	Bits                 int64  `json:"bits"`
	IsAnonymous          bool   `json:"is_anonymous"`
}

type CheerEventSubResponse struct {
	Subscription EventsubSubscription `json:"subscription"`
	Event        CheerEventSubEvent   `json:"event"`
}
