<p align="center">
    <img src="https://raw.githubusercontent.com/StarpTech/fay/main/assets/logo.png" alt="fay logo"/>
</p>
<h3 align="center">Fay</h3>
<p align="center">Stateless, Fast and Reliable PDF rendering service.</p>
<p align="center"><a href="https://hub.docker.com/repository/docker/starptech/fay">Dockerhub</a> &#183; <a href="/loadtesting/README.md">Load tests</a> &#183; <a href="/.github/CONTRIBUTING.md">Contributing</a></p>
<p align="center"><a href="https://github.com/StarpTech/fay/actions?query=workflow%3Atests" rel="nofollow"><img src="https://github.com/StarpTech/fay/workflows/tests/badge.svg" style="max-width:100%;"></a></p>

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
  - `javascript` (form,query, **default:** `true`): Enable javascript on the website.
  - `marginTop,marginRight,marginBottom,marginLeft` (form,query, **default:** `0`): Set page margin.
  - `headerTemplate` (file, **default:** `<span></span>`): Header template.
  - `footerTemplate` (file, **default:** `<span></span>`): Footer template.
- `/ping` - **Check if the server is ready to accept requests.**
- `/swagger/index.html` - **Swagger introspection**

For detail description of the pdf options check the [playwright](https://playwright.dev/docs/api/class-page?_highlight=pdf#pagepdfoptions) documentation.

## Environment variables

- `FAY_MAX_ACTIVE_PAGES` (**default:** `0`): Controls how many pages can be opened at the same time before responding with status code `429`.

## Scalability

Fay is staless and can be scaled infinitely. If you run fay on your infrastructure you should keep some things in mind.
Fay will open as many pages as possible depending on the available host resources. You can control the maximum active pages by the environment variable `FAY_MAX_ACTIVE_PAGES=20`. As a general thumb you can calculate the base memory consumption in the following way. The chrome instance takes around `~45MB`. Every additional page `~15MB`. In case of the limit is reached the server will respond with status code `429`. The client is responsible to implement a backoff strategy.

A single fay instance (static HTML mode) is capable to serve 20 parallel virtual users with an average request duration of `~0.5`s. The memory consumption was `~500MB`.

For more informations check the [load-test](./loadtesting/README.md).

## Best practice

In order to produce reproducible results try to avoid downloading external resources you can't control and executing javascript. You can inline images and styles in the document. Use the options `offline=true` and `javascript=false` to enforce that. Fonts can be embedded in the docker image to make them accesible to the chromium browser.

## Credits

- https://github.com/mxschmitt/playwright-go
- https://github.com/thecodingmachine/gotenberg
- <a target="_blank" href="https://icons8.com/icons/set/elf--v2">Elf icon</a> icon by <a target="_blank" href="https://icons8.com">Icons8</a>
