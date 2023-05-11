package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"kazchain/x/kazchain/types"
)

func (k Keeper) ArticleAll(goCtx context.Context, req *types.QueryAllArticleRequest) (*types.QueryAllArticleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var articles []types.Article
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	articleStore := prefix.NewStore(store, types.KeyPrefix(types.ArticleKey))

	pageRes, err := query.Paginate(articleStore, req.Pagination, func(key []byte, value []byte) error {
		var article types.Article
		if err := k.cdc.Unmarshal(value, &article); err != nil {
			return err
		}

		articles = append(articles, article)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllArticleResponse{Article: articles, Pagination: pageRes}, nil
}

func (k Keeper) Article(goCtx context.Context, req *types.QueryGetArticleRequest) (*types.QueryGetArticleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	article, found := k.GetArticle(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetArticleResponse{Article: article}, nil
}
