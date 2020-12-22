package controller

import (
	"io/ioutil"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"github.com/mxschmitt/playwright-go"
)

type Http struct {
	Browser *playwright.Browser
}

func (ctrl *Http) Ping(c echo.Context) error {
	if ctrl.Browser.IsConnected {
		return c.HTML(http.StatusOK, "")
	}
	return c.HTML(http.StatusServiceUnavailable, "")
}

type ConvertRequest struct {
	URL               string `form:"url" query:"url" valid:"url"`
	Locale            string `form:"locale" query:"locale"`
	JavaScriptEnabled bool   `form:"javaScriptEnabled" query:"javaScriptEnabled"`
	Format            string `form:"format" query:"format" valid:"in(Letter|Legal|Tabloid|Ledger|A0|A1|A2|A4|A5|A6)"`
	Offline           bool   `form:"offline" query:"offline"`
	Media             string `form:"media" query:"media" valid:"in(screen,print)"`

	MarginTop    string `form:"marginTop" query:"marginTop"`
	MarginRight  string `form:"marginRight" query:"marginRight"`
	MarginBottom string `form:"marginBottom" query:"marginBottom"`
	MarginLeft   string `form:"marginLeft" query:"marginLeft"`

	FooterTemplate string `form:"footerTemplate"`
	HeaderTemplate string `form:"headerTemplate"`
	HTML           string `form:"html"`
}

func (ctrl *Http) ConvertHTML(c echo.Context) error {

	/*
		Extract data from request
	*/
	u := new(ConvertRequest)
	if err := c.Bind(u); err != nil {
		return c.HTML(http.StatusBadRequest, "")
	}

	/*
		Request validation
	*/
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		c.Logger().Errorf("request validation: %s", err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	/*
		Footer template
	*/
	footerTemplateFile, err := c.FormFile("footerTemplate")
	if err == nil {
		footerTemplateSrc, err := footerTemplateFile.Open()
		if err != nil {
			c.Logger().Errorf("could not open footerTemplate: %s", err)
			return c.String(http.StatusBadRequest, "")
		}
		defer footerTemplateSrc.Close()

		if b, err := ioutil.ReadAll(footerTemplateSrc); err == nil {
			u.FooterTemplate = string(b)
		} else {
			c.Logger().Errorf("could not read footerTemplate: %s", err)
		}
	} else if err != http.ErrMissingFile {
		c.Logger().Errorf("could not get form value footerTemplate: %s", err)
		return c.String(http.StatusBadRequest, "")
	}

	/*
		Header template
	*/
	headerTemplateFile, err := c.FormFile("headerTemplate")
	if err == nil {
		headerTemplateSrc, err := headerTemplateFile.Open()
		if err != nil {
			c.Logger().Errorf("could not open headerTemplate: %s", err)
			return c.String(http.StatusBadRequest, "")
		}
		defer headerTemplateSrc.Close()

		if b, err := ioutil.ReadAll(headerTemplateSrc); err == nil {
			u.HeaderTemplate = string(b)
		} else {
			c.Logger().Errorf("could not read headerTemplate: %s", err)
		}
	} else if err != http.ErrMissingFile {
		c.Logger().Errorf("could not get form value headerTemplate: %s", err)
		return c.String(http.StatusBadRequest, "")
	}

	/*
		HTML template
	*/
	htmlTemplateFile, err := c.FormFile("html")
	if err == nil {
		htmlTemplateSrc, err := htmlTemplateFile.Open()
		if err != nil {
			c.Logger().Errorf("could not open html: %s", err)
			return c.String(http.StatusBadRequest, "")
		}
		defer htmlTemplateSrc.Close()

		if b, err := ioutil.ReadAll(htmlTemplateSrc); err == nil {
			u.HTML = string(b)
		} else {
			c.Logger().Errorf("could not read html: %s", err)
		}
	} else if err != http.ErrMissingFile {
		c.Logger().Errorf("could not get form value html: %s", err)
		return c.String(http.StatusBadRequest, "")
	}

	/*
		Defaults
	*/
	if u.FooterTemplate == "" {
		u.FooterTemplate = "<span></span>"
	}
	if u.HeaderTemplate == "" {
		u.HeaderTemplate = "<span></span>"
	}
	if u.Format == "" {
		u.Format = "A4"
	}
	if u.Media == "" {
		u.Format = "print"
	}

	/*
		Create new browser context to avoid side-effects (cookies, storage etc...)
	*/
	browserContextOptions := playwright.BrowserNewContextOptions{
		JavaScriptEnabled: playwright.Bool(u.JavaScriptEnabled),
		Locale:            playwright.String(u.Locale),
	}
	browserContext, err := ctrl.Browser.NewContext(browserContextOptions)
	if err != nil {
		c.Logger().Errorf("could not create new context: %s", err)
		return c.HTML(http.StatusInternalServerError, "")
	}
	defer browserContext.Close()

	/*
		Open a new page. Playwright will handle the queue.
	*/
	page, err := browserContext.NewPage(playwright.BrowserNewPageOptions{
		Offline: playwright.Bool(u.Offline),
	})
	if err != nil {
		c.Logger().Error("could not create new page")
		return c.HTML(http.StatusInternalServerError, "")
	}
	if u.URL != "" {
		_, err = page.Goto(u.URL, playwright.PageGotoOptions{
			Timeout: playwright.Int(10000),
		})
		if err != nil {
			c.Logger().Errorf("could not go to page: %s", err)
			return c.HTML(http.StatusBadGateway, "")
		}
	} else {
		err := page.SetContent(u.HTML, playwright.PageSetContentOptions{
			Timeout: playwright.Int(10000),
		})
		if err != nil {
			c.Logger().Errorf("could not set page content: %s", err)
			return c.HTML(http.StatusInternalServerError, "")
		}

	}

	page.EmulateMedia(playwright.PageEmulateMediaOptions{Media: u.Media})

	/*
		Render page
	*/
	pdfBytes, err := page.PDF(playwright.PagePdfOptions{
		DisplayHeaderFooter: playwright.Bool(true),
		PrintBackground:     playwright.Bool(true),
		FooterTemplate:      playwright.String(u.FooterTemplate),
		HeaderTemplate:      playwright.String(u.HeaderTemplate),
		Format:              playwright.String(u.Format),
		Margin: &playwright.PagePdfMargin{
			Top:    u.MarginTop,
			Right:  u.MarginRight,
			Bottom: u.MarginBottom,
			Left:   u.MarginLeft,
		},
	})

	return c.Blob(200, "application/pdf", pdfBytes)
}
