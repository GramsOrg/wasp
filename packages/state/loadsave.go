package state

import (
	"errors"
	"github.com/iotaledger/hive.go/kvstore"
	"github.com/iotaledger/wasp/packages/database/dbkeys"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/trie"
	"github.com/iotaledger/wasp/packages/util"
	"golang.org/x/xerrors"
)

type mustKVStoreBatch struct {
	prefix byte
	batch  kvstore.BatchedMutations
}

func newKVStoreBatch(prefix byte, batch kvstore.BatchedMutations) *mustKVStoreBatch {
	return &mustKVStoreBatch{
		prefix: prefix,
		batch:  batch,
	}
}

func (k *mustKVStoreBatch) Set(key kv.Key, value []byte) {
	if err := k.batch.Set(dbkeys.MakeKey(k.prefix, []byte(key)), value); err != nil {
		panic(err)
	}
}

func (k *mustKVStoreBatch) Del(key kv.Key) {
	if err := k.batch.Delete(dbkeys.MakeKey(k.prefix, []byte(key))); err != nil {
		panic(err)
	}
}

// Save saves updates collected in the virtual state to DB together with the provided blocks in one transaction
// Mutations must be non-empty otherwise it is NOP
func (vs *virtualStateAccess) Save(blocks ...Block) error {
	if vs.kvs.Mutations().IsEmpty() {
		// nothing to commit
		return nil
	}
	vs.Commit()

	batch := vs.db.Batched()

	vs.trie.PersistMutations(newKVStoreBatch(dbkeys.ObjectTypeTrie, batch))
	vs.kvs.Mutations().Apply(newKVStoreBatch(dbkeys.ObjectTypeState, batch))
	for _, blk := range blocks {
		key := dbkeys.MakeKey(dbkeys.ObjectTypeBlock, util.Uint32To4Bytes(blk.BlockIndex()))
		if err := batch.Set(key, blk.Bytes()); err != nil {
			return err
		}
	}
	if err := batch.Commit(); err != nil {
		return err
	}

	// call flush explicitly, because batched.Commit doesn't actually write the changes to disk
	if err := vs.db.Flush(); err != nil {
		return err
	}

	vs.trie.ClearCache()
	vs.kvs.ClearMutations()
	vs.kvs.Mutations().ResetModified()
	return nil
}

// LoadSolidState establishes VirtualStateAccess interface with the solid state in DB.
// Checks root commitment to chainID
func LoadSolidState(store kvstore.KVStore, chainID *iscp.ChainID) (VirtualStateAccess, bool, error) {
	// check the existence of terminalCommitment at key ''. chainID is expected
	v, err := store.Get(dbkeys.MakeKey(dbkeys.ObjectTypeState))
	if errors.Is(err, kvstore.ErrKeyNotFound) {
		// state does not exist
		return nil, false, nil
	}
	if err != nil {
		return nil, false, xerrors.Errorf("LoadSolidState: %v", err)
	}
	chID, err := iscp.ChainIDFromBytes(v)
	if err != nil {
		return nil, false, xerrors.Errorf("LoadSolidState: %v", err)
	}
	if !chID.Equals(chainID) {
		return nil, false, xerrors.Errorf("LoadSolidState: expected chainID: %s, got: %s", chainID, chID)
	}
	ret := newVirtualState(store)

	// explicit use of merkle trie model. Asserting that the chainID is committed by the root at the key ''
	merkleProof := CommitmentModel.Proof(nil, ret.trie)
	if err = merkleProof.Validate(trie.RootCommitment(ret.trie), chainID.Bytes()); err != nil {
		return nil, false, xerrors.Errorf("LoadSolidState: can't prove inclusion of chain ID %s in the root: %v", chainID, err)
	}
	ret.kvs.Mutations().ResetModified()
	return ret, true, nil
}

// LoadBlockBytes loads block bytes of the specified block index from DB
func LoadBlockBytes(store kvstore.KVStore, stateIndex uint32) ([]byte, error) {
	data, err := store.Get(dbkeys.MakeKey(dbkeys.ObjectTypeBlock, util.Uint32To4Bytes(stateIndex)))
	if errors.Is(err, kvstore.ErrKeyNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

// LoadBlock loads block from DB and decodes it
func LoadBlock(store kvstore.KVStore, stateIndex uint32) (Block, error) {
	data, err := LoadBlockBytes(store, stateIndex)
	if err != nil {
		return nil, err
	}
	return BlockFromBytes(data)
}
