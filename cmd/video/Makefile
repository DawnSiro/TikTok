server:
	kitex --thrift-plugin validator -module douyin idl/video.thrift # 在项目根目录下执行
	kitex --thrift-plugin validator -module douyin -service video -use douyin/kitex_gen ../../idl/video.thrift # 在 cmd/comment 目录下执行