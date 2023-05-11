package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"kazchain/x/kazchain/types"
)

// GetArticleCount get the total number of article
func (k Keeper) GetArticleCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ArticleCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetArticleCount set the total number of article
func (k Keeper) SetArticleCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ArticleCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendArticle appends a article in the store with a new id and update the count
func (k Keeper) AppendArticle(
	ctx sdk.Context,
	article types.Article,
) uint64 {
	// Create the article
	count := k.GetArticleCount(ctx)

	// Set the ID of the appended value
	article.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ArticleKey))
	appendedValue := k.cdc.MustMarshal(&article)
	store.Set(GetArticleIDBytes(article.Id), appendedValue)

	// Update article count
	k.SetArticleCount(ctx, count+1)

	return count
}

// SetArticle set a specific article in the store
func (k Keeper) SetArticle(ctx sdk.Context, article types.Article) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ArticleKey))
	b := k.cdc.MustMarshal(&article)
	store.Set(GetArticleIDBytes(article.Id), b)
}

// GetArticle returns a article from its id
func (k Keeper) GetArticle(ctx sdk.Context, id uint64) (val types.Article, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ArticleKey))
	b := store.Get(GetArticleIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveArticle removes a article from the store
func (k Keeper) RemoveArticle(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ArticleKey))
	store.Delete(GetArticleIDBytes(id))
}

// GetAllArticle returns all article
func (k Keeper) GetAllArticle(ctx sdk.Context) (list []types.Article) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ArticleKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Article
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetArticleIDBytes returns the byte representation of the ID
func GetArticleIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetArticleIDFromBytes returns ID in uint64 format from a byte array
func GetArticleIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
