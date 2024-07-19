## 画像のＰＤＦ化（大量枚数対応）
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

## PDFtkを使用してアウトプットされたPDFを結合

### Mac環境での手順

Mac環境でPDFtkを使用して複数のPDFを結合する手順は以下の通りです。

1. **Homebrewのインストール**（インストールされていない場合）

Homebrewは、macOS用のパッケージ管理システムです。インストールされていない場合は、以下のコマンドを実行してインストールします。

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

2. **PDFtkのインストール**

Homebrewを使用してPDFtkをインストールします。

```bash
brew install pdftk-java
```

3. **PDFの結合**

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
