# go-api-architecture

This repository is for learning purposes only.

このリポジトリは学習用です。

## 使用技術一覧
- go
- mysql


## フォルダ構成
```shell
root
├── cmd
│   └── app
├── docs
│   ├── README.md
│   └── learning
├── internal
│   ├── http
│   │   └── app
│   │       ├── handler
│   │       └── router
│   └── websocket
│       └── app
│           ├── handler
│           └── router
├── pkg
└── README.md
```

### 各ディレクトリの役割
- `cmd/app`: アプリのエントリーポイント（`main.go`）
- `internal/http/app/router`: HTTPルーター定義
- `internal/http/app/handler`: HTTPハンドラー
- `internal/websocket/app/router`: WebSocketルーター定義
- `internal/websocket/app/handler`: WebSocketハンドラー
- `pkg`: 共有ユーティリティや共通コード
- `docs`: ドキュメント

## クイックスタート

1. .envをコピーする
  ```shell
  $ cp .env.example .env
  ```

2. DBを起動する（任意）
現状のコードではDB未使用のため省略可。

3. アプリを起動する（1行）
```shell
$ go run ./cmd/app
```

サーバは `http://localhost:8085` で待ち受けます。

## エンドポイント

- GET `/test`
  - 用途: 動作確認
  - 例:
  ```shell
  $ curl -sS http://localhost:8085/test
  success
  ```

- POST `/api`
  - 用途: サンプルAPI（`.env` の `APP_NAME` を返却）
  - 例:
  ```shell
  $ curl -sS -X POST http://localhost:8085/api
  {"message":"Success","app_name":"APP"}
  ```

## コミット直前に実行

### pint 実行
編集中
```shell
```

### lint 実行
編集中
```shell
```

## コマンドリスト

```shell
$ go mod init github.com/Fybrid/go-api-architecture
```

### MySQL起動（ローカルインストール版）

1. mysql起動
```shell
$ mysql.server start
```
2. mysqlシャットダウン
```shell
$ mysql.server stop
```
## ドキュメント

- [ドキュメント入口](./docs/README.md)
- [コミットルール](docs/branch.md)
- [ブランチルール](docs/command.md)
- 学習ログ: `docs/learning/`
