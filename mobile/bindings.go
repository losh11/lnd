package lndmobile

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/lightningnetwork/lnd"
)

func Start(appDir string, callback Callback) {
	// Call the "real" main in a nested manner so the defers will properly
	// be executed in the case of a graceful shutdown.
	go func() {
		if err := lnd.Main(appDir, bufLightningLis, bufWalletUnlockerLis); err != nil {
			if e, ok := err.(*flags.Error); ok &&
				e.Type == flags.ErrHelp {
			} else {
				fmt.Fprintln(os.Stderr, err)
			}
			os.Exit(1)
		}
	}()

	// TODO(halseth): callback when RPC server is running.
	callback.OnResponse([]byte("started"))
}
