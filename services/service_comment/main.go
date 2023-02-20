package main

//func main() {
//	// 监听本地 5678 端口
//	listen, err := net.Listen("tcp", ":5678")
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//
//	// 创建 gRPC 服务器
//	s := grpc.NewServer()
//	// 将实现的接口注册进 gRPC 服务器
//	model.RegisterUserSrvServer(s, &controller.Server{})
//	log.Println("gRPC server starts running...")
//	// 启动 gRPC 服务器
//	err = s.Serve(listen)
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//}
