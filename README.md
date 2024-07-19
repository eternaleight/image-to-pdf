<p align="center">
  <img src="https://github.com/user-attachments/assets/4a808643-b165-43bb-9b4d-cf41ab6e7b28" alt="Image-to-PDF Logo" width="300px">
</p>

windowsだと最大600×600のPDFしか標準機能で作れなかったので画像のPDF化（大量枚数 高画質対応）を作成しました。
## 画像のPDF化
### 1. **コードの実行**：

次のようにコマンドライン引数を指定して実行します。

-imageFolder PDF化したい画像が入っているディレクトリ

-batchSize PDFに入る枚数の指定(多すぎるとPDF化の処理ができないため指定する必要がある、PDFを分割してから後で結合するため)

-outputFolders 出力先のディレクトリ

```zsh
go run main.go -imageFolder="/c/〇〇/inputImage" -batchSize=20 -outputFolder="/c/〇〇/pdfs"
```

### 説明

- `flag.String` と `flag.Int` を使用してコマンドライン引数を定義しています。
  - `imageFolder` は画像フォルダのパスを指定します。
  - `batchSize` は1つのPDFに含める画像の枚数を指定します。
  - `outputFolder` は生成されたPDFバッチを保存するフォルダを指定します。
- `flag.Parse` を呼び出してコマンドライン引数を解析します。
- 必要な引数が指定されていない場合、エラーメッセージを表示して終了します。
- `os.MkdirAll` を使用して、出力フォルダが存在しない場合は作成します。

入力画像フォルダ、バッチサイズ、およびアウトプットフォルダをコマンドラインから指定
### main.go

```main.go
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
		pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}})

		for _, imageFile := range batchFiles {
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
```

## PDFtkを使用してアウトプットされたPDFを結合

### WSLまたはLinux環境での手順

WSLまたはLinux環境でPDFtkを使用して複数のPDFを結合する手順は以下の通りです。

1. **PDFtkのインストール**

まず、PDFtkがインストールされていない場合はインストールします。

```bash
sudo apt-get install pdftk
```

2. **PDFの結合**

生成されたバッチPDFファイルを一つに結合します。以下のコマンドを実行して、複数のPDFファイルを一つのPDFファイルに結合します。

```bash
pdftk batch_0.pdf batch_1.pdf batch_2.pdf batch_3.pdf batch_4.pdf batch_5.pdf batch_6.pdf batch_7.pdf cat output combined.pdf
```

これにより、全てのバッチPDFファイルが `combined.pdf` という名前の一つのPDFファイルに結合されます。


### Homebrewを使用しての手順

1. **PDFtkのインストール**

```bash
brew install pdftk-java
```

2. **PDFの結合**

生成されたバッチPDFファイルを一つに結合します。以下のコマンドを実行して、複数のPDFファイルを一つのPDFファイルに結合します。

```bash
pdftk batch_0.pdf batch_1.pdf batch_2.pdf batch_3.pdf batch_4.pdf batch_5.pdf batch_6.pdf batch_7.pdf cat output combined.pdf
```

### 結合手順の説明

1. **バッチPDFのリスト**:
   - 生成されたバッチPDFファイル（ `batch_0.pdf` から `batch_7.pdf` など）をリストアップします。

2. **PDFの結合コマンド**:
   - `pdftk` コマンドを使用して、リストアップしたPDFファイルを `cat` コマンドで結合し、 `combined.pdf` に出力します。

以下は、すべての手順をまとめた例です。

```bash
pdftk batch_0.pdf batch_1.pdf batch_2.pdf batch_3.pdf batch_4.pdf batch_5.pdf batch_6.pdf batch_7.pdf cat output combined.pdf
```

### 追加の手順

もしバッチの数が多く手動でリストアップするのが大変な場合は、以下のようなシェルスクリプトを使用することもできます。

```bash
pdftk $(ls batch_*.pdf | sort) cat output combined.pdf
```

このコマンドは、 `batch_` で始まるすべてのPDFファイルを自動的に結合します。

これで、複数のバッチPDFファイルを一つのPDFファイルに結合することができます。
