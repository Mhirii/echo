package internal

func WriteServer(proto Template) string {
	content := "package server" + "\n"
	content += "\n"
	content += writeServerImports(proto)
	content += "\n"
	content += "type Server struct {" + "\n"
	content += "\t" + "pb." + proto.Pkg + "Server" + "\n"
	content += "} " + "\n"
	content += "\n"
	for _, rpc := range proto.Service.Rpcs {
		content += writeEndpoint(rpc)
		content += "\n"
	}

	return content
}

func writeServerImports(proto Template) string {
	imports := []string{
		"\"context\"",
		"\"errors\"",
		"\"github.com/google/uuid\"",
		"\"github.com/gookit/slog\"",
		"pb " + "\"" + proto.Pkg + "/proto" + "\"",
	}

	content := writeImports(imports)
	return content
}

func writeEndpoint(method string) string {
	content := ""
	content += "func (s *Server) " + method + "(ctx context.Context, in *pb." + method + "Request) (*pb." + method + "Response, error) {" + "\n"
	content += "\t" + "return nil, errors.New(" + "\"" + method + " not implemented" + "\"" + ")" + "\n"
	content += "}" + "\n"
	return content
}

func writeImports(imports []string) string {
	content := "import (" + "\n"
	for _, imp := range imports {
		content += "\t" + imp + "\n"
	}
	content += ")\n"
	return content
}

func WriteMainGo(proto Template) string {
	content := "package main" + "\n"
	content += "\n"
	content += writeImports([]string{
		"\"flag\"",
		"\"fmt\"",
		"\"net\"",
		"pb " + "\"" + proto.Pkg + "/proto" + "\"",
		"\"" + proto.Pkg + "/internal/server" + "\"",
		"\"github.com/gookit/slog\"",
		"\"google.golang.org/grpc\"",
	})
	content += "\n"
	content += "var port = flag.Int(\"port\", 50051, \"The server port\")" + "\n"
	content += "\n"

	content += "func main() {" + "\n"
	content += "flag.Parse()" + "\n"

	content += "slog.Info(\"Connecting to Database\")" + "\n"
	content += "\n"

	content += "slog.Info(\"Starting Server\")" + "\n"
	content += "lis, err := net.Listen(\"tcp\", fmt.Sprintf(\":%d\", *port))" + "\n"
	content += "if err != nil {" + "\n"
	content += "slog.Fatalf(\"failed to listen: %v\", err)" + "\n"
	content += "}" + "\n"
	content += "\n"

	content += "s := grpc.NewServer()" + "\n"
	content += "pb.Register" + proto.Pkg + "Server(s, &server.Server{})" + "\n"
	content += "\n"

	content += "slog.Printf(\"server listening at %v\", lis.Addr())" + "\n"
	content += "if err := s.Serve(lis); err != nil {" + "\n"
	content += "slog.Fatalf(\"failed to serve: %v\", err)" + "\n"
	content += "}" + "\n"
	content += "\n"
	content += "}"
	return content
}
