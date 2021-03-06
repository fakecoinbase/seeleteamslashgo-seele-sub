/**
*  @file
*  @copyright defined in go-seele/LICENSE
 */

package core

import (
	"math/big"

	"github.com/seeleteam/go-seele/common"
	"github.com/seeleteam/go-seele/consensus/pow"
	"github.com/seeleteam/go-seele/core/store"
	"github.com/seeleteam/go-seele/core/types"
	"github.com/seeleteam/go-seele/database/leveldb"
)

func newTestGenesis() *Genesis {
	accounts := map[common.Address]*big.Int{
		types.TestGenesisAccount.Addr: types.TestGenesisAccount.Amount,
	}

	return GetGenesis(NewGenesisInfo(accounts, 1, 0, big.NewInt(0), types.PowConsensus, nil))
}

func NewTestBlockchain() *Blockchain {
	return NewTestBlockchainWithVerifier(nil)
}

func NewTestBlockchainWithVerifier(verifier types.DebtVerifier) *Blockchain {
	db1, _ := leveldb.NewTestDatabase()
	db2, _ := leveldb.NewTestDatabase()
	db3, _ := leveldb.NewTestDatabase()

	bcStore := store.NewCachedStore(store.NewBlockchainDatabase(db1))

	genesis := newTestGenesis()
	if err := genesis.InitializeAndValidate(bcStore, db1); err != nil {
		panic(err)
	}

	bc, err := NewBlockchain(bcStore, db1, db2, db3, "", pow.NewEngine(1), verifier, -1)
	if err != nil {
		panic(err)
	}

	return bc
}
