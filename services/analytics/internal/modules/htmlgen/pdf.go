package htmlmodule

import (
	"bytes"
	"log"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func PDF(data []byte) error {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
		return err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(data)))
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}
	pdfg.WriteFile("output.pdf")
	return nil
}
