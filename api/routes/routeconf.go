package routes

import (
	"goji.io"
)

func getMux() *goji.Mux {
	mux := goji.NewMux()
	return mux
}
