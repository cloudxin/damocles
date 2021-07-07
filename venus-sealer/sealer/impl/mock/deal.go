package mock

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/dtynn/venus-cluster/venus-sealer/sealer/api"
)

var _ api.DealManager = (*nullDeal)(nil)

func NewDealManager() api.DealManager {
	return &nullDeal{}
}

type nullDeal struct {
}

func (*nullDeal) Acquire(context.Context, abi.SectorID, *uint) (api.Deals, error) {
	return nil, nil
}