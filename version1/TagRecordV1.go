package version1

import "time"

type TagRecordV1 struct {
	Tag      string    `json:"tag"`
	Count    int       `json:"count"`
	LastTime time.Time `json:"last_time"`
}

func NewTagRecordV1(tag string, count int) *TagRecordV1 {
	return &TagRecordV1{
		Tag:      tag,
		Count:    count,
		LastTime: time.Now(),
	}
}
