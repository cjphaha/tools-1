package main

import (
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/config/instance"
	//"dubbo.apache.org/dubbo-go/v3/metadata/service/local"
	_ "dubbo.apache.org/dubbo-go/v3/metadata/service/remote"
	"go.uber.org/zap"
)

var (
	log *zap.SugaredLogger
)

func init() {
	logger, _ := zap.NewDevelopment()
	log = logger.Sugar()
}

func main() {
	//RemoteMateData()
	//LocalMateData()
	Instance()
}

//func LocalMateData()  {
//	lsv, err := local.GetLocalMetadataService()
//	if err != nil {
//		log.Error(err)
//	}
//	_ = lsv
//}
//
//func RemoteMateData() {
//	rsv, err := remote.GetRemoteMetadataService()
//	if err != nil {
//		log.Error(err)
//	}
//	_ = rsv
//}

func Delegate() {
	//report, _ := delegate.NewMetadataReport()
	//report.GetAppMetadata()
}

func Instance() {
	//fmt.Sprintf("dubbo://192.168.1.%v:20000/com.ikurento.user.UserProvider", i)
	url, err := common.NewURL("tri://127.0.0.1:20000/com.apache.dubbo.sample.basic.IGreeter")
	if err != nil {
		log.Error(err)
	}
	mdp := instance.GetMetadataReportInstance(url)
	_ = mdp
}
