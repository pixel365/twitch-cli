// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package database

import (
	"fmt"
	"log"
)

type Video struct {
	CategoryID       *string             `db:"category_id" dbs:"v.category_id" json:"-"`
	StreamID         *string             `db:"stream_id" json:"stream_id"`
	PublishedAt      string              `db:"published_at" json:"published_at"`
	Duration         string              `db:"duration" json:"duration"`
	BroadcasterName  string              `db:"broadcaster_name" json:"user_name" dbi:"false"`
	Title            string              `db:"title" json:"title"`
	VideoDescription string              `db:"video_description" json:"description"`
	CreatedAt        string              `db:"created_at" json:"created_at"`
	ID               string              `db:"id" json:"id" dbs:"v.id"`
	Viewable         string              `db:"viewable" json:"viewable"`
	PeriodDate       string              `db:"period_date" dbi:"false" json:"-"`
	BroadcasterLogin string              `db:"broadcaster_login" json:"user_login" dbi:"false"`
	VideoLanguage    string              `db:"video_language" json:"language"`
	ThumbnailURL     string              `json:"thumbnail_url"`
	BroadcasterID    string              `db:"broadcaster_id" json:"user_id"`
	Type             string              `db:"type" json:"type"`
	URL              string              `json:"url"`
	MutedSegments    []VideoMutedSegment `json:"muted_segments"`
	ViewCount        int                 `db:"view_count" json:"view_count"`
}

type VideoMutedSegment struct {
	VideoID     string `db:"video_id" json:"-"`
	VideoOffset int    `db:"video_offset" json:"offset"`
	Duration    int    `db:"duration" json:"duration"`
}

type Clip struct {
	Title           string  `db:"title" json:"title"`
	EndedAt         string  `db:"ended_at" dbi:"false" json:"-"`
	ID              string  `db:"id" json:"id" dbs:"c.id"`
	CreatorID       string  `db:"creator_id" json:"creator_id"`
	CreatorName     string  `db:"creator_name" json:"creator_name" dbi:"false"`
	VideoID         string  `db:"video_id" json:"video_id"`
	GameID          string  `db:"game_id" json:"game_id"`
	Language        string  `db:"language" dbi:"false" json:"language"`
	BroadcasterName string  `db:"broadcaster_name" json:"broadcaster_name" dbi:"false"`
	BroadcasterID   string  `db:"broadcaster_id" json:"broadcaster_id"`
	StartedAt       string  `db:"started_at" dbi:"false" json:"-"`
	CreatedAt       string  `db:"created_at" json:"created_at"`
	EmbedURL        string  `json:"embed_url"`
	ThumbnailURL    string  `json:"thumbnail_url"`
	URL             string  `json:"url"`
	VodOffset       int     `db:"vod_offset" json:"vod_offset"`
	Duration        float64 `db:"duration" json:"duration"`
	ViewCount       int     `db:"view_count" json:"view_count"`
	IsFeatured      bool    `json:"is_featured"`
}

var sortMap = map[string]string{
	"time":     " order by v.created_at desc",
	"trending": "",
	"views":    " order by v.view_count desc",
}

func (q *Query) GetVideos(v Video, period string, sort string) (*DBResponse, error) {
	var r []Video
	sql := generateSQL("select v.*, u1.user_login as broadcaster_login, u1.display_name as broadcaster_name from videos v join users u1 on v.broadcaster_id = u1.id", v, SEP_AND)

	if period != "" {
		sql += " and datetime(v.created_at) > datetime(:period_date) "
		v.PeriodDate = period
	}

	rows, err := q.DB.NamedQuery(sql+sortMap[sort]+q.SQL, v)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	for rows.Next() {
		var v Video
		err := rows.StructScan(&v)
		if err != nil {
			log.Print(err)
			return nil, err
		}

		v.ThumbnailURL = "https://static-cdn.jtvnw.net/cf_vods/d2nvs31859zcd8/twitchdev/335921245/ce0f3a7f-57a3-4152-bc06-0c6610189fb3/thumb/index-0000000000-%{width}x%{height}.jpg"
		v.URL = fmt.Sprintf("https://www.twitch.tv/videos/%v", v.ID)
		r = append(r, v)
	}

	for i, v := range r {
		vms := []VideoMutedSegment{}
		err = q.DB.Select(&vms, "select * from video_muted_segments where video_id=$1", v.ID)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		r[i].MutedSegments = vms

		if r[i].Type == "" {
			r[i].Type = "archive"
		}
	}

	dbr := DBResponse{
		Data:  r,
		Limit: q.Limit,
		Total: len(r),
	}

	if len(r) != q.Limit {
		q.PaginationCursor = ""
	}

	dbr.Cursor = q.PaginationCursor

	return &dbr, err
}

func (q *Query) InsertVideo(v Video) error {
	stmt := generateInsertSQL("videos", "", v, false)
	_, err := q.DB.NamedExec(stmt, v)
	return err
}

func (q *Query) DeleteVideo(id string) error {
	tx := q.DB.MustBegin()
	tx.MustExec("delete from stream_markers where video_id=$1", id)
	tx.MustExec("delete from video_muted_segments where video_id=$1", id)
	tx.MustExec("delete from videos where id = $1", id)
	return tx.Commit()
}

func (q *Query) InsertMutedSegmentsForVideo(vms VideoMutedSegment) error {
	stmt := generateInsertSQL("video_muted_segments", "", vms, false)
	_, err := q.DB.NamedExec(stmt, vms)
	return err
}

func (q *Query) InsertClip(c Clip) error {
	stmt := generateInsertSQL("clips", "", c, false)
	_, err := q.DB.NamedExec(stmt, c)
	return err
}

func (q *Query) GetClips(c Clip, startDate string, endDate string) (*DBResponse, error) {
	var r []Clip
	sql := generateSQL("select c.id, c.broadcaster_id, c.creator_id, c.video_id, c.game_id, c.title, c.view_count, c.duration, c.vod_offset, datetime(c.created_at) as created_at,  u1.display_name as broadcaster_name, u1.stream_language as language, u2.display_name as creator_name from clips c join users u1 on c.broadcaster_id = u1.id join users u2 on c.creator_id = u2.id ", c, SEP_AND)
	if startDate != "" {
		c.StartedAt = startDate
		c.EndedAt = endDate
		sql += " and datetime(c.created_at) > datetime(:started_at) and datetime(c.created_at) < datetime(:ended_at) "
	}
	sql += q.SQL
	rows, err := q.DB.NamedQuery(sql, c)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	for rows.Next() {
		var c Clip
		err := rows.StructScan(&c)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		c.EmbedURL = fmt.Sprintf("https://clips.twitch.tv/embed?clip=%v", c.ID)
		c.ThumbnailURL = "https://clips-media-assets2.twitch.tv/157589949-preview-480x272.jpg"
		c.URL = fmt.Sprintf("https://clips.twitch.tv/%v", c.ID)
		r = append(r, c)
	}

	dbr := DBResponse{
		Data:  r,
		Limit: q.Limit,
		Total: len(r),
	}

	if len(r) != q.Limit {
		q.PaginationCursor = ""
	}

	dbr.Cursor = q.PaginationCursor

	return &dbr, err
}
