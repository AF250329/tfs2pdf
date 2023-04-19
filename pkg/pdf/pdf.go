// It based on wkhtmltopdf package from https://wkhtmltopdf.org/index.html
package pdf

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"

	"github.com/AF250329/tfs2pdf/pkg/tfs"
	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

const (
	EXECUTABLE_FILE_NAME = "wkhtmltopdf.exe"
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

func copyBinFile() error {

	dir, _ := os.Executable()

	dir = filepath.Dir(dir)

	sourceFileName := filepath.Join(dir, "bin", EXECUTABLE_FILE_NAME)

	destinationFileName := filepath.Join(dir, EXECUTABLE_FILE_NAME)
	_, err := os.Stat(destinationFileName)
	if errors.Is(err, fs.ErrExist) {
		// File already exist
		return nil
	}

	err = Copy(sourceFileName, destinationFileName)
	if err != nil {
		return err
	}

	return nil
}

func (p *PdfData) Create(data *tfs.Data) error {

	pathToTemplateFolder = getPathToTemplateFolder()

	copyBinFile()

	pdfDocument := createNewPdfDocument()

	page1 := p.createPage(data)
	pdfDocument.AddPage(page1)

	err := pdfDocument.Create()
	if err != nil {
		return err
	}

	outputFileName := filepath.Join(p.OutputFolder, p.FileName)

	err = pdfDocument.WriteFile(outputFileName)
	if err != nil {
		return err
	}

	return nil
}

func createNewPdfDocument() *wkhtmltopdf.PDFGenerator {
	pdfDocument, _ := wkhtmltopdf.NewPDFGenerator()
	pdfDocument.Dpi.Set(800)
	pdfDocument.MarginBottomUnit.Set("1.27cm")
	pdfDocument.MarginLeftUnit.Set("1.27cm")
	pdfDocument.MarginRightUnit.Set("1.27cm")
	pdfDocument.MarginTopUnit.Set("1.27cm")
	pdfDocument.NoCollate.Set(true)
	pdfDocument.NoOutline.Set(true)
	pdfDocument.Orientation.Set("landscape")
	pdfDocument.PageSize.Set("A4")

	return pdfDocument
}

func (p PdfData) createPage(data *tfs.Data) *wkhtmltopdf.PageReader {

	sourceFileName := filepath.Join(pathToTemplateFolder, "template.htm")

	tmpl, err := template.ParseFiles(sourceFileName)
	if err != nil {
		panic(err)
	}

	destinationFileName := filepath.Join(os.TempDir(), "tmp-template.tmp")
	os.Remove(destinationFileName)

	destinationFile, err := os.Create(destinationFileName)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(destinationFile, data)
	if err != nil {
		panic(err)
	}

	destinationFile.Close()

	pageSourceFile, err := os.Open(destinationFileName)
	if err != nil {
		panic(err)
	}

	pdfPage1 := wkhtmltopdf.NewPageReader(pageSourceFile)
	pdfPage1.EnableLocalFileAccess.Set(true)

	return pdfPage1
}

func CopyDirectory(scrDir, dest string) error {
	entries, err := os.ReadDir(scrDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(scrDir, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		// stat, ok := fileInfo.Sys().(*syscall.Stat_t)
		// if !ok {
		// 	return fmt.Errorf("failed to get raw syscall.Stat_t data for '%s'", sourcePath)
		// }

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err := CreateIfNotExists(destPath, 0755); err != nil {
				return err
			}
			if err := CopyDirectory(sourcePath, destPath); err != nil {
				return err
			}
		case os.ModeSymlink:
			if err := CopySymLink(sourcePath, destPath); err != nil {
				return err
			}
		default:
			if err := Copy(sourcePath, destPath); err != nil {
				return err
			}
		}

		// if err := os.Lchown(destPath, int(stat.Uid), int(stat.Gid)); err != nil {
		// 	return err
		// }

		fInfo, err := entry.Info()
		if err != nil {
			return err
		}

		isSymlink := fInfo.Mode()&os.ModeSymlink != 0
		if !isSymlink {
			if err := os.Chmod(destPath, fInfo.Mode()); err != nil {
				return err
			}
		}
	}
	return nil
}

func Copy(srcFile, dstFile string) error {
	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	defer out.Close()

	in, err := os.Open(srcFile)
	defer in.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}

func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateIfNotExists(dir string, perm os.FileMode) error {
	if Exists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

func CopySymLink(source, dest string) error {
	link, err := os.Readlink(source)
	if err != nil {
		return err
	}
	return os.Symlink(link, dest)
}
