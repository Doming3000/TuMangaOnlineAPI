package tumangaonline

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	s "strings"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"github.com/julioolivares90/TumangaOnlineApi/models"
	"github.com/julioolivares90/TumangaOnlineApi/utilities"
)

// Limpiar texto para facilitar la lectura
func LimpiarTexto(texto string) string {
	// Reemplazar saltos de línea y tabulaciones por espacios
	texto = s.ReplaceAll(texto, "\n", " ")
	texto = s.ReplaceAll(texto, "\t", " ")

	// Eliminar múltiples espacios
	re := regexp.MustCompile(`\s+`)
	texto = re.ReplaceAllString(texto, " ")

	// Eliminar espacios al inicio y al final
	return s.TrimSpace(texto)
}

// GetInfoManga obtiene la informacion de un manga
func GetInfoManga(url string) models.MangaInfoTMO {
	c := colly.NewCollector()
	mangaInfo := models.MangaInfoTMO{}

	c.OnHTML("#app > section", func(element *colly.HTMLElement) {
		// Extraer información general del manga
		mangaInfo.Title = LimpiarTexto(element.ChildText("header > section.element-header-content > div.container.h-100 > div > div.col-12.col-md-9.element-header-content-text > h1"))
		mangaInfo.Image = element.ChildAttr("header > section.element-header-content > div.container.h-100 > div > div.col-12.col-md-3.text-center > div > img", "src")
		mangaInfo.Type = element.ChildText("header > section.element-header-content > div.container.h-100 > div > div.col-12.col-md-3.text-center > h1")
		mangaInfo.Score = element.ChildText("header > section.element-header-content > div.container.h-100 > div > div.col-12.col-md-3.text-center > div > div.score > a > span")
		mangaInfo.Demography = element.ChildText("header > section.element-header-content > div.container > div.row > div.col-12 > div.element-image > div.demography")
		mangaInfo.Synopsis = LimpiarTexto(element.ChildText("header > section.element-header-content > div.container.h-100 > div > div.col-12.col-md-9.element-header-content-text > p.element-description"))
		mangaInfo.Status = element.ChildText("header > section.element-header-content > div.container.h-100 > div > div.col-12.col-md-9.element-header-content-text > span.book-status")

		// Extraer lista de generos
		var generos []string
		element.ForEach("header > section > div.container > div.row > div.col-12 > h6", func(i int, el *colly.HTMLElement) {
			generos = append(generos, LimpiarTexto(el.Text))
		})
		mangaInfo.Genres = generos

		// Extraer lista de capitulos
		var capitulos []models.Chapter

		element.ForEach("#chapters > ul.list-group > li, #chapters > ul.list-group > div > li", func(i int, el *colly.HTMLElement) {
			cap := models.Chapter{
				Title: LimpiarTexto(el.ChildText("h4 > div.row > div > a.btn-collapse")),
			}
			capitulos = append(capitulos, cap)
		})

		// Asignar lista final de capítulos
		mangaInfo.Chapters = capitulos
	})
	c.Visit(url)
	return mangaInfo
}

// Obtener la URL de un capitulo por su número
func FindChapterURLByNumber(mangaURL, capNumber string) ([]models.ChapterData, error) {
	var results []models.ChapterData

	userCapNum, err := parseChapterNumber(capNumber)
	if err != nil {
		return nil, fmt.Errorf("capítulo inválido: %w", err)
	}

	c := colly.NewCollector()

	c.OnHTML("li.upload-link", func(e *colly.HTMLElement) {
		chapterTitle := e.ChildText("a.btn-collapse")
		chNum, _ := extractChapterNumber(chapterTitle)
		if chNum == 0 || chNum != userCapNum {
			return
		}
		title := s.TrimSpace(chapterTitle)

		// Recorre todos los scans de este capítulo
		e.ForEach("ul.chapter-list > li.list-group-item", func(_ int, scanEl *colly.HTMLElement) {
			scan := s.TrimSpace(scanEl.ChildText("div.col-4.col-md-6.text-truncate span a"))
			date := s.TrimSpace(scanEl.ChildText("div.col-4.col-md-2.text-center span.badge-primary"))
			readURL := scanEl.ChildAttr("div.col-2.col-sm-1.text-right a.btn.btn-default.btn-sm", "href")
			if readURL == "" {
				return
			}
			if !s.HasPrefix(readURL, "http") {
				readURL = "https://zonatmo.com" + readURL
			}
			viewerURL, err := ResolveChapterURL(readURL)
			if err != nil {
				viewerURL = readURL
			}
			results = append(results, models.ChapterData{
				Title:   title,
				UrlRead: viewerURL,
				Scan:    scan,
				Date:    extractDate(date),
			})
		})
	})

	if err := c.Visit(mangaURL); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, errors.New("capítulo no encontrado")
	}

	return results, nil
}

// Extraer el número de capítulo como float64 desde el texto, ej: "Capítulo 1.00  ..." => 1.0
func extractChapterNumber(text string) (float64, error) {
	re := regexp.MustCompile(`Cap[ií]tulo\s+([0-9]+(?:\.[0-9]+)?)`)
	matches := re.FindStringSubmatch(text)
	if len(matches) < 2 {
		return 0, errors.New("no se encontró número de capítulo")
	}
	return strconv.ParseFloat(matches[1], 64)
}

// Parsear el número de capítulo como float64
func parseChapterNumber(input string) (float64, error) {
	return strconv.ParseFloat(input, 64)
}

// Extraer la fecha de un string
func extractDate(dateRaw string) string {
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	match := re.FindString(dateRaw)
	return match
}

// Seguir la redirección y devuelve la URL final del visor.
func ResolveChapterURL(initialURL string) (string, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest("GET", initialURL, nil)
	if err != nil {
		return "", fmt.Errorf("error creando la solicitud HTTP: %w", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9")
	req.Header.Set("Referer", "https://zonatmo.com")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error al hacer la solicitud HTTP: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 && resp.StatusCode < 400 {
		location := resp.Header.Get("Location")
		if location == "" {
			return "", errors.New("la redirección no contiene encabezado Location")
		}
		if s.HasPrefix(location, "/") {
			return "https://zonatmo.com" + location, nil
		}
		return location, nil
	}

	if resp.StatusCode == 200 {
		return initialURL, nil
	}

	return "", fmt.Errorf("error en la solicitud: %d", resp.StatusCode)
}

// Obtener los mangas populares
func GetMangasPopulares(pageNumber int) []models.MangaTMO {
	var mangasPopulares []models.MangaTMO
	url := utilities.TUMANGAONLINE_BASE_URL
	c := colly.NewCollector()

	c.OnHTML("#app > main > div:nth-child(2) > div.col-12.col-lg-8.col-xl-9 > div:nth-child(1)", func(element *colly.HTMLElement) {
		element.ForEach("div.element", func(i int, element *colly.HTMLElement) {
			dataItentificador := element.Attr("data-identifier")
			//fmt.Println(dataItentificador)
			mangaPopular := models.MangaTMO{
				Title:      element.ChildText("a > div > div > h4"),
				MangaUrl:   element.ChildAttr("a", "href"),
				Type:       element.ChildText("a > div > span.book-type"),
				Demography: element.ChildText("a > div > span.demography"),
				Score:      element.ChildText("a > div > span.score > span"),
				MangaImage: getImagenManga(element.ChildText("a > div > style"), dataItentificador),
			}
			mangasPopulares = append(mangasPopulares, mangaPopular)

		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.Visit(fmt.Sprintf("%s/populars?page=%d", url, pageNumber))
	return mangasPopulares
}

// GetMangasPopularesJosei
func GetMangasPopularesJosei() []models.MangaTMO {
	var mangasPopulares []models.MangaTMO
	url := utilities.TUMANGAONLINE_BASE_URL
	c := colly.NewCollector()

	c.OnHTML("#app > main > div:nth-child(2) > div.col-12.col-lg-8.col-xl-9 > div:nth-child(1)", func(element *colly.HTMLElement) {
		element.ForEach("div.element", func(i int, element *colly.HTMLElement) {
			dataItentificador := element.Attr("data-identifier")
			//fmt.Println(dataItentificador)
			mangaPopular := models.MangaTMO{
				Title:      element.ChildText("a > div > div > h4"),
				MangaUrl:   element.ChildAttr("a", "href"),
				Type:       element.ChildText("a > div > span.book-type"),
				Demography: element.ChildText("a > div > span.demography"),
				Score:      element.ChildText("a > div > span.score > span"),
				MangaImage: getImagenManga(element.ChildText("a > div > style"), dataItentificador),
			}
			mangasPopulares = append(mangasPopulares, mangaPopular)

		})
	})
	c.Visit(fmt.Sprintf("%s/", url))
	return mangasPopulares
}

// GetMangasPopularesSeinen
func GetMangasPopularesSeinen() []models.MangaTMO {
	var mangasPopulares []models.MangaTMO
	url := utilities.TUMANGAONLINE_BASE_URL

	// initializing a chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// navigate to the target web page and select the HTML elements of interest
	var nodes []*cdp.Node
	chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Nodes("body", &nodes, chromedp.ByQueryAll),
	)

	var url2, image, name, price string
	for _, node := range nodes {

		chromedp.Run(ctx,
			chromedp.AttributeValue("a", "href", &url2, nil, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.AttributeValue("img", "src", &image, nil, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text("h2", &name, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(".price", &price, chromedp.ByQuery, chromedp.FromNode(node)),
		)
	}

	return mangasPopulares
}

// getImagenManga obtiene la imagen de fondo de la lista de mangas de la pagina de inicio
func getImagenManga(imagenUrl string, mangaIdentificador string) string {
	cadenaBorrar := fmt.Sprintf(".book-thumbnail-%s::before{\n                     ", mangaIdentificador)

	cadenaSinEspacios := s.Trim(imagenUrl, "")

	cad1 := s.ReplaceAll(cadenaSinEspacios, "background-image: url('", " ")
	cad2 := s.ReplaceAll(cad1, "');\n                }", "")
	cad3 := s.ReplaceAll(cad2, cadenaBorrar, "")

	return s.TrimLeft(cad3, "")
}

// getImagenListaManga obtiene la imagen de fondo de la lista de mangas de la pagina de inicio
func getImagenListaManga(imagenUrl string, _ string) string {

	cadenaSinBackground := s.Replace(imagenUrl, "background-image: url('", " ", -1)

	cadenaSinClase := s.Replace(cadenaSinBackground, ".list-thumbnail", " ", -1)

	cadenaSinBefore := s.Replace(cadenaSinClase, "::before{\n", " ", -1)

	cadenaSinLLave := s.Replace(cadenaSinBefore, "}", " ", -1)
	cad1 := s.Replace(cadenaSinLLave, "');\n", " ", -1)

	//cad2 := fmt.Sprintf("-%s", mangaIdentificador)
	//cadSinDataSrc := s.Replace(cad1, cad2, " ", -1)

	return cad1

}

// Buscar manga por título
func BuscarMangas(order_item string, order_dir string, title string, _page string, filter_by string, Type string, demography string, status string, translation_status string, webcomic string, yonkoma string, amateur string, erotic string) []models.Library {
	var mangas []models.Library
	c := colly.NewCollector()
	encodedTitle := url.QueryEscape(title)

	url := fmt.Sprintf("https://zonatmo.com/library?order_item=%s&order_dir=%s&title=%s&_pg=%s&filter_by=%s&type=%s&demography=%s&status=%s&translation_status=%s&webcomic=%s&yonkoma=%s&amateur=%s&erotic=%s",
		order_item, order_dir, encodedTitle, _page, filter_by, Type, demography, status, translation_status, webcomic, yonkoma, amateur, erotic)

	c.OnHTML("#app > main > div:nth-child(2) > div.col-12.col-lg-8.col-xl-9 > div:nth-child(3)", func(element *colly.HTMLElement) {

		element.ForEach("div.element", func(i int, element *colly.HTMLElement) {

			dataItentificador := element.Attr("data-identifier")

			library := models.Library{
				Title:      element.ChildText("a > div > div > h4"),
				MangaUrl:   element.ChildAttr("a", "href"),
				Type:       element.ChildText("a > div > span.book-type"),
				Demography: element.ChildText("a > div > span.demography"),
				Score:      element.ChildText("a > div > span.score > span"),
				MangaImage: getImagenManga(element.ChildText("a > div > style"), dataItentificador),
			}

			mangas = append(mangas, library)

		})
	})

	c.Visit(url)
	return mangas

}

func GetLibraryMangas() []models.ListaManga {
	c := colly.NewCollector()
	var lista []models.ListaManga
	c.OnHTML("#app > main > div:nth-child(2) > div.col-12.col-lg-8.col-xl-9 > div:nth-child(3)", func(element *colly.HTMLElement) {
		element.ForEach("div.col-12.col-sm-12", func(i int, element *colly.HTMLElement) {

			urlLista := element.ChildAttr("a", "href")
			titleLista := element.ChildText("div.thumbnail > div.thumbnail-title > h4.text-truncate")
			descripcionLista := element.ChildText("div.thumbnail > div.thumbnail-description > p")
			cantidadDeSeguidoresLista := element.ChildText("div.thumbnail > div.thumbnail-container > span.followers_count")

			imagenLista := element.ChildText("div.thumbnail > style")
			dataSRC := utilities.GetDataSRCFromURL2(urlLista, titleLista)

			//fmt.Println(imagenLista)
			imagenFinal := getImagenListaManga(imagenLista, dataSRC)
			datos := models.ListaManga{
				Url:                       urlLista,
				Title:                     titleLista,
				Descripcion:               descripcionLista,
				CantidadDeSeguidoresLista: cantidadDeSeguidoresLista,
				ImagenLista:               imagenFinal,
				DataSRC:                   dataSRC,
			}
			lista = append(lista, datos)
		})
	})
	c.Visit("https://lectortmo.com/lists")
	return lista
}
