package zk

import (
	"dubbo.apache.org/dubbo-go/v3/metadata/identifier"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetData(t *testing.T) {
	_ = &identifier.SubscriberMetadataIdentifier{
		Revision: "",
		BaseApplicationMetadataIdentifier: identifier.BaseApplicationMetadataIdentifier{
			Application: "com.apache.dubbo.sample.basic.IGreeter",
			Group:       "",
		},
	}
	//info, err := NewZookeeperMetadataReport().GetAppMetadata(id)
	//assert.NoError(t, err)
	//_ = info
	mr := NewZookeeperMetadataReport()
	data, zs, err := mr.client.GetContent("/dubbo/com.apache.dubbo.sample.basic.IGreeter/providers")
	assert.NoError(t, err)
	_ = data
	_ = zs
}
