package grpchandler

import (
	"context"

	pb "github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/pb/rest_api_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type StorygRPCHandler struct {
	pb.UnimplementedStoryServiceServer
}

func NewStorygRPCHandler() pb.StoryServiceServer {
	return &StorygRPCHandler{}
}

func (h *StorygRPCHandler) FindAll(context.Context, *pb.StoryFindAllRequest) (*pb.Stories, error) {
	stories := &pb.Stories{
		Stories: []*pb.Story{
			{
				Id:      1,
				Title:   "Example",
				Content: "Example",
			},
		},
	}

	return stories, nil
}

func (h *StorygRPCHandler) FindByID(context.Context, *pb.StoryByIDRequest) (*pb.Story, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByID not implemented")
}

func (h *StorygRPCHandler) Create(context.Context, *pb.StoryCreateRequest) (*pb.Story, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func (h *StorygRPCHandler) Update(context.Context, *pb.StoryUpdateRequest) (*pb.Story, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func (h *StorygRPCHandler) Delete(context.Context, *pb.StoryByIDRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
