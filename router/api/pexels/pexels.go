package pexels

import (
	"github.com/lyeka/gopexels"
	"os"
)

var pc *gopexels.Client

func init() {
	pc = gopexels.NewClient(os.Getenv("PexelsToken"))
}

