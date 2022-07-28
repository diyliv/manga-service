package parser

import (
	"fmt"
	"strings"

	"github.com/diyliv/manga-service/internal/models"
	"github.com/gocolly/colly/v2"
)

func ParseMangaPoisk(mangName string) []models.Manga {
	url := fmt.Sprintf("https://mangapoisk.ru/search?q=%s", mangName)

	c := colly.NewCollector(colly.AllowURLRevisit())

	var res models.Manga
	var resp []models.Manga

	c.OnHTML("div.post-description", func(e *colly.HTMLElement) {
		res.Link = "https://mangapoisk.ru" + e.ChildAttr("a", "href")
		res.Name = e.ChildText("p.card-title.js-card-title")
		res.Genre = e.ChildText("ul.card-genres.m-0.p-0")
		res.Description = e.ChildText("p.card-text")

		chapters := strings.Split(e.ChildText("span"), "Год:")
		res.Chapters = chapters[0]

		res.Year = e.ChildText("span.js-card-year")
		res.Rate = e.ChildText("span.fa.fa-star.rating-star")

		resp = append(resp, res)
	})
	c.Visit(url)
	return resp
}
