package htmlmodule

import (
	"bytes"
	"log"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func PDF(data []byte) ([]byte, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(data)))
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	pdfg.Bytes()
	// pdfg.WrsiteFile("output.pdf")
	return pdfg.Bytes(), nil
}
