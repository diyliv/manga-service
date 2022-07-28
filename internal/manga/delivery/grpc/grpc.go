package grpc

import (
	"context"

	"github.com/diyliv/manga-service/pkg/parser"
	mangapb "github.com/diyliv/manga-service/proto/manga"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mangaService struct {
	logger *zap.Logger
}

func NewMangaService(logger *zap.Logger) *mangaService {
	return &mangaService{logger: logger}
}

func (s *mangaService) Search(ctx context.Context, req *mangapb.SearchReq) (*mangapb.SearchResp, error) {
	name := req.GetMangaName()

	if name == "" {
		return nil, status.Error(codes.InvalidArgument, "You didnt specify a name.")
	}

	parse := parser.ParseMangaPoisk(name)
	if len(parse) == 0 {
		return nil, status.Error(codes.NotFound, "Nothing was found.")
	}

	resp := make([]*mangapb.Manga, 0)

	for _, v := range parse {
		if v.Rate == "" {
			v.Rate = "Отзывов нет"
		}

		resp = append(resp, &mangapb.Manga{
			Name:        v.Name,
			Link:        v.Link,
			Genre:       v.Genre,
			Chapters:    v.Chapters,
			Description: v.Description,
			Year:        v.Year,
			Rate:        v.Rate,
		})
	}

	return &mangapb.SearchResp{SearchResp: resp}, nil
}
