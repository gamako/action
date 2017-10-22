# action

go言語をつかった簡単なアクションゲームのプロジェクト。
rasberry pieで動かすことを目的としています。

# 開発環境の準備

- go
    - cgoビルド環境。クロスコンパイルするためにはターゲット用のcコンパイルができる必要がある。
- glide
    - goの依存ライブラリを準備するためのツール
- sdl2
    - 実行環境にはsdl2が必要。
- VisualStudioCodeなど

## コマンドラインツール インストール方法(macの場合)

homebrewでコマンドラインツールのインストール

```bash
brew install golang glide sdl2
```

## Visual Studio Codeの準備(macの場合)

1. Webからインストール (https://code.visualstudio.com/download)
2. 拡張機能からGoの拡張をインストール
3. 設定を開いて、ユーザー設定に以下の項目を設定
    - go.goroot ... goのインストール先
    - go.toolsGopath ...（オプション）指定しなければ、gopathにインストールされる。複数のプロジェクトを扱うのであれば、共通にパスを指定しておくと良いでしょう。

例
```json
{
    "go.goroot": "/usr/local/opt/go/libexec",
    "go.toolsGopath": "/Users/gamako/.go",
}
```

3. 設定の右上からワークスペースの設定を選んで、以下の項目を追加
    - go.gopath ... このプロジェクトのパス

例
```json
{
    "go.gopath":"/Users/gamako/project/raspberrypi/action"
}
```

4. デバッグ環境の準備

-  go delveのインストール
    - brew install -v go-delve/delve/delve --HEAD
      をしてもエラーになるので
    - dlv-certという名前のcodesigin用自己証明書をつくって、keychainに登録（コマンドでのやりかたは、homebrewのformulaのrbスクリプトに書いてある）
    - command-lineツールをappleからダウンロードしてインストール。
      - osバージョン、Xcodeバージョンが合ったものをインストールする。

参考

https://qiita.com/TsuyoshiUshio@github/items/ba15b1a7e6c6e5ffaf43

## ビルド方法

1. 環境変数GOPATHにこのプロジェクトのディレクトリを指定しておく
2. src/appでglide installを実行して、依存ソースを取得
3. go build app

例
```
export GOPATH=$(pwd)
cd src/app
glide install
cd ../..
go build app
```

