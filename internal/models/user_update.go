// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package models

type UserUpdateEventSubResponse struct {
	Event        StreamUpEventSubEvent `json:"event"`
	Subscription EventsubSubscription  `json:"subscription"`
}

type UserUpdateEventSubEvent struct {
	UserID        string `json:"user_id"`
	UserLogin     string `json:"user_login"`
	UserName      string `json:"user_name"`
	Email         string `json:"email"`
	Description   string `json:"description"`
	EmailVerified bool   `json:"email_verified"`
}
