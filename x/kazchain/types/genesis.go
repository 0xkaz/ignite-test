package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ArticleList: []Article{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in article
	articleIdMap := make(map[uint64]bool)
	articleCount := gs.GetArticleCount()
	for _, elem := range gs.ArticleList {
		if _, ok := articleIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for article")
		}
		if elem.Id >= articleCount {
			return fmt.Errorf("article id should be lower or equal than the last id")
		}
		articleIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
