# KeyEvent

KeyEventは、Googleカレンダーに素早くイベントを追加するためのデスクトップアプリケーションです。

## 特徴
- キーボード操作に最適化
- シンプルで直感的なUI
- Googleカレンダーとの連携

## 対応プラットフォーム
- macOS (Apple Silicon)
- Windows (64-bit)

## インストール

### プリビルドバイナリ
1. [最新のリリース](https://github.com/imaikosuke/KeyEvent/releases/latest)ページから、お使いのOSに合わせたバイナリをダウンロードします。
   - macOS (Apple Silicon): `KeyEvent-mac`
   - Windows: `KeyEvent.exe`
2. ダウンロードしたファイルを解凍します（必要な場合）。
3. macOSユーザーの場合、初回実行時に「開発元を確認できないため開けません」というメッセージが表示される場合があります。その場合は、Controlキーを押しながらアプリケーションをクリックし、「開く」を選択してください。
4. 解凍したフォルダ内の実行ファイルを実行します。

### ソースからのビルド
1. Go言語の開発環境をセットアップします。
2. このリポジトリをクローンします：
   ```
   git clone https://github.com/imaikosuke/KeyEvent.git
   ```
3. プロジェクトディレクトリに移動し、依存関係をインストールします：
   ```
   cd KeyEvent
   go mod tidy
   ```
4. アプリケーションをビルドします：
   - macOS (Apple Silicon):
     ```
     GOOS=darwin GOARCH=arm64 go build -o KeyEvent-mac
     ```
   - Windows:
     ```
     GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 go build -o KeyEvent.exe
     ```

## 使用方法
1. アプリケーションを起動します。
2. Googleアカウントでの認証を行います。
3. イベントの詳細を入力し、色を選択します。
4. 「送信」ボタンをクリックするか、Ctrl+Enterを押してイベントを作成します。
