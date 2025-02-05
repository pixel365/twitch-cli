// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package models

type RewardEventSubEvent struct {
	DefaultImage                      RewardImage          `json:"default_image"`
	Image                             RewardImage          `json:"image"`
	BroadcasterUserName               string               `json:"broadcaster_user_name"`
	BroadcasterUserLogin              string               `json:"broadcaster_user_login"`
	ID                                string               `json:"id"`
	BroadcasterUserID                 string               `json:"broadcaster_user_id"`
	Title                             string               `json:"title"`
	Prompt                            string               `json:"prompt"`
	BackgroundColor                   string               `json:"background_color"`
	CooldownExpiresAt                 string               `json:"cooldown_expires_at"`
	MaxPerUserPerStream               RewardMax            `json:"max_per_user_per_stream"`
	GlobalCooldown                    RewardGlobalCooldown `json:"global_cooldown"`
	MaxPerStream                      RewardMax            `json:"max_per_stream"`
	Cost                              int64                `json:"cost"`
	RedemptionsRedeemedCurrentStream  int64                `json:"redemptions_redeemed_current_stream"`
	IsPaused                          bool                 `json:"is_paused"`
	ShouldRedemptionsSkipRequestQueue bool                 `json:"should_redemptions_skip_request_queue"`
	IsUserInputRequired               bool                 `json:"is_user_input_required"`
	IsInStock                         bool                 `json:"is_in_stock"`
	IsEnabled                         bool                 `json:"is_enabled"`
}

type RewardMax struct {
	IsEnabled bool  `json:"is_enabled"`
	Value     int64 `json:"value"`
}

type RewardGlobalCooldown struct {
	IsEnabled bool  `json:"is_enabled"`
	Seconds   int64 `json:"seconds"`
}

type RewardImage struct {
	URL1x string `json:"url_1x"`
	URL2x string `json:"url_2x"`
	URL4x string `json:"url_4x"`
}

type RewardEventSubResponse struct {
	Subscription EventsubSubscription `json:"subscription"`
	Event        RewardEventSubEvent  `json:"event"`
}
