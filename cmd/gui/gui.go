package gui

import (
	"fmt"
	"github.com/p9c/pod/pkg/conte"
	"github.com/p9c/pod/pkg/gui/webview"
	"github.com/p9c/pod/pkg/log"
	"github.com/shurcooL/vfsgen"
	"net/http"
	"net/url"
)

func GUI(cx *conte.Xt) {
	rc := rcvar{
		cx:     cx,
		alert:  DuOSalert{},
		status: DuOStatus{},
		txs:    DuOStransactionsExcerpts{},
		lastxs: DuOStransactions{},
	}
	var fs http.FileSystem = http.Dir("./pkg/gui/widgets/CDNSTATIC")
	err := vfsgen.Generate(fs, vfsgen.Options{})
	if err != nil {
		log.FATAL(err)
	}
	rc.fs = fs

	rc.w = webview.New(webview.Settings{
		Width:  1024,
		Height: 760,
		Title:  "ParallelCoin - DUO - True Story",
		URL:    "data:text/html," + url.PathEscape(getFile("/index.html", fs)),
	})

	fmt.Println("dadada", getFile("/index.html", fs))
	//b := Bios{
	//	Theme:      false,
	//	IsBoot:     true,
	//	IsBootMenu: true,
	//	IsBootLogo: true,
	//	IsLoading:  false,
	//	IsDev:      true,
	//	IsScreen:   "overview",
	//}
	log.INFO("starting GUI")

	defer rc.w.Exit()
	rc.w.Dispatch(func() {
		// Load JavaScript Files
		evalJs(&rc)

		// Load CSS files
		injectCss(&rc)
	})
	rc.w.Run()

	//
	//go func() {
	//	for _ = range time.NewTicker(time.Second * 1).C {
	//
	//
	//		//status, err := json.Marshal(rc.GetDuOStatus())
	//		//if err != nil {
	//		//}
	//		//transactions, err := json.Marshal(rc.GetTransactions(0, 555, ""))
	//		//if err != nil {
	//		//}
	//}
	//}()

}
