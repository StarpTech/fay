# 🧚‍♂️ Fay - HTML to PDF rendering service

Fay is a HTTP Service which expose the PDF rendering capabilities of [Playwright](https://github.com/microsoft/playwright) to automate [Chromium](https://www.chromium.org/Home).

## Installation

```
docker run --rm -p 3000:3000 starptech/fay
```
_The image is relatively big due to the playwright base image. This might be improved in the future. Versioning will be added soon as well._

## Endpoints

- `/convert` - **Converts a website to PDF document.**
  - `url` (form,query): The url of the website to convert.
  - `locale` (form,query, **default:** `en-US`): Browser locale.
  - `format` (form,query, **default:** `A4`): Page format.
  - `offline` (form,query, **default:** `false`): Enable loading external resources.
  - `media` (form,query, **default:** `print`): Page media emulation.
  - `javaScriptEnabled` (form,query, **default:** `false`): Enable javascript on the website.
  - `marginTop,marginRight,marginBottom,marginLeft` (form,query, **default:** `0`): Set page margin.
  - `html` (file, **default:** ``): Convert the HTML to PDF instead `url`.
  - `headerTemplate` (file, **default:** `<span></span>`): Header template.
  - `footerTemplate` (file, **default:** `<span></span>`): Footer template.
- `/ping` - **Check if the server is ready to accept requests.**

For detail description of the pdf options check the [playwright](https://playwright.dev/docs/api/class-page?_highlight=pdf#pagepdfoptions) documentation.

## Scalability

Playwright is capable to maintain a pool of browser instances. This allows to handle many requests at the same time.

In a simple test we could serve 20 virtual users (~10.4 requests per second) with an average request duration under 1s.
For more informations check the [load-test](./loadtesting/README.md).

## Best practice

In order to produce reproducible results try to avoid downloading external resources you can't control and executing javascript. You can inline images and styles in the document. Use the options `offline=true` and `javaScriptEnabled=false` to enforce that. Fonts can be embedded in the docker image to make them accesible to the chromium browser.

## Development

1. Install `npm install -g serve`.
2. Serve example template `serve -l 3001 ./example`.
3. Run server `go run cmd/fay/main.go`.

Convert a document.

```
curl --request POST \
    --url http://localhost:3000/convert \
    --header 'Content-Type: multipart/form-data' \
    --form html=http://localhost:3001/page.html \
    --form locale=de-DE \
    --form marginTop=0.5in \
    --form headerTemplate=@example/header.html \
    -o html.pdf
```

## Credits

https://github.com/thecodingmachine/gotenberg