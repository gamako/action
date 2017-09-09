# action

## 開発環境の準備

- go
    - cgoビルド環境。クロスコンパイルするためにはターゲット用のcコンパイルができる必要がある。
- glide
    - goの依存ライブラリを準備するためのツール
- sdl2
    - 実行環境にはsdl2が必要。

### インストール方法

#### macの場合

```
brew install golang glide sdl2
```

## ビルド方法

```
cd src/app
glide install
cd ../..
go build
```

