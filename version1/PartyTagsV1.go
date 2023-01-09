package version1

import "time"

type PartyTagsV1 struct {
	Id         string         `json:"id"`
	Tags       []*TagRecordV1 `json:"tags"`
	ChangeTime time.Time      `json:"change_time"`
}

func NewPartyTagsV1(id string, tags []*TagRecordV1) *PartyTagsV1 {
	c := &PartyTagsV1{
		Id:         id,
		ChangeTime: time.Now(),
	}
	if tags == nil {
		c.Tags = make([]*TagRecordV1, 0)
	} else {
		c.Tags = tags
	}

	return c
}
