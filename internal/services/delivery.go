package services

import (
	"context"
	"gitlab.com/dh-backend/delivery-service/proto"
)

func (s *Server) DeliveryService(ctx context.Context, in *proto.InDeliveryRequest) (*proto.InDeliveryResponse, error) {

	return &proto.InDeliveryResponse{
		DeliveryFee: 2000,
	}, nil
}
