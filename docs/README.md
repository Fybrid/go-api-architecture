# ドキュメント（docs）

このフォルダは、学習ログや補助ドキュメントの置き場です。最小・テキスト中心・継続更新を基本方針とします。

構成

- 学習ログ: `docs/learning/` — セッションごとの時刻付きメモ
- コマンド集: `docs/command.md` — よく使うコマンドや手順

命名規則

- 学習ログ: `YYYY-MM-DD-HHMM-topic.md`（ローカル時刻、例: `2025-09-06-1522-kickoff.md`）

テンプレート

- 学習ログ（日本語）: `docs/learning/TEMPLATE.md`
- 学習ログ（英語）: `docs/learning/TEMPLATE.en.md`

使い方（例）

学習ログを作る（時刻入りファイル名）:

```shell
ts=$(date +%Y-%m-%d-%H%M)
cp docs/learning/TEMPLATE.md docs/learning/${ts}-topic.md
```

運用のヒント

- セッションの最後に5分だけ振り返って書くと継続しやすいです。
- Mermaid などテキスト図を優先（画像は最小限）。

注意事項

- 秘密情報や個人情報は書かない（ログ/スクショ含む）。
- 陳腐化しやすい情報は簡潔に保つ。
