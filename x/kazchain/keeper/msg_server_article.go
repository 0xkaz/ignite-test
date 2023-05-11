package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"kazchain/x/kazchain/types"
)

func (k msgServer) CreateArticle(goCtx context.Context, msg *types.MsgCreateArticle) (*types.MsgCreateArticleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var article = types.Article{
		Creator: msg.Creator,
		Name:    msg.Name,
		Value:   msg.Value,
	}

	id := k.AppendArticle(
		ctx,
		article,
	)

	return &types.MsgCreateArticleResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateArticle(goCtx context.Context, msg *types.MsgUpdateArticle) (*types.MsgUpdateArticleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var article = types.Article{
		Creator: msg.Creator,
		Id:      msg.Id,
		Name:    msg.Name,
		Value:   msg.Value,
	}

	// Checks that the element exists
	val, found := k.GetArticle(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetArticle(ctx, article)

	return &types.MsgUpdateArticleResponse{}, nil
}

func (k msgServer) DeleteArticle(goCtx context.Context, msg *types.MsgDeleteArticle) (*types.MsgDeleteArticleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetArticle(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveArticle(ctx, msg.Id)

	return &types.MsgDeleteArticleResponse{}, nil
}
