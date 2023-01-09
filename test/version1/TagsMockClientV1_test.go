package test_version1

import (
	"testing"

	"github.com/pip-services-content2/client-tags-go/version1"
)

type tagsMockClientV1Test struct {
	client  *version1.TagsMockClientV1
	fixture *TagsClientFixtureV1
}

func newTagsMockClientV1Test() *tagsMockClientV1Test {
	return &tagsMockClientV1Test{}
}

func (c *tagsMockClientV1Test) setup(t *testing.T) *TagsClientFixtureV1 {
	c.client = version1.NewTagsMockClientV1()

	c.fixture = NewTagsClientFixtureV1(c.client)

	return c.fixture
}

func (c *tagsMockClientV1Test) teardown(t *testing.T) {
	c.client = nil
}

func TestMockOperations(t *testing.T) {
	c := newTagsMockClientV1Test()

	fixture := c.setup(t)
	fixture.TestGetAndSetTags(t)
	c.teardown(t)

	fixture = c.setup(t)
	fixture.TestRecordTags(t)
	c.teardown(t)
}
