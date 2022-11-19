// Command remote is a chromedp example demonstrating how to connect to an
// existing Chrome DevTools instance using a remote WebSocket URL.
package main

import (
	"context"
	"flag"
	"log"
	// "time"
	"os"

	"github.com/chromedp/chromedp"
)

func main() {
	devtoolsWsURL := flag.String("url", "", "DevTools WebSsocket URL/IP and port e.g. IP:9223")
	meetId := flag.String("id", "", "Google Meet id")

	flag.Parse()

	if *devtoolsWsURL == "" {
		log.Fatal("must specify -url")
	}

	if *meetId == "" {
		log.Fatal("must specify -id")
	}

	// create allocator context for use with creating a browser context later
	allocatorContext, cancel := chromedp.NewRemoteAllocator(context.Background(), *devtoolsWsURL)
	defer cancel()

	// create context
	ctx, cancel := chromedp.NewContext(allocatorContext)
	defer cancel()

	// run task list
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://meet.google.com/" + *meetId), // nkf-kfkr-wne
		// chromedp.WaitVisible(`body > c-wiz`),
		// chromedp.Sleep(2 * time.Second),
		chromedp.Click(`//*/c-wiz/div/div/div[13]/div[3]/div/div[1]/div[4]/div/div/div[2]/div/div[2]/div/div[1]/div[1]/button/span`,chromedp.NodeVisible),
		// chromedp.Sleep(3600 * time.Second),
	); err != nil {
		log.Fatalf("Failed getting body of duckduckgo.com: %v", err)
	}
	os.Exit(0)
}