// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package models

type PredictionEventSubResponse struct {
	Event        PredictionEventSubEvent `json:"event"`
	Subscription EventsubSubscription    `json:"subscription"`
}

type PredictionEventSubEvent struct {
	ID                   string                            `json:"id"`
	BroadcasterUserID    string                            `json:"broadcaster_user_id"`
	BroadcasterUserLogin string                            `json:"broadcaster_user_login"`
	BroadcasterUserName  string                            `json:"broadcaster_user_name"`
	Title                string                            `json:"title"`
	WinningOutcomeID     string                            `json:"winning_outcome_id,omitempty"`
	StartedAt            string                            `json:"started_at"`
	LocksAt              string                            `json:"locks_at,omitempty"`
	LockedAt             string                            `json:"locked_at,omitempty"`
	EndedAt              string                            `json:"ended_at,omitempty"`
	Status               string                            `json:"status,omitempty"`
	Outcomes             []PredictionEventSubEventOutcomes `json:"outcomes"`
}

type PredictionEventSubEventOutcomes struct {
	Users         *int                                    `json:"users,omitempty"`
	ChannelPoints *int                                    `json:"channel_points,omitempty"`
	TopPredictors *[]PredictionEventSubEventTopPredictors `json:"top_predictors,omitempty"`
	ID            string                                  `json:"id"`
	Title         string                                  `json:"title"`
	Color         string                                  `json:"color"`
}

type PredictionEventSubEventTopPredictors struct {
	ChannelPointsWon  *int   `json:"channel_points_won"`
	UserID            string `json:"user_id"`
	UserLogin         string `json:"user_login"`
	UserName          string `json:"user_name"`
	ChannelPointsUsed int    `json:"channel_points_used"`
}
