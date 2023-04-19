// Main logic of application
package tfs2pdf

import (
	"fmt"
	"strconv"

	"github.com/AF250329/tfs2pdf/pkg/pdf"
	"github.com/AF250329/tfs2pdf/pkg/tfs"
)

// Function execute all logic of application
func Run(args []string, outputFolder string) error {
	fmt.Println("started")

	itemId, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	data := tfs.ReadTfsItem(itemId)

	p := &pdf.PdfData{
		OutputFolder: outputFolder,
		FileName:     fmt.Sprintf("%s.pdf", args[0]),
	}

	return p.Create(data)
}
