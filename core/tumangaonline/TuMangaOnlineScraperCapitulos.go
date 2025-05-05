package tumangaonline

import (
	"net/http"
	"strings"

	"github.com/gocolly/colly"
)

// GetCookiesFromTMO obtiene las cookies de una sesion en la pagina de tmo
func GetCookiesFromTMO() (map[string]string, error) {
	cookies := make(map[string]string)
	client := http.Client{}

	request, err := http.NewRequest("GET", "https://lectortmo.com/", nil)
	if err != nil {
		return cookies, err
	}

	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.80 Safari/537.36 Edg/86.0.622.43")
	request.Header.Set("sec-fetch-dest", "document")
	request.Header.Set("referer", "https://lectortmo.com/")
	request.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	request.Header.Set("method", "GET")
	request.Header.Set("authority", "lectortmo.com")
	request.Header.Set("accept-encoding", "gzip, deflate, br")
	request.Header.Set("accept-language", "es-419,es;q=0.9,es-ES;q=0.8,en;q=0.7,en-GB;q=0.6,en-US;q=0.5")

	response, err := client.Do(request)
	if err != nil {
		return cookies, err
	}

	cookiesResp := response.Cookies()
	for _, cookie := range cookiesResp {
		cookies[cookie.Name] = cookie.Value
	}

	return cookies, nil
}

// Obtener las imágenes de un capítulo
func GetImageChapter(urlRefer string, url string) ([]string, error) {
	var imagenes []string

	// Forzar directamente a /cascade si venía con /paginated
	if strings.Contains(url, "/paginated") {
		url = strings.Replace(url, "/paginated", "/cascade", 1)
	}

	c := colly.NewCollector()

	c.OnHTML("body", func(element *colly.HTMLElement) {
		images := getImagesFromHTMLParsed(element)
		imagenes = append(imagenes, images...)
	})

	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("referer", urlRefer)
		request.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)...")
	})

	err := c.Visit(url)
	return imagenes, err
}

// Obtener los enlaces a las imágenes de los capítulos
func getImagesFromHTMLParsed(document *colly.HTMLElement) []string {
	var images []string

	document.ForEach("div.viewer-image-container", func(i int, element *colly.HTMLElement) {
		image := element.ChildAttr("img", "data-src")
		if image != "" {
			images = append(images, image)
		}
	})

	document.ForEach("div.img-container", func(i int, element *colly.HTMLElement) {
		image := element.ChildAttr("img", "data-src")
		if image != "" {
			images = append(images, image)
		}
	})

	return images
}

// Pendiente: No se puede acceder a las urls por que la página las bloquea (403)
