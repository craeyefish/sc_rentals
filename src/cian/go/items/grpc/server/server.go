package items

import (
	"context"
	"rentals/items"
	pb "rentals/items/grpc"
	item "rentals/items/ops"
	"sync"

	"github.com/shopspring/decimal"
)

type itemsServer struct {
	pb.UnimplementedItemsServer

	mu sync.Mutex // protects routeNotes
}

func (s *itemsServer) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	v, err := decimal.NewFromString(req.Item.GetValue())
	if err != nil {
		return nil, err
	}

	err = item.CreateItem(ctx, items.Item{
		Name:  req.Item.GetName(),
		Value: v,
		Type:  items.Type(req.Item.GetType()),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateItemResponse{Result: true}, nil
}
