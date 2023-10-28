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

func init() {
	dbClient := db.NewDynamoDBClient()

	boardRepo := repository.NewBoardRepo(dbClient)
	noteRepo := repository.NewNoteRepo(dbClient)

	boardService := service.NewBoardService(boardRepo)
	noteService := service.NewNoteService(noteRepo)

	BoardHandler = handler.NewBoardHandler(boardService)
	NoteHandler = handler.NewNoteHandler(noteService)
}
