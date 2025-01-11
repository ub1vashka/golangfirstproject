package server

import (
	"errors"
	"net/http"

	"github.com/ub1vashka/golangfirstproject/internal/domain/models"
	"github.com/ub1vashka/golangfirstproject/internal/logger"
	"github.com/ub1vashka/golangfirstproject/internal/storage/storageerror"

	"github.com/gin-gonic/gin"
)

func (s *Server) getBookByIDHandler(ctx *gin.Context) {
	log := logger.Get()
	bid := ctx.Param("id")
	book, err := s.bService.GetBook(bid)
	if err != nil {
		log.Error().Err(err).Msg("get book by ID failed")
		if errors.Is(err, storageerror.ErrBookIDNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func (s *Server) deleteBookHandler(ctx *gin.Context) {
	log := logger.Get()
	bid := ctx.Param("id")
	err := s.bService.DeleteBook(bid)
	if err != nil {
		log.Error().Err(err).Msg("delete book failed")
		if errors.Is(err, storageerror.ErrBookIDNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "book deleted successfully"})
}

func (s *Server) addBookHandler(ctx *gin.Context) {
	log := logger.Get()
	var book models.Book
	err := ctx.ShouldBindBodyWithJSON(&book)
	if err != nil {
		log.Error().Err(err).Msg("unmarshall body failed")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bid, err := s.bService.AddBook(book)
	if err != nil {
		log.Error().Err(err).Msg("save book failed")
		if errors.Is(err, storageerror.ErrBookAlredyExist) {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.String(http.StatusCreated, "Book %s was saved", bid)
}

func (s *Server) getBooksHandler(ctx *gin.Context) {
	log := logger.Get()
	books, err := s.bService.GetBooks()
	if err != nil {
		log.Error().Err(err).Msg("get all books form storage failed")
		if errors.Is(err, storageerror.ErrEmptyStorage) {
			ctx.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, books)
}
