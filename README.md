<p align="center">
    <img src="https://raw.githubusercontent.com/StarpTech/fay/main/assets/logo.png" alt="fay logo"/>
</p>
<h3 align="center">Fay</h3>
<p align="center">HTTP Service that expose the PDF rendering capabilities of <a href="https://github.com/microsoft/playwright">Playwright</a>.</p>
<p align="center"><a href="https://hub.docker.com/repository/docker/starptech/fay">Dockerhub</a> &#183; <a href="/loadtesting/README.md">Load tests</a> &#183; <a href="/.github/CONTRIBUTING.md">Contributing</a></p>
<p align="center"><a href="https://travis-ci.org/thecodingmachine/gotenberg" rel="nofollow"><img src="https://github.com/StarpTech/fay/workflows/tests/badge.svg" style="max-width:100%;"></a></p>

## Features

- Fast and reliable due to [Playwright](https://github.com/microsoft/playwright).
- Use Chromium 89.0.4344.0.
- Render any URL or static HTML to PDF.
- Sandbox mode with network and javascript disabled.
- Healthcheck endpoint.
- Ready to use docker image.
- Swagger endpoint.

## Installation

```
docker run --rm -p 3000:3000 starptech/fay
```

_The image is relatively big due to the playwright base image. This might be improved in the future. Versioning will be added soon as well._

## Endpoints

- `/convert` - **Converts a website to PDF document.**
  - `filename` (form,query, **default:** `"result.pdf"`): Filename of the resulting pdf.
  - `url` (form,query, **default:** `""`): The url of the website to convert.
  - `html` (file, **default:** `""`): Convert the HTML to PDF instead `url`.
  - `locale` (form,query, **default:** `en-US`): Browser locale.
  - `format` (form,query, **default:** `A4`): Page format.
  - `offline` (form,query, **default:** `false`): Disable network connectivity.
  - `media` (form,query, **default:** `print`): Page media emulation.
  - `javaScriptEnabled` (form,query, **default:** `false`): Enable javascript on the website.
  - `marginTop,marginRight,marginBottom,marginLeft` (form,query, **default:** `0`): Set page margin.
  - `headerTemplate` (file, **default:** `<span></span>`): Header template.
  - `footerTemplate` (file, **default:** `<span></span>`): Footer template.
- `/ping` - **Check if the server is ready to accept requests.**
- `/swagger/index.html` - **Swagger introspection**

For detail description of the pdf options check the [playwright](https://playwright.dev/docs/api/class-page?_highlight=pdf#pagepdfoptions) documentation.

## Scalability

Playwright is capable to maintain a pool of browser instances. This allows to handle many requests at the same time.

In a simple test we could serve up to 20 virtual users over a time window of 5min with an average request duration of ~1s and ~13 req/s. We were even capable to serve up to 100 virtual users over a time window of 5min without any error with an average request duration of ~4.25s and ~11 req/s.
For more informations check the [load-test](./loadtesting/README.md).

## Best practice

In order to produce reproducible results try to avoid downloading external resources you can't control and executing javascript. You can inline images and styles in the document. Use the options `offline=true` and `javaScriptEnabled=false` to enforce that. Fonts can be embedded in the docker image to make them accesible to the chromium browser.

## Credits

- https://github.com/mxschmitt/playwright-go
- https://github.com/thecodingmachine/gotenberg
- <a target="_blank" href="https://icons8.com/icons/set/elf--v2">Elf icon</a> icon by <a target="_blank" href="https://icons8.com">Icons8</a>
