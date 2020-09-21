package legacy

import (
	"encoding/binary"
	"github.com/stalker-loki/app/slog"

	"github.com/urfave/cli"

	wtxmgr "github.com/stalker-loki/pod/pkg/chain/tx/mgr"
	"github.com/stalker-loki/pod/pkg/db/walletdb"
	"github.com/stalker-loki/pod/pkg/wallet"
)

func DropWalletHistory(w *wallet.Wallet) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		// cfg := w.PodConfig
		var (
			err error
			// Namespace keys.
			syncBucketName    = []byte("sync")
			waddrmgrNamespace = []byte("waddrmgr")
			wtxmgrNamespace   = []byte("wtxmgr")
			// Sync related key names (sync bucket).
			syncedToName     = []byte("syncedto")
			startBlockName   = []byte("startblock")
			recentBlocksName = []byte("recentblocks")
		)
		// dbPath := filepath.Join(*cfg.DataDir,
		// 	*cfg.Network, "wallet.db")
		// Info("dbPath", dbPath)
		// db, err := walletdb.Open("bdb",
		// 	dbPath)
		// if Check(err) {
		// 	// DBError("failed to open database:", err)
		// 	return err
		// }
		// defer db.Close()
		slog.Debug("dropping wtxmgr namespace")
		err = walletdb.Update(w.Database(), func(tx walletdb.ReadWriteTx) error {
			slog.Debug("deleting top level bucket")
			if err := tx.DeleteTopLevelBucket(wtxmgrNamespace); slog.Check(err) {
			}
			if err != nil && err != walletdb.ErrBucketNotFound {
				return err
			}
			var ns walletdb.ReadWriteBucket
			slog.Debug("creating new top level bucket")
			if ns, err = tx.CreateTopLevelBucket(wtxmgrNamespace); slog.Check(err) {
				return err
			}
			if err = wtxmgr.Create(ns); slog.Check(err) {
				return err
			}
			ns = tx.ReadWriteBucket(waddrmgrNamespace).NestedReadWriteBucket(syncBucketName)
			startBlock := ns.Get(startBlockName)
			slog.Debug("putting start block", startBlock)
			if err = ns.Put(syncedToName, startBlock); slog.Check(err) {
				return err
			}
			recentBlocks := make([]byte, 40)
			copy(recentBlocks[0:4], startBlock[0:4])
			copy(recentBlocks[8:], startBlock[4:])
			binary.LittleEndian.PutUint32(recentBlocks[4:8], uint32(1))
			defer slog.Debug("put recent blocks")
			return ns.Put(recentBlocksName, recentBlocks)
		})
		if slog.Check(err) {
			return err
		}
		slog.Debug("updated wallet")
		// if w != nil {
		// 	// Rescan chain to ensure balance is correctly regenerated
		// 	job := &wallet.RescanJob{
		// 		InitialSync: true,
		// 	}
		// 	// Submit rescan job and log when the import has completed.
		// 	// Do not block on finishing the rescan.  The rescan success
		// 	// or failure is logged elsewhere, and the channel is not
		// 	// required to be read, so discard the return value.
		// 	errC := w.SubmitRescan(job)
		// 	select {
		// 	case err := <-errC:
		// 		DBError(err)
		// 		// case <-time.After(time.Second * 5):
		// 		// 	break
		// 	}
		// }
		return err
	}
}
