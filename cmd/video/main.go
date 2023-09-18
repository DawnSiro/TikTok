package main

import (
	"net"

	"douyin/dal/db"
	"douyin/dal/rdb"
	"douyin/kitex_gen/video/videoservice"
	"douyin/pkg/constant"
	"douyin/pkg/mw"
	"douyin/pkg/viper"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	viper.InitConfig()
	db.InitMySQL()
	rdb.InitRedis()
	klog.SetLogger(logrus.NewLogger())
	klog.SetLevel(klog.LevelInfo)
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constant.ETCDAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr(constant.TCP, constant.VideoServiceAddr)
	if err != nil {
		panic(err)
	}
	Init()
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constant.VideoServiceName),
		provider.WithExportEndpoint(constant.ExportEndpoint),
		provider.WithInsecure(),
	)
	svr := videoservice.NewServer(new(VideoServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.VideoServiceName}),
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
