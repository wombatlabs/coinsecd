package common

import (
	"fmt"
	"github.com/coinsec/coinsecd/domain/dagconfig"
	"os"
	"sync/atomic"
	"syscall"
	"testing"
)

// RunCoinsecdForTesting runs coinsecd for testing purposes
func RunCoinsecdForTesting(t *testing.T, testName string, rpcAddress string) func() {
	appDir, err := TempDir(testName)
	if err != nil {
		t.Fatalf("TempDir: %s", err)
	}

	coinsecdRunCommand, err := StartCmd("COINSECD",
		"coinsecd",
		NetworkCliArgumentFromNetParams(&dagconfig.DevnetParams),
		"--appdir", appDir,
		"--rpclisten", rpcAddress,
		"--loglevel", "debug",
	)
	if err != nil {
		t.Fatalf("StartCmd: %s", err)
	}
	t.Logf("Coinsecd started with --appdir=%s", appDir)

	isShutdown := uint64(0)
	go func() {
		err := coinsecdRunCommand.Wait()
		if err != nil {
			if atomic.LoadUint64(&isShutdown) == 0 {
				panic(fmt.Sprintf("Coinsecd closed unexpectedly: %s. See logs at: %s", err, appDir))
			}
		}
	}()

	return func() {
		err := coinsecdRunCommand.Process.Signal(syscall.SIGTERM)
		if err != nil {
			t.Fatalf("Signal: %s", err)
		}
		err = os.RemoveAll(appDir)
		if err != nil {
			t.Fatalf("RemoveAll: %s", err)
		}
		atomic.StoreUint64(&isShutdown, 1)
		t.Logf("Coinsecd stopped")
	}
}
