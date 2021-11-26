package zk

import (
	"encoding/json"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/metadata/identifier"
	gxzookeeper "github.com/dubbogo/gost/database/kv/zk"
)

// ZookeeperMetadataReport is the implementation of
// MetadataReport based on zookeeper.
type ZookeeperMetadataReport struct {
	client  *gxzookeeper.ZookeeperClient
	rootDir string
}

func NewZookeeperMetadataReport() *ZookeeperMetadataReport {
	client, err := gxzookeeper.NewZookeeperClient("demoZK", []string{"127.0.0.1:2181"}, true)
	if err != nil {
		panic(err)
	}
	return &ZookeeperMetadataReport{
		client:  client,
		rootDir: "/dubbo/",
	}
}

// GetAppMetadata get metadata info from zookeeper
func (m *ZookeeperMetadataReport) GetAppMetadata(metadataIdentifier *identifier.SubscriberMetadataIdentifier) (*common.MetadataInfo, error) {
	k := m.rootDir + metadataIdentifier.GetFilePathKey()
	data, _, err := m.client.GetContent(k)
	if err != nil {
		return nil, err
	}
	var metadataInfo common.MetadataInfo
	err = json.Unmarshal(data, &metadataInfo)
	if err != nil {
		return nil, err
	}
	return &metadataInfo, nil
}
