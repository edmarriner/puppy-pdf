package main

import (
	"context"
	"io/ioutil"
	"pdfpuppy"
)

func main() {

	opts := pdfpuppy.Options{
		Page: pdfpuppy.Page{
			Source: pdfpuppy.Source{
				Url: "https://apple.co.uk",
			},
		},
		Header: pdfpuppy.Header{
			Source: pdfpuppy.Source{
				Url: "https://apple.co.uk",
			},
		},
		Footer: pdfpuppy.Footer{
			Source: pdfpuppy.Source{
				Url: "https://apple.co.uk",
			},
		},
	}

	buffer, err := pdfpuppy.GeneratePDF(opts, context.Background())

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("sample.pdf", buffer, 0644)

	if err != nil {
		panic(err)
	}

}
