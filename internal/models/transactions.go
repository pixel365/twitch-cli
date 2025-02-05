// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package models

type TransactionEventSubEvent struct {
	ID                   string                     `json:"id"`
	ExtensionClientID    string                     `json:"extension_client_id"`
	BroadcasterUserID    string                     `json:"broadcaster_user_id"`
	BroadcasterUserLogin string                     `json:"broadcaster_user_login"`
	BroadcasterUserName  string                     `json:"broadcaster_user_name"`
	UserName             string                     `json:"user_name"`
	UserLogin            string                     `json:"user_login"`
	UserID               string                     `json:"user_id"`
	Product              TransactionEventSubProduct `json:"product"`
}

type TransactionEventSubProduct struct {
	Name          string `json:"name"`
	Sku           string `json:"sku"`
	Bits          int64  `json:"bits"`
	InDevelopment bool   `json:"in_development"`
}

type TransactionEventSubResponse struct {
	Subscription EventsubSubscription     `json:"subscription"`
	Event        TransactionEventSubEvent `json:"event"`
}

type TransactionProduct struct {
	Sku           string          `json:"sku"`
	DisplayName   string          `json:"displayName"`
	Domain        string          `json:"domain"`
	Expiration    string          `json:"expiration"`
	Cost          TransactionCost `json:"cost"`
	InDevelopment bool            `json:"inDevelopment"`
	Broadcast     bool            `json:"broadcast"`
}

type TransactionCost struct {
	Type   string `json:"type"`
	Amount int64  `json:"amount"`
}
