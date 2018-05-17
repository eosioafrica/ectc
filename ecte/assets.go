package ecte

import (
	"github.com/cavaliercoder/grab"
	"fmt"
	"time"
	"os"
	"net/url"
)

func (env *Environment) DownloadSeedBashInstallAsset() error {

	_, err := url.ParseRequestURI(env.Config.App.Seed)
	if err != nil {

		return WrapErrors(ErrDownloadingBashSeedAsset, err)
	}

	// create client
	client := grab.NewClient()
	req, _ := grab.NewRequest(env.Config.Dirs.AssetsFull,
		env.Config.App.Seed)

	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		return err
	}

	env.SeedBashScript = resp.Filename

	fmt.Printf("Download saved to %v \n", resp.Filename)
	return nil
}