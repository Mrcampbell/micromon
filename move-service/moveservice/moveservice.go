package moveservice

import (
	"context"
	"fmt"

	"github.com/Mrcampbell/pgo2/move-service/moveservice/convert"
	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
)

var _ pokemon.MoveServiceServer = &MoveService{}

type MoveService struct {
	set map[string]pokemon.MoveDetail
}

func NewMoveService() *MoveService {

	set := make(map[string]pokemon.MoveDetail, 0)

	ms := &MoveService{
		set: set,
	}

	ms.load()
	return ms
}

func (ms *MoveService) GetMoveDetail(ctx context.Context, req *pokemon.GetMoveDetailRequest) (*pokemon.GetMoveDetailResponse, error) {
	move, ok := ms.set[req.Id]
	if !ok {
		return nil, fmt.Errorf("GetMoveDetail: Requested Move ID not found in Records: %v", req.Id)
	}
	return &pokemon.GetMoveDetailResponse{
		Move: &move,
	}, nil
}

func (ms *MoveService) GetMoveSummary(ctx context.Context, req *pokemon.GetMoveSummaryRequest) (*pokemon.GetMoveSummaryResponse, error) {
	move, ok := ms.set[req.Id]
	if !ok {
		fmt.Printf("GetMoveDetail: Requested Move ID not found in Records: %v\n", req.Id)
		return &pokemon.GetMoveSummaryResponse{
			Move: &pokemon.MoveSummary{
				Id:   req.Id,
				Name: "UNKNOWN",
				Type: pokemon.Type_UNKNOWN,
			},
		}, nil
	}

	summary := convert.MoveDetailToSummary(move)
	return &pokemon.GetMoveSummaryResponse{
		Move: &summary,
	}, nil
}
