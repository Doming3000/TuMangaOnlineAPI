# TuMangaOnlineAPI
### **:robot: Autor original**
[_*Julio Olivares*_](https://github.com/julioolivares90)

## **:clipboard: Por hacer / ToDo**
- ~~Desplegar~~ *Fixed*
- ~~Arreglar la búsqueda~~ *Fixed*
- Arreglar la obtención de imágenes de un manga
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

<details>
<summary>Desplegar contenido JSON del ejemplo</summary>

```json
{
  "statusCode": 200,
  "data": [
    {
      "title": "Tensei Shitara Slime Datta Ken",
      "score": "9.20",
      "type": "MANGA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manga/12223/tenseishitaraslimedattaken",
      "mangaImagen": "https://otakuteca.com/images/books/cover/64cb883e56d87.webp"
    },
    {
      "title": "Chainsaw Man",
      "score": "8.81",
      "type": "MANGA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manga/336/chainsawman",
      "mangaImagen": "https://otakuteca.com/images/books/cover/648944edc1e26.webp"
    },
    {
      "title": "Goblin Slayer",
      "score": "8.15",
      "type": "MANGA",
      "demography": "Seinen",
      "mangaUrl": "https://zonatmo.com/library/manga/16354/goblinslayer",
      "mangaImagen": "https://otakuteca.com/images/books/cover/65ec500805f64.webp"
    },
    {
      "title": "Mago Infinito",
      "score": "9.06",
      "type": "MANHWA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manhwa/74407/elmagoinfinito",
      "mangaImagen": "https://otakuteca.com/images/books/cover/6655e1f55aba4.webp"
    },
    {
      "title": "Probablemente tambien va tras mi hermano",
      "score": "0.00",
      "type": "MANGA",
      "demography": "Shoujo",
      "mangaUrl": "https://zonatmo.com/library/manga/88751/osorakukanojowaorenoanikiwoneratteru",
      "mangaImagen": "https://otakuteca.com/images/books/cover/680fbb9ce3117.webp"
    },
    {
      "title": "Hataraku Maou-sama!",
      "score": "8.25",
      "type": "MANGA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manga/8657/hatarakumaousama",
      "mangaImagen": "https://otakuteca.com/images/books/cover/5eb35538d6b7d.webp"
    },
    {
      "title": "Mimi, una chica completamente normal",
      "score": "0.00",
      "type": "MANGA",
      "demography": "Seinen",
      "mangaUrl": "https://zonatmo.com/library/manga/88748/mimiunachicacompletamentenormal",
      "mangaImagen": "https://otakuteca.com/images/books/cover/680f5941c293f.webp"
    },
    {
      "title": "Dandadan",
      "score": "9.20",
      "type": "MANGA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manga/2005/dandadan",
      "mangaImagen": "https://otakuteca.com/images/books/cover/614cdc1808414.webp"
    },
    {
      "title": "Isekai Craft Gurashi: Jiyu Kimama na Seisan Shoku no Honobono Slow Life",
      "score": "10.00",
      "type": "MANGA",
      "demography": "Seinen",
      "mangaUrl": "https://zonatmo.com/library/manga/68554/isekai-craft-gurashi-jiyu-kimama-na-seisan-shoku-no-honobono-slow-life",
      "mangaImagen": "https://otakuteca.com/images/books/cover/6798e7957281b.webp"
    },
    {
      "title": "Regresión del bastardo del clan de la espada",
      "score": "6.00",
      "type": "MANHWA",
      "demography": "Seinen",
      "mangaUrl": "https://zonatmo.com/library/manhwa/86436/regressingasthereincarnatedbastardoftheswordclan",
      "mangaImagen": "https://otakuteca.com/images/books/cover/6717d5513d2bf.webp"
    },
    {
      "title": "Isekai Maou no Successor",
      "score": "7.00",
      "type": "MANGA",
      "demography": "Seinen",
      "mangaUrl": "https://zonatmo.com/library/manga/55813/isekai-maou-no-successor",
      "mangaImagen": "https://otakuteca.com/images/books/cover/66d139e5b84b6.webp"
    },
    {
      "title": "Alguien ha poseído mi cuerpo",
      "score": "8.00",
      "type": "MANHWA",
      "demography": "Shoujo",
      "mangaUrl": "https://zonatmo.com/library/manhwa/77595/alguien-ha-poseido-mi-cuerpo",
      "mangaImagen": "https://otakuteca.com/images/books/cover/647772e7973d0.webp"
    },
    {
      "title": "Shiunji-ke no Kodomotachi",
      "score": "7.88",
      "type": "MANGA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manga/67271/shiunjikenokodomotachi",
      "mangaImagen": "https://otakuteca.com/images/books/cover/67d0d1fc22229.webp"
    },
    {
      "title": "La regresión 100 del jugador de nivel máximo",
      "score": "10.00",
      "type": "MANHWA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manhwa/80382/the100thregressionofthemaxlevelplayer",
      "mangaImagen": "https://otakuteca.com/images/books/cover/6795926d34b45.webp"
    },
    {
      "title": "Mí Nueva Novia no es Humana?",
      "score": "9.50",
      "type": "MANGA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manga/87017/minuevanovianoeshumana",
      "mangaImagen": "https://otakuteca.com/images/books/cover/67e698739676d.webp"
    },
    {
      "title": "El amor de Nanase es anormal",
      "score": "10.00",
      "type": "MANGA",
      "demography": "Seinen",
      "mangaUrl": "https://zonatmo.com/library/manga/86293/nanasesannokoigaijou",
      "mangaImagen": "https://otakuteca.com/images/books/cover/670718c2f228c.webp"
    },
    {
      "title": "Megane, Tokidoki, Yankee-kun",
      "score": "9.25",
      "type": "MANGA",
      "demography": "Shoujo",
      "mangaUrl": "https://zonatmo.com/library/manga/56564/megane-tokidoki-yankee-kun",
      "mangaImagen": "https://otakuteca.com/images/books/cover/67b7540e7fe69.webp"
    },
    {
      "title": "El Apocalipsis Ya Está Aquí",
      "score": "0.00",
      "type": "MANHWA",
      "demography": "Seinen",
      "mangaUrl": "https://zonatmo.com/library/manhwa/88604/theapocalypseishere",
      "mangaImagen": "https://otakuteca.com/images/books/cover/68017a892de3c.webp"
    },
    {
      "title": "Kage no Jitsuryokusha ni Naritakute! - crónica de las 7 sombras",
      "score": "10.00",
      "type": "MANGA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manga/84143/kagenojitsuryokushaninaritakutemasterofgardenshichikageretsuden",
      "mangaImagen": "https://otakuteca.com/images/books/cover/672f6059e045b.webp"
    },
    {
      "title": "Me convertí en el genio bastardo de un noble clan oscuro",
      "score": "0.00",
      "type": "MANHWA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manhwa/88416/ibecamethebastardgeniusofanobledarkclan",
      "mangaImagen": "https://otakuteca.com/images/books/cover/67ec71a55f2bd.webp"
    },
    {
      "title": "Multiverse no Watashi, Koishite Ii desu ka?",
      "score": "9.00",
      "type": "MANGA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manga/85093/multiversenowatashikoishiteiidesuka",
      "mangaImagen": "https://otakuteca.com/images/books/cover/67a5f99c144e3.webp"
    },
    {
      "title": "Gorin no Megami-sama",
      "score": "9.50",
      "type": "MANGA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manga/64060/gorin-no-megami-sama",
      "mangaImagen": "https://otakuteca.com/images/books/cover/6292e89658167.webp"
    },
    {
      "title": "A solas en la habitación de una hermosa chica en pijama",
      "score": "0.00",
      "type": "MANGA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manga/87550/muboubikawaiipajamasugatanobishoujotoheyadefutarikiri",
      "mangaImagen": "https://otakuteca.com/images/books/cover/6788aae16d55d.webp"
    },
    {
      "title": "Me convertí en la sirvienta del Tirano.",
      "score": "9.00",
      "type": "MANHWA",
      "demography": "Shoujo",
      "mangaUrl": "https://zonatmo.com/library/manhwa/84518/meconvertienlasirvientadeltirano",
      "mangaImagen": "https://otakuteca.com/images/books/cover/672993e454e41.webp"
    },
    {
      "title": "El chico a mi lado me está preocupando de nuevo hoy",
      "score": "0.00",
      "type": "MANGA",
      "demography": "Shoujo",
      "mangaUrl": "https://zonatmo.com/library/manga/88263/otonarikunwakyoumowatashiwokomaraseru",
      "mangaImagen": "https://otakuteca.com/images/books/cover/67db485b6da5c.webp"
    },
    {
      "title": "Majo Taisen - The War of Greedy Witches",
      "score": "8.14",
      "type": "MANGA",
      "demography": "Seinen",
      "mangaUrl": "https://zonatmo.com/library/manga/54731/majo-taisen-32-nin-no-isai-no-majo-wa-koroshiau",
      "mangaImagen": "https://otakuteca.com/images/books/cover/67e99bf3505cf.webp"
    },
    {
      "title": "Solo quiero derrotarte",
      "score": "9.00",
      "type": "MANHWA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manhwa/53552/solo-quiero-vencerte",
      "mangaImagen": "https://otakuteca.com/images/books/cover/5f674c13d1039.webp"
    },
    {
      "title": "Kagurabachi",
      "score": "8.63",
      "type": "MANGA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manga/2054/kagurabachi",
      "mangaImagen": "https://otakuteca.com/images/books/cover/65076987c0141.webp"
    },
    {
      "title": "Destinado a ser Amado por Villanas",
      "score": "8.50",
      "type": "MANHWA",
      "demography": "Seinen",
      "mangaUrl": "https://zonatmo.com/library/manhwa/88040/destinadoaseramadoporvillanas",
      "mangaImagen": "https://otakuteca.com/images/books/cover/67be79886bee1.webp"
    },
    {
      "title": "The Villain Of Destiny",
      "score": "10.00",
      "type": "MANHUA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manhua/63572/el-villano-del-destino",
      "mangaImagen": "https://otakuteca.com/images/books/cover/6159b0b038dd6.webp"
    }
  ]
}
```
</details>

## 🔥 Obtener todos los mangas populares seinen:
https://tumangaonlineapi-production.up.railway.app/api/v1/manga/populares-seinen

<details>
<summary>Desplegar contenido JSON del ejemplo</summary>

```json
{
    "statusCode": 200,
    "data": [
        {
            "title": "Dueña de la Pensión",
            "score": "0.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/53758/duena-de-la-pension",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f7186dfd183c.jpg"
        },
        {
            "title": "Jimiko wa Igai ni Ero Katta",
            "score": "0.00",
            "type": "MANGA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manga/49040/jimiko-wa-igai-ni-erokatta",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5e87a536c19a4.jpg"
        },
        {
            "title": "Young Boss",
            "score": "10.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/49943/young-boss",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5eb71213bcb93.jpg"
        },
        {
            "title": "SISTERS AND FATHER",
            "score": "0.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/47560/sisters-and-father",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5e0ba6dcc40ae.jpg"
        },
        {
            "title": "Luvslave",
            "score": "10.00",
            "type": "MANGA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manga/44087/luvslave",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5cd105350800b.jpg"
        },
        {
            "title": "A KNOWING SISTER",
            "score": "0.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/53156/a-knowing-sister",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f53afdf115ed.jpg"
        },
        {
            "title": "The Girl Hiding in the Wall",
            "score": "0.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/49956/the-girl-hiding-in-the-wall",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5ec0c9a5e9b9d.jpg"
        },
        {
            "title": "Cowgirl´s Riding-Position Makes Me Cum",
            "score": "10.00",
            "type": "MANGA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manga/46302/cowgirls-riding-position-makes-me-cum",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5d94b98f26c01.jpg"
        },
        {
            "title": "Bullied Boy's Tongue Revenge",
            "score": "10.00",
            "type": "MANGA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manga/46118/bullied-boys-tongue-revenge",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5d84e2f18a859.jpg"
        },
        {
            "title": "Cram School Scandal",
            "score": "8.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/50813/cram-school-scandal",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f5e66e7894f2.jpg"
        },
        {
            "title": "Uq Holder",
            "score": "7.67",
            "type": "MANGA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manga/8777/uq-holder",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5e283e1726909.jpg"
        },
        {
            "title": "A Killer Woman",
            "score": "0.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/49207/a-killer-woman",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5e90fd0651625.jpg"
        },
        {
            "title": "Depredadores de la Estética",
            "score": "0.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/51526/depredadores-de-la-estetica",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f03fe5f3e512.jpg"
        },
        {
            "title": "Mala Sangre",
            "score": "0.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/50385/bad-blood",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5ecc89bc91261.jpg"
        },
        {
            "title": "Grabbed My Boobs For 16 Hours",
            "score": "10.00",
            "type": "MANGA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manga/46218/grabbed-my-boobs-for-16-hours",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5d8b7e4dd3c8b.jpg"
        },
        {
            "title": "Ejaculation",
            "score": "8.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/48710/ejaculation",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5e6eeca783872.jpg"
        },
        {
            "title": "El vastago",
            "score": "0.00",
            "type": "NOVELA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/novel/54074/el-vastago",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f80f6d734cbc.jpg"
        },
        {
            "title": "Tengo una Mansión en el mundo Post-apocalíptico",
            "score": "8.50",
            "type": "NOVELA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/novel/45123/tengo-una-mansion-en-el-mundo-post-apocaliptico",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5d2e05cfaec2c.jpg"
        },
        {
            "title": "La novia de mi amigo",
            "score": "8.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/49044/la-novia-de-mi-amigo",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5e87d57437924.jpg"
        },
        {
            "title": "SISTERS WIVE",
            "score": "9.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/41955/sisters-wive",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5c13d0d017c81.jpg"
        },
        {
            "title": "Necesito que me ayudes",
            "score": "9.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/48708/necesito-que-me-ayudes",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5e6eea57e8bb3.jpg"
        }
    ]
}
```
</details>

## 📋 Obtener la informacion de un manga:
https://tumangaonlineapi-production.up.railway.app/api/v1/manga/info?mangaUrl=https://lectortmo.com/library/manga/23741/dr-stone

<details>
<summary>Desplegar contenido JSON del ejemplo</summary>

```json
{
  "statusCode": 200,
  "data": {
    "title": "Dr. Stone ( 2017 - 2022 )",
    "image": "https://otakuteca.com/images/books/cover/5d924d4309b18.webp",
    "tipo": "MANGA",
    "score": "8.96",
    "demografia": "Shounen",
    "descripcion": "Senkuu y Taiju son dos amigos inseparables durante sus años de preparatoria: el primero, un prodigio de la química, y el segundo, un grandullón musculoso de gran corazón, aunque no precisamente brillante. Cinco años después, cuando Taiju decide declararle su amor a Yuzuriha, un suceso inimaginable sacude la Tierra: un cataclismo transforma a la humanidad en piedra. Ahora, en un mundo postapocalíptico, ambos deberán combinar la astucia científica de Senkuu y la determinación inquebrantable de Taiju para luchar por la supervivencia y, quizá, devolver la vida a un planeta petrificado.",
    "estado": "Finalizado",
    "generos": [
      "Acción",
      "Aventura",
      "Apocalíptico",
      "Comedia",
      "Supervivencia",
      "Misterio",
      "Ciencia Ficción"
    ],
    "capitulo": [
      {
        "Title": "Capítulo 236.00 Spin Off: Capítulo 3, Ciencia del Futuro",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 235.00 Spin Off Capítulo 2",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 234.00 Spin Off Capítulo 1",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 233.00 Epílogo - Extra",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 232.10 Último Capítulo A COLOR",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 232.00 Último Capítulo",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 231.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 230.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 229.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 228.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 227.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 226.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 225.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 224.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 223.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 222.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 221.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 220.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 219.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 218.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 217.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 216.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 215.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 214.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 213.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 212.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 211.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 210.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 209.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 208.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 207.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 206.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 205.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 204.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 203.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 202.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 201.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 200.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 199.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 198.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 197.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 196.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 195.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 194.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 193.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 192.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 191.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 190.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 189.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 188.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 187.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 186.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 185.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 184.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 183.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 182.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 181.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 180.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 179.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 178.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 177.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 176.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 175.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 174.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 173.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 172.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 171.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 170.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 169.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 168.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 167.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 166.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 165.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 164.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 163.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 162.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 161.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 160.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 159.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 158.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 157.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 156.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 155.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 154.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 153.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 152.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 151.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 150.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 149.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 148.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 147.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 146.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 145.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 144.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 143.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 142.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 141.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 140.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 139.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 138.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 137.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 136.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 135.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 134.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 133.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 132.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 131.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 130.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 129.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 128.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 127.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 126.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 125.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 124.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 123.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 122.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 121.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 120.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 119.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 118.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 117.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 116.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 115.00 Un Segundo y un Grano",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 114.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 113.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 112.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 111.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 110.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 109.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 108.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 107.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 106.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 105.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 104.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 103.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 102.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 101.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 100.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 99.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 98.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 97.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 96.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 95.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 94.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 93.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 92.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 91.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 90.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 89.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 88.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 87.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 86.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 85.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 84.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 83.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 82.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 81.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 80.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 79.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 78.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 77.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 76.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 75.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 74.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 73.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 72.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 71.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 70.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 69.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 68.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 67.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 66.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 65.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 64.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 63.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 62.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 61.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 60.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 59.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 58.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 57.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 56.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 55.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 54.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 53.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 52.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 51.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 50.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 49.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 48.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 47.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 46.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 45.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 44.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 43.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 42.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 41.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 40.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 39.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 38.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 37.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 36.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 35.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 34.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 33.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 32.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 31.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 30.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 29.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 28.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 27.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 26.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 25.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 24.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 23.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 22.50",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 22.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 21.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 20.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 19.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 18.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 17.00",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 16.00 Kohaku",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 15.00 Los dos países del mundo de piedra",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 14.00 En lo que crees",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 13.00 El comienzo del mundo de piedra",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 12.00 Epílogo del Prólogo",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 11.00 El arma de la ciencia",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 10.00 La banda de la ciencia",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 9.50 Extra - ¡¡La cerámica: Intentémos hacerlo!!",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 9.00 Senku vs Tsukasa",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 8.00 Encender la señal",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 7.00 La aventura de la pólvora",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 6.00 Taiju vs Tsukasa",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 5.00 Yuzuriha",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 4.00 La almeja blanca pura",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 3.00 Rey del mundo de piedra",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 2.00 Fantasia vs Ciencia",
        "UrlLeer": ""
      },
      {
        "Title": "Capítulo 1.00 Mundo de piedra",
        "UrlLeer": ""
      }
    ]
  }
}
```
</details>

### obtener las imagenes de un capitulo de un manga
https://tumangaonlineapi-production.up.railway.app/api/v1/get-manga?urlPage=https://lectortmo.com/view_uploads/569910
```json
{
    "statusCode": 200,
    "data": [
        "https://img1.tucomiconline.com/uploads/5f2701e221574/f64623e3.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/1e55867a.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/863bd767.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/2fd25756.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/7a465f81.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/ded9be6a.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/54739f43.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/a2b12cfa.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/1db24da4.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/d073729a.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/71042aa9.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/c22a10c7.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/8286bd16.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/930dfe36.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/cbebfd4a.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/611e4bb0.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/b6c20dc9.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/42671792.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/e5b57b1e.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/7558c109.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/a110143b.jpg"
    ]
}
```
### realizar una busqueda de un manga
### listado de parametros 
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

## Enlace 
https://tumangaonlineapi-production.up.railway.app/api/v1/manga/library?title=naruto

<details>
<summary>Desplegar contenido JSON del ejemplo</summary>

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
      "mangaImagen": "https://otakuteca.com/images/books/cover/61eabe6e3934b.webp"
    },
    {
      "title": "Naruto Shippuden parejas Doujines Yaoi",
      "score": "7.00",
      "type": "DOUJINSHI",
      "demography": "Josei",
      "mangaUrl": "https://zonatmo.com/library/doujinshi/57108/naruto-shippuden-parejas-doujines-yaoi",
      "mangaImagen": "https://otakuteca.com/images/books/cover/601ca58f9692d.webp"
    },
    {
      "title": "Naruto Full Color",
      "score": "9.25",
      "type": "MANGA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manga/475/naruto-digital-colored-comics",
      "mangaImagen": "https://otakuteca.com/images/books/cover/617ebcafa0486.webp"
    },
    {
      "title": "Joke Box 7 Naruto Fanbook-Kakashi x Iruka",
      "score": "7.00",
      "type": "DOUJINSHI",
      "demography": "Shoujo",
      "mangaUrl": "https://zonatmo.com/library/doujinshi/40948/joke-box-7-naruto-fanbook-kakashi-x-iruka",
      "mangaImagen": "https://otakuteca.com/images/books/cover/5ba4189b74e65.webp"
    },
    {
      "title": "Tengo la mano fría",
      "score": "0.00",
      "type": "DOUJINSHI",
      "demography": "Shoujo",
      "mangaUrl": "https://zonatmo.com/library/doujinshi/59723/naruto-dj-tenohira-ga-atsui",
      "mangaImagen": "https://otakuteca.com/images/books/cover/6098261956cee.webp"
    },
    {
      "title": "Naruto: Sasuke Retsuden - The Manga",
      "score": "7.00",
      "type": "MANGA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manga/73141/narutosasukeretsuden",
      "mangaImagen": "https://otakuteca.com/images/books/cover/635833b8423ea.webp"
    },
    {
      "title": "NARUTO Gaiden: Uzu no Naka no Tsumujikaze",
      "score": "8.50",
      "type": "ONE SHOT",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/one_shot/78459/naruto-gaiden-uzu-no-naka-no-tsumujikaze",
      "mangaImagen": "https://otakuteca.com/images/books/cover/64b4339e9a7c7.webp"
    },
    {
      "title": "Naruto Gaiden: Nanadaime Hokage to Akairo no Hanatsuzuki",
      "score": "8.00",
      "type": "MANGA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manga/11344/naruto-gaiden-nanadaime-hokage-to-akairo-no-hanatsuzuki",
      "mangaImagen": "https://otakuteca.com/images/books/cover/5edb2d0bc3bb2.webp"
    },
    {
      "title": "Naruto - Watashi no Kareshi wa Keikokubijin (Doujinshi)",
      "score": "0.00",
      "type": "DOUJINSHI",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/doujinshi/46640/naruto-watashi-no-kareshi-wa-keikokubijin-doujinshi",
      "mangaImagen": "https://otakuteca.com/images/books/cover/5db834ffccd43.webp"
    },
    {
      "title": "Renge to Naruto!",
      "score": "10.00",
      "type": "MANGA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manga/80575/renge-to-naruto",
      "mangaImagen": "https://otakuteca.com/images/books/cover/6550532b16a6c.webp"
    },
    {
      "title": "Naruto Especial: Boruto Road to B",
      "score": "8.35",
      "type": "ONE SHOT",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/one_shot/12697/Naruto-Especial-Boruto-Road-to",
      "mangaImagen": "https://otakuteca.com/images/books/cover/12697_TMOmanga043300.webp"
    },
    {
      "title": "Naruto- festival de verano",
      "score": "0.00",
      "type": "DOUJINSHI",
      "demography": "Shoujo",
      "mangaUrl": "https://zonatmo.com/library/doujinshi/52147/naruto-festival-de-verano",
      "mangaImagen": "https://otakuteca.com/images/books/cover/5f23b7fbea104.webp"
    },
    {
      "title": "La Supremacia De Naruto",
      "score": "7.50",
      "type": "NOVELA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/novel/57666/la-supremacia-de-naruto",
      "mangaImagen": "https://otakuteca.com/images/books/cover/6041fae923aaa.webp"
    },
    {
      "title": "Naruto (Piloto)",
      "score": "8.00",
      "type": "ONE SHOT",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/one_shot/31871/naruto-piloto",
      "mangaImagen": "https://otakuteca.com/images/books/cover/5f46aa3cb4e9c.webp"
    },
    {
      "title": "Naruto: The Path Lit by the Full Moon",
      "score": "8.00",
      "type": "ONE SHOT",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/one_shot/15552/naruto-gaiden-michita-tsuki-ga-terasu-michi",
      "mangaImagen": "https://otakuteca.com/images/books/cover/15552_TMOManga5719e0cca25b9.webp"
    },
    {
      "title": "Naruto: The Path Lit by the Full Moon",
      "score": "9.00",
      "type": "ONE SHOT",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/one_shot/12844/Naruto-ga-Hokage-ni-Natta-Hi",
      "mangaImagen": "https://otakuteca.com/images/books/cover/12844_TMOmanga024024.webp"
    },
    {
      "title": "Naruto Feliz Cumpleaños",
      "score": "0.00",
      "type": "ONE SHOT",
      "demography": "Shoujo",
      "mangaUrl": "https://zonatmo.com/library/one_shot/58997/sasunaru-naruto-feliz-cumpleanos",
      "mangaImagen": "https://otakuteca.com/images/books/cover/607502731bf66.webp"
    },
    {
      "title": "Naruto-sazanka",
      "score": "6.00",
      "type": "DOUJINSHI",
      "demography": "Shoujo",
      "mangaUrl": "https://zonatmo.com/library/doujinshi/52754/naruto-sazanka",
      "mangaImagen": "https://otakuteca.com/images/books/cover/5f40eaa6178fb.webp"
    },
    {
      "title": "Sistema Naruto en el mundo de One Piece",
      "score": "7.00",
      "type": "NOVELA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/novel/57622/sistema-naruto-en-el-mundo-de-one-piece",
      "mangaImagen": "https://otakuteca.com/images/books/cover/6033c3ab501df.webp"
    },
    {
      "title": "Konoha's Story—The Steam Ninja Scrolls",
      "score": "0.00",
      "type": "MANGA",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/manga/73582/naruto-konohas-story-the-steam-ninja-scrolls",
      "mangaImagen": "https://otakuteca.com/images/books/cover/6394e019eacaa.webp"
    },
    {
      "title": "In a Different World with the Naruto System",
      "score": "0.00",
      "type": "NOVELA",
      "demography": "Seinen",
      "mangaUrl": "https://zonatmo.com/library/novel/41779/in-a-different-world-with-the-naruto-system",
      "mangaImagen": "https://otakuteca.com/images/books/cover/5c027a1806966.webp"
    },
    {
      "title": "NARUTO GAIDEN NUEVO CICLO",
      "score": "5.59",
      "type": "DOUJINSHI",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/doujinshi/19625/NARUTO-GAIDEN-NUEVO-CICLO",
      "mangaImagen": "https://otakuteca.com/images/books/cover/19625_TMOManga57c76213ecc34.webp"
    },
    {
      "title": "Naruto ND",
      "score": "6.75",
      "type": "DOUJINSHI",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/doujinshi/16018/Naruto-ND",
      "mangaImagen": "https://otakuteca.com/images/books/cover/16018_TMOManga5737cbea47cd2.webp"
    },
    {
      "title": "Ruki (Naruto Universe)",
      "score": "0.00",
      "type": "DOUJINSHI",
      "demography": "Shounen",
      "mangaUrl": "https://zonatmo.com/library/doujinshi/80650/ruki-naruto-universe",
      "mangaImagen": "https://otakuteca.com/images/books/cover/6555afe6f12d2.webp"
    }
  ]
}
```
</details>
