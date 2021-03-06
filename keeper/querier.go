package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/irismod/record/types"
)

// NewQuerier creates a querier for record REST endpoints
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryRecord:
			return queryRecord(ctx, k, req)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown query path: %s", path[0])
		}
	}
}

func queryRecord(ctx sdk.Context, k Keeper, req abci.RequestQuery) ([]byte, error) {
	var params types.QueryRecordParams
	err := k.cdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	record, found := k.GetRecord(ctx, params.RecordID)
	if !found {
		return nil, types.ErrUnknownRecord
	}

	recordOutput := types.RecordOutput{
		TxHash:   record.TxHash.String(),
		Contents: record.Contents,
		Creator:  record.Creator,
	}

	bz, err := codec.MarshalJSONIndent(k.cdc, recordOutput)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
