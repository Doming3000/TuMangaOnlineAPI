# TuMangaOnlineAPI
### **:robot: Autor original**
[_*Julio Olivares*_](https://github.com/julioolivares90)

## **:clipboard: Por hacer / ToDo**
- ~~Desplegar~~ *Fixed*
- ~~Arreglar la búsqueda por nombre~~ *Fixed*
- Revisar y testear la búsqueda por otros parámetros
- Arreglar la obtención de imágenes de un manga
- Arreglar la búsqueda por generos
- Arreglar los enlaces a los capitulos al obtener la información de un manga

## **:package: Principales herramientas usadas**

- [x] fiber
- [x] go colly

## :rocket: TuMangaOnlineAPI API enlace:
https://tumangaonlineapi-production.up.railway.app

### Iniciar el servidor usando go:

```
go run App.go
```

### Request & Response Ejemplos de uso:

## 🗞️ Obtener todos los mangas populares:
https://tumangaonlineapi-production.up.railway.app/api/v1/manga/populares

Ejemplo de respuesta (limitado a 2 resultados para facilitar la lectura):
```json
{
  "statusCode": 200,
  "data": [
    {
      "title": "Kage no Jitsuryokusha ni Naritakute!",
      "score": "8.42",
      "type": "MANGA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manga/42185/kagenojitsuryokushaninaritakute",
      "mangaImage": "https://otakuteca.com/images/books/cover/64cacac18e007.webp"
    },
    {
      "title": "La Vida Después de la Muerte",
      "score": "8.79",
      "type": "MANHWA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manhwa/47049/lavidadespuesdelamuerte",
      "mangaImage": "https://otakuteca.com/images/books/cover/5ddde8a92558c.webp"
    }
  ]
}
```

## 📋 Obtener la informacion de un manga:
https://tumangaonlineapi-production.up.railway.app/api/v1/manga/info?mangaUrl=https://lectortmo.com/library/manga/23741/dr-stone

Ejemplo de respuesta (limitado a 3 capítulos para facilitar la lectura):
```json
{
  "statusCode": 200,
  "data": {
    "title": "Dr. Stone ( 2017 - 2022 )",
    "image": "https://otakuteca.com/images/books/cover/5d924d4309b18.webp",
    "type": "MANGA",
    "score": "8.96",
    "demography": "Shounen",
    "synopsis": "Senkuu y Taiju son dos amigos inseparables durante sus años de preparatoria: el primero, un prodigio de la química, y el segundo, un grandullón musculoso de gran corazón, aunque no precisamente brillante. Cinco años después, cuando Taiju decide declararle su amor a Yuzuriha, un suceso inimaginable sacude la Tierra: un cataclismo transforma a la humanidad en piedra. Ahora, en un mundo postapocalíptico, ambos deberán combinar la astucia científica de Senkuu y la determinación inquebrantable de Taiju para luchar por la supervivencia y, quizá, devolver la vida a un planeta petrificado.",
    "status": "Finalizado",
    "genres": [
      "Acción",
      "Aventura",
      "Apocalíptico",
      "Comedia",
      "Supervivencia",
      "Misterio",
      "Ciencia Ficción"
    ],
    "chapters": [
      {
        "Title": "Capítulo 3.00 Rey del mundo de piedra"
      },
      {
        "Title": "Capítulo 2.00 Fantasia vs Ciencia"
      },
      {
        "Title": "Capítulo 1.00 Mundo de piedra"
      }
    ]
  }
}
```
## 🔎 Realizar una busqueda de un manga
https://tumangaonlineapi-production.up.railway.app/api/v1/manga/library?title=naruto

### Listado de parametros 
- title
- _page
- orderItem
- orderDir
- filter_by
- Type
- demography
- status
- translation_status
- webcomic
- yonkoma
- amateur
- erotic

Ejemplo de respuesta (limitado a 3 resultados para facilitar la lectura):
```json
{
  "statusCode": 200,
  "data": [
    {
      "title": "Boruto: Naruto Next Generations",
      "score": "5.00",
      "type": "MANGA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manga/2991/boruto-naruto-next-generations",
      "mangaImage": "https://otakuteca.com/images/books/cover/61eabe6e3934b.webp"
    },
    {
      "title": "Naruto Shippuden parejas Doujines Yaoi",
      "score": "7.00",
      "type": "DOUJINSHI",
      "demography": "Josei",
      "mangaUrl": "https://zonatmo.com/library/doujinshi/57108/naruto-shippuden-parejas-doujines-yaoi",
      "mangaImage": "https://otakuteca.com/images/books/cover/601ca58f9692d.webp"
    },
    {
      "title": "Naruto Full Color",
      "score": "9.25",
      "type": "MANGA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manga/475/naruto-digital-colored-comics",
      "mangaImage": "https://otakuteca.com/images/books/cover/617ebcafa0486.webp"
    }
  ]
}
```
## 📷 Obtener las imágenes de un capitulo (WIP)
https://tumangaonlineapi-production.up.railway.app/api/v1/get-manga?urlPage=https://zonatmo.com/viewer/0e59105a34f7d280b19054cdf2c90d9c/paginated
