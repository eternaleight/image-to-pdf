package main

import (
	"flag"
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
	// コマンドライン引数の定義
	imageFolder := flag.String("imageFolder", "", "Path to the folder containing images")
	batchSize := flag.Int("batchSize", 20, "Number of images per PDF batch")
	outputFolder := flag.String("outputFolder", "", "Path to the folder to save PDF batches")
	flag.Parse()

	// 引数のバリデーション
	if *imageFolder == "" || *outputFolder == "" {
		fmt.Println("Error: imageFolder and outputFolder are required")
		flag.Usage()
		return
	}

	// 出力フォルダの作成（存在しない場合）
	err := os.MkdirAll(*outputFolder, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating output folder:", err)
		return
	}

	imageFiles, err := filepath.Glob(filepath.Join(*imageFolder, "*"))
	if err != nil {
		fmt.Println("Error reading image files:", err)
		return
	}

	fmt.Printf("Found %d files\n", len(imageFiles))

	batchNumber := 0

	for i := 0; i < len(imageFiles); i += *batchSize {
		batchEnd := i + *batchSize
		if batchEnd > len(imageFiles) {
			batchEnd = len(imageFiles)
		}
		batchFiles := imageFiles[i:batchEnd]

		pdf := gopdf.GoPdf{}
		pdf.Start(gopdf.Config{})

		for _, imageFile := range batchFiles {
			fmt.Printf("Processing file: %s\n", imageFile)
			file, err := os.Open(imageFile)
			if err != nil {
				fmt.Println("Error opening image file:", err)
				continue
			}
			img, _, err := image.Decode(file)
			if err != nil {
				fmt.Println("Error decoding image file:", err)
				file.Close()
				continue
			}
			file.Close()

			imgWidth := float64(img.Bounds().Dx())
			imgHeight := float64(img.Bounds().Dy())

			pdf.AddPageWithOption(gopdf.PageOption{PageSize: &gopdf.Rect{W: imgWidth, H: imgHeight}})
			err = pdf.Image(imageFile, 0, 0, &gopdf.Rect{W: imgWidth, H: imgHeight})
			if err != nil {
				fmt.Println("Error adding image to PDF:", err)
				continue
			}

			fmt.Printf("Added image to PDF: %s\n", imageFile)
		}

		batchOutputPDF := filepath.Join(*outputFolder, fmt.Sprintf("batch_%d.pdf", batchNumber))
		err = pdf.WritePdf(batchOutputPDF)
		if err != nil {
			fmt.Println("Error saving PDF:", err)
			return
		}

		fmt.Println("PDF generated successfully:", batchOutputPDF)
		batchNumber++
	}

	fmt.Println("All batches processed. Please use a PDF merging tool to combine the batch PDFs into a single file.")
}
