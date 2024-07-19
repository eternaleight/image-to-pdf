package main

import (
	"fmt"
	"github.com/signintech/gopdf"
	_ "golang.org/x/image/bmp"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
)

func main() {
	imageFolder := ""         // 画像フォルダのパスを設定
	outputPDF := "output.pdf" // 出力PDFのファイル名を設定

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) // 初期設定でA4サイズを使用

	// 画像ファイルの取得
	imageFiles, err := filepath.Glob(filepath.Join(imageFolder, "*"))
	if err != nil {
		fmt.Println("Error reading image files:", err)
		return
	}

	fmt.Printf("Found %d files\n", len(imageFiles))

	for _, imageFile := range imageFiles {
		fmt.Printf("Processing file: %s\n", imageFile)
		file, err := os.Open(imageFile)
		if err != nil {
			fmt.Println("Error opening image file:", err)
			continue
		}
		img, format, err := image.Decode(file)
		if err != nil {
			fmt.Println("Error decoding image file:", err)
			file.Close()
			continue
		}
		file.Close()

		width := float64(img.Bounds().Dx()) * 0.75
		height := float64(img.Bounds().Dy()) * 0.75

		pdf.AddPageWithOption(gopdf.PageOption{PageSize: &gopdf.Rect{W: width, H: height}})
		err = pdf.Image(imageFile, 0, 0, &gopdf.Rect{W: width, H: height})
		if err != nil {
			fmt.Println("Error adding image to PDF:", err)
			continue
		}

		fmt.Printf("Added %s image to PDF: %s\n", format, imageFile)
	}

	err = pdf.WritePdf(outputPDF)
	if err != nil {
		fmt.Println("Error saving PDF:", err)
		return
	}

	fmt.Println("PDF generated successfully:", outputPDF)
}
