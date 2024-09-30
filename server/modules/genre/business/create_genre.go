package business

import (
	"context"
	"io"
	"strings"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/genre/model"
	"github.com/gosimple/slug"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type CreateGenreStorage interface {
	CreateGenre(ctx context.Context, data *model.GenreCreation) error
}

type createGenreBiz struct {
	storage CreateGenreStorage
}

func NewCreateGenreBiz(storage CreateGenreStorage) *createGenreBiz {
	return &createGenreBiz{storage}
}

func (biz *createGenreBiz) CreateGenre(ctx context.Context, data *model.GenreCreation) error {
	name := strings.TrimSpace(data.GenreName)

	if name == "" {
		return model.ErrGenreNameIsBlank
	}

	data.GenreSlug = generateSlug(data.GenreName)

	if err := biz.storage.CreateGenre(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(err, model.GenreEntityName)
	}

	return nil
}

func generateSlug(input string) string {
	// Normalize the string
	t := transform.NewReader(strings.NewReader(input), norm.NFD)
	normalized, _ := io.ReadAll(t)

	// Convert to slug
	return slug.Make(string(normalized))
}
