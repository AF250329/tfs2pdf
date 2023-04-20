// It based on wkhtmltopdf package from https://wkhtmltopdf.org/index.html
package pdf

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/AF250329/tfs2pdf/pkg/tfs"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

var (
	pathToTemplateFolder string
)

type PdfData struct {
	OutputFolder string
	FileName     string
}

func getPathToTemplateFolder() string {
	if pathToTemplateFolder == "" {
		dir, _ := os.Executable()

		dir = filepath.Dir(dir)

		pathToTemplateFolder = filepath.Join(dir, "templates")
	}

	return pathToTemplateFolder
}

func (p *PdfData) Create(data *tfs.Data) error {

	pathToTemplateFolder = getPathToTemplateFolder()

	htmlFile := p.parseTemplate(data)

	htmlFile = p.convertPath(htmlFile)

	log.Default().Printf("trying to print it to PDF printer")

	pdfRawData := p.loadPage(htmlFile)

	outputFileName := filepath.Join(p.OutputFolder, p.FileName)

	err := os.WriteFile(outputFileName, pdfRawData, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (p PdfData) convertPath(htmlFilePath string) string {

	htmlFilePath = strings.ReplaceAll(htmlFilePath, "\\", "/")

	htmlFilePath = fmt.Sprintf("file://%s", htmlFilePath)

	return htmlFilePath
}

func (p PdfData) parseTemplate(data *tfs.Data) string {
	sourceFileName := filepath.Join(pathToTemplateFolder, "template.htm")

	tmpl, err := template.ParseFiles(sourceFileName)
	if err != nil {
		panic(err)
	}

	destinationFileName := filepath.Join(os.TempDir(), "tmp-template.tmp.html")
	os.Remove(destinationFileName)

	destinationFile, err := os.Create(destinationFileName)
	if err != nil {
		panic(err)
	}

	CopyDirectory(filepath.Join(pathToTemplateFolder, "template_files"), filepath.Join(os.TempDir(), "template_files"))

	err = tmpl.Execute(destinationFile, data)
	if err != nil {
		panic(err)
	}

	destinationFile.Close()

	return destinationFileName
}

func (p PdfData) loadPage(htmlFilePath string) []byte {
	taskCtx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	var pdfBuffer []byte

	if err := chromedp.Run(taskCtx, printToPDF(htmlFilePath, &pdfBuffer)); err != nil {
		log.Fatal(err)
	}

	return pdfBuffer
}

func printToPDF(url string, res *[]byte) chromedp.Tasks {

	start := time.Now()

	return chromedp.Tasks{

		// emulation.SetUserAgentOverride("alex 1.0"),

		chromedp.Navigate(url),

		// wait for footer element is visible (ie, page is loaded)
		// chromedp.WaitVisible(`body`, chromedp.ByQuery),

		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(true).Do(ctx)

			if err != nil {
				return err
			}

			*res = buf

			fmt.Printf("\nTook: %f secs\n", time.Since(start).Seconds())

			return nil
		}),
	}
}
