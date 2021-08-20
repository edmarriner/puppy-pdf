package pdfpuppy

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"io"
	"net/url"
)

type Options struct {
	Page   Page
	Header Header
	Footer Footer
}

type puppy struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	options    Options
}

type Page struct {
	Source Source
}

type Source struct {
	Html string
	Url  string
	File io.Reader
}

type Header struct {
	Source Source
}

type Footer struct {
	Source Source
}

func GeneratePDF(options Options, ctx context.Context) ([]byte, error) {

	err := validateOptions(options)

	if err != nil {
		return nil, err
	}

	// Create the Chromedp context
	dpContext, dpCancel := chromedp.NewContext(ctx)

	pup := &puppy{
		cancelFunc: dpCancel,
		ctx:        dpContext,
		options:    options,
	}

	buffer, err := pup.render()

	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func validateOptions(opts Options) error {

	if opts.Page.Source.Url != "" {
		_, err := url.ParseRequestURI(opts.Page.Source.Url)

		if err != nil {
			return err
		}
	}

	return nil
}

func (pdf *puppy) render() ([]byte, error) {

	var buffer []byte
	task := chromedp.Tasks{
		chromedp.Navigate(pdf.options.Page.Source.Url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(false).Do(ctx)
			if err != nil {
				return err
			}
			buffer = buf
			return nil
		}),
	}

	err := chromedp.Run(pdf.ctx, task)

	if err != nil {
		return nil, err
	}

	return buffer, nil
}
