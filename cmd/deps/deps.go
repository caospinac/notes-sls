package deps

import (
	"github.com/caospinac/notes-sls/internal/db"
	"github.com/caospinac/notes-sls/internal/handler"
	"github.com/caospinac/notes-sls/internal/repository"
	"github.com/caospinac/notes-sls/internal/service"
)

var (
	BoardHandler handler.BoardHandler
	NoteHandler  handler.NoteHandler
)

const (
	database = "notes-sls"
)

func init() {
	dbClient := db.GetMongoDB(database)

	boardRepo := repository.NewBoardsRepo(dbClient)
	noteRepo := repository.NewNotesRepo(dbClient)

	boardService := service.NewBoardService(boardRepo)
	noteService := service.NewNoteService(noteRepo)

	BoardHandler = handler.NewBoardHandler(boardService)
	NoteHandler = handler.NewNoteHandler(noteService)
}
