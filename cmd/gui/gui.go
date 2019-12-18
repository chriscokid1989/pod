package gui

import (
	"gioui.org/app"
	"github.com/p9c/pod/cmd/gui/duoui"
	"github.com/p9c/pod/pkg/log"
)

func WalletGUI(duo *duoui.DuoUI) (err error) {
	go func() {
		if err := duoui.DuoUImainLoop(duo); err != nil {
			log.FATAL(err)
		}
	}()
	app.Main()
	return
}
