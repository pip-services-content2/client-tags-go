package version1

import "context"

type ITagsClientV1 interface {
	GetTags(ctx context.Context, correlationId string, partyId string) (*PartyTagsV1, error)

	SetTags(ctx context.Context, correlationId string, partyTags *PartyTagsV1) (*PartyTagsV1, error)

	RecordTags(ctx context.Context, correlationId string, partyId string, tags []string) (*PartyTagsV1, error)
}
