package environment

import (
	"os"
	"gopkg.in/src-d/go-git.v4"
)

func (env *Environment) PullAssetsFromGit(){

	if env.Err != nil { return  }

	_, err := git.PlainClone(env.Config.Dirs.AssetsFull, false, &git.CloneOptions{
		URL:      "https://github.com/src-d/go-git",
		Progress: os.Stdout,
	})

	env.Err = WrapErrors(ErrPullingAssetsFromGit, err)

}

/*
func (env *Environment) DownloadSeedBashInstallAsset() {

	if env.Err != nil { return  }

	_, err := url.ParseRequestURI(env.Config.App.Seed)
	if err != nil {

		env.Err = WrapErrors(ErrDownloadingBashSeedAsset, err)
		return
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
		env.Err = WrapErrors(ErrDownloadingBashSeedAsset, err)
	}

	env.SeedBashScript = resp.Filename

	fmt.Printf("Download saved to %v \n", env.SeedBashScript)

	return
}*/