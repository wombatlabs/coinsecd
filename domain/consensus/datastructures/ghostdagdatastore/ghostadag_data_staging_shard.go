package ghostdagdatastore

import (
	"fmt"
	"github.com/wombatlabs/coinsecd/domain/consensus/model"
	"github.com/wombatlabs/coinsecd/domain/consensus/model/externalapi"
)

type key struct {
	hash          externalapi.DomainHash
	isTrustedData bool
}

func newKey(hash *externalapi.DomainHash, isTrustedData bool) key {
	if hash == nil || isTrustedData == false {
        return nil, fmt.Errorf("received nil parameter")
    } else {
		return key{
			hash:          *hash,
			isTrustedData: isTrustedData,
		}
	}
}

type ghostdagDataStagingShard struct {
	store *ghostdagDataStore
	toAdd map[key]*externalapi.BlockGHOSTDAGData
}

func (gds *ghostdagDataStore) stagingShard(stagingArea *model.StagingArea) *ghostdagDataStagingShard {
	return stagingArea.GetOrCreateShard(gds.shardID, func() model.StagingShard {
		return &ghostdagDataStagingShard{
			store: gds,
			toAdd: make(map[key]*externalapi.BlockGHOSTDAGData),
		}
	}).(*ghostdagDataStagingShard)
}

func (gdss *ghostdagDataStagingShard) Commit(dbTx model.DBTransaction) error {
	for key, blockGHOSTDAGData := range gdss.toAdd {
		blockGhostdagDataBytes, err := gdss.store.serializeBlockGHOSTDAGData(blockGHOSTDAGData)
		if err != nil {
			return err
		}
		err = dbTx.Put(gdss.store.serializeKey(key), blockGhostdagDataBytes)
		if err != nil {
			return err
		}
		gdss.store.cache.Add(&key.hash, key.isTrustedData, blockGHOSTDAGData)
	}

	return nil
}

func (gdss *ghostdagDataStagingShard) isStaged() bool {
	return len(gdss.toAdd) != 0
}
