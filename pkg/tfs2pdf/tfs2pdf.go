// Main logic of application
package tfs2pdf

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"

	"github.com/AF250329/tfs2pdf/pkg/pdf"
	"github.com/AF250329/tfs2pdf/pkg/tfs"
)

// Function execute all logic of application
func Run(args []string, outputFolder, tfsAddress, tfsToken string) error {

	fmt.Println("started")

	itemId, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	data, err := tfs.ReadTfsItem(itemId, tfsAddress, tfsToken)
	if err != nil {
		return err
	}

	log.Default().Printf("received data from TFS server")

	p := &pdf.PdfData{
		OutputFolder: outputFolder,
		FileName:     fmt.Sprintf("%s.pdf", args[0]),
	}

	err = p.Create(data)
	if err != nil {
		return err
	}

	log.Default().Printf("successfully create PDF file at: %v", filepath.Join(p.OutputFolder, p.FileName))

	return nil
}
