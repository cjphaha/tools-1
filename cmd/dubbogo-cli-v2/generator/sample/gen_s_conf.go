package sample

const (
	serverConfigFile = `dubbo:
  registries:
    demoZK:
      protocol: zookeeper
      address: 127.0.0.1:2181
  protocols:
    triple:
      name: tri
      port: 20000
  provider:
    services:
      GreeterProvider:
        interface: com.apache.dubbo.sample.basic.IGreeter # must be compatible with grpc or dubbo-java`
)

func init() {
	fileMap["srvConfGenerator"] = &fileGenerator{
		path:    "./go-server/conf",
		file:    "dubbogo.yml",
		context: serverConfigFile,
	}
}
