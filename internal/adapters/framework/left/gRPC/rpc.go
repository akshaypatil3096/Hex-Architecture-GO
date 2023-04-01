package grpc

import (
	"context"

	"github.com/akshaypatil3096/Hex-Architecture-GO/internal/adapters/framework/left/gRPC/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetAddition(int(req.A), int(req.B))
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpecter error")
	}

	ans = &pb.Answer{
		Value: int32(answer),
	}

	return ans, nil
}

func (grpca Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetSubtraction(int(req.A), int(req.B))
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpecter error")
	}

	ans = &pb.Answer{
		Value: int32(answer),
	}

	return ans, nil
}

func (grpca Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetMultiplication(int(req.A), int(req.B))
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpecter error")
	}

	ans = &pb.Answer{
		Value: int32(answer),
	}

	return ans, nil
}

func (grpca Adapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetDivision(int(req.A), int(req.B))
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpecter error")
	}

	ans = &pb.Answer{
		Value: int32(answer),
	}

	return ans, nil
}
