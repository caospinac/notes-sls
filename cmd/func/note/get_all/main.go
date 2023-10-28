package main

import (
	"github.com/caospinac/notes-sls/cmd/deps"
	"github.com/caospinac/notes-sls/pkg/util"
)

func main() {
	util.Start(deps.NoteHandler.GetAll)
}
