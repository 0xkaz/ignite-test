package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"kazchain/x/kazchain/types"
)

func TestArticleMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateArticle(ctx, &types.MsgCreateArticle{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestArticleMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateArticle
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateArticle{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateArticle{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateArticle{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateArticle(ctx, &types.MsgCreateArticle{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateArticle(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestArticleMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteArticle
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteArticle{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteArticle{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteArticle{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateArticle(ctx, &types.MsgCreateArticle{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteArticle(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
