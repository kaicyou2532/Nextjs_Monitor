＃Next.js モニター
このリポジトリは、Next.js 開発サーバーを監視する小さな Go プログラムを提供します。指定した URL を定期的にチェックし、サーバーが停止している場合は設定されたディレクトリで npm run dev を起動します。

前提条件
Go 1.23 以降

Node.js と npm

npm run dev で起動する Next.js プロジェクト

ビルド方法
bash
コピーする
編集する
# リポジトリをクローンしてディレクトリに移動
cd Nextjs_Monitor

# 依存関係をインストールし、monitor バイナリをビルド
go build -o monitor
実行方法
モニターはコマンドラインフラグで設定します:

-dir – npm run dev を実行するディレクトリ（デフォルト: カレントディレクトリ）

-url – ヘルスチェックに使用する URL（デフォルト: http://localhost:3000）

-interval – チェックを実行する間隔（デフォルト: 1m）

-pattern – 実行中の npm プロセスを検出するために pgrep で使用するパターン（デフォルト: npm.*run.*dev）

例:

bash
コピーする
編集する
./monitor -dir /path/to/nextjs/app -url http://localhost:3000 -interval 30s
このプログラムは指定した URL を 30 秒ごとにチェックし、サーバーが応答せず、かつ npm run dev プロセスが検出されない場合にのみ、指定ディレクトリでサーバーを再起動します。

バックグラウンドサービスとして systemd で起動すれば、Next.js 開発サーバーを常に稼働させ続けることも可能です。

systemd での設定例
リポジトリにはサンプルのサービスファイル nextjs-monitor.service が含まれています。まず、ExecStart 行を monitor バイナリへのパスと Next.js アプリのディレクトリに合わせて修正し、以下のように配置します:

bash
コピーする
編集する
sudo cp nextjs-monitor.service /etc/systemd/system/
その後、systemd をリロードし、サービスを有効化して起動します:

bash
コピーする
編集する
sudo systemctl daemon-reload
sudo systemctl enable --now nextjs-monitor.service
これでモニターがバックグラウンドで実行され、万が一終了しても自動的に再起動されます。
