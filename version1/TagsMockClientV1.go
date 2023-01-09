package version1

import (
	"context"
	"sort"
	"time"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type TagsMockClientV1 struct {
	maxTagsCount int
	tags         []*PartyTagsV1
}

func NewTagsMockClientV1() *TagsMockClientV1 {
	return &TagsMockClientV1{
		tags:         make([]*PartyTagsV1, 0),
		maxTagsCount: 100,
	}
}

func (c *TagsMockClientV1) GetTags(ctx context.Context, correlationId string, partyId string) (result *PartyTagsV1, err error) {
	for _, v := range c.tags {
		if v.Id == partyId {
			buf := *v
			result = &buf
			break
		}
	}

	if result == nil {
		result = NewPartyTagsV1(partyId, make([]*TagRecordV1, 0))
	}

	return result, nil
}

func (c *TagsMockClientV1) SetTags(ctx context.Context, correlationId string, partyTags *PartyTagsV1) (*PartyTagsV1, error) {
	updateIndex := -1
	for i, tag := range c.tags {
		if tag.Id == partyTags.Id {
			updateIndex = i
			break
		}
	}

	partyTags.ChangeTime = time.Now()

	buf := *partyTags
	if updateIndex != -1 {
		c.tags[updateIndex] = partyTags
	} else {
		c.tags = append(c.tags, &buf)
	}

	return partyTags, nil
}

func (c *TagsMockClientV1) updateTags(partyTags *PartyTagsV1, tags []string) *PartyTagsV1 {
	if partyTags.Tags == nil {
		partyTags.Tags = make([]*TagRecordV1, 0)
	}

	// Add or update tags, increment their count and update last used time
	for _, tag := range tags {
		var tagRecord *TagRecordV1
		for _, r := range partyTags.Tags {
			if data.TagsProcessor.EqualTags(r.Tag, tag) {
				tagRecord = r
				break
			}
		}

		if tagRecord != nil {
			tagRecord.Tag = tag
			tagRecord.Count += 1
			tagRecord.LastTime = time.Now()
		} else {
			partyTags.Tags = append(partyTags.Tags, NewTagRecordV1(tag, 1))
		}
	}

	return partyTags
}

func (c *TagsMockClientV1) trimTags(partyTags *PartyTagsV1, maxLength int) *PartyTagsV1 {
	if partyTags.Tags == nil {
		partyTags.Tags = make([]*TagRecordV1, 0)
	}

	if maxLength == 0 {
		maxLength = 1000
	}
	if maxLength > len(partyTags.Tags) {
		maxLength = len(partyTags.Tags)
	}

	// Limit number of tags. Remove older less used tags
	if len(partyTags.Tags) > maxLength {
		sort.Slice(partyTags.Tags, func(i, j int) bool {
			return partyTags.Tags[i].LastTime.Unix() > partyTags.Tags[j].LastTime.Unix()
		})
		partyTags.Tags = partyTags.Tags[0:maxLength]
	}

	return partyTags
}

func (c *TagsMockClientV1) RecordTags(ctx context.Context, correlationId string, partyId string, tags []string) (*PartyTagsV1, error) {
	if tags == nil {
		tags = make([]string, 0)
	}

	tags = data.TagsProcessor.NormalizeTags(tags)

	if len(tags) == 0 {
		return nil, nil
	}

	partyTags, err := c.GetTags(ctx, correlationId, partyId)
	if err != nil {
		return nil, err
	}

	if partyTags == nil {
		partyTags = NewPartyTagsV1(partyId, make([]*TagRecordV1, 0))
	}

	partyTags = c.updateTags(partyTags, tags)
	partyTags = c.trimTags(partyTags, c.maxTagsCount)

	return c.SetTags(ctx, correlationId, partyTags)

}
