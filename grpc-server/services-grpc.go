package grpc_server

type Server struct{
	//empty
}

func (s* Server) DisableOlt(ctx context.Context, empty *pb.Empty) (*pb.Empty, error) {
	log.Println("Info: DisableOlt() called")
	//TODO Send OltDown
}

func (s* Server) ReenableOlt(ctx context.Context, empty *pb.Empty) (*pb.Empty, error) {
	log.Println("Info: ReenableOlt() called")
	//TODO Send 
}
//TODO Implement All Services

