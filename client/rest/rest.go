package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/irismod/record/types"
)

// Rest variable names
// nolint
const (
	RestRecordID = "record-id"
)

// RegisterHandle defines routes that get registered by the main application
func RegisterHandle(cliCtx client.Context, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}

type RecordCreateReq struct {
	BaseReq  rest.BaseReq    `json:"base_req" yaml:"base_req"` // base req
	Contents []types.Content `json:"contents" yaml:"contents"`
	Creator  sdk.AccAddress  `json:"creator" yaml:"creator"`
}
