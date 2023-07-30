### Dockerfile

Version

Go: 1.19.11

InstallPackage

git
vim
curl
tmux
npm
nodejs

### Docker-compose

作業ディレクトリ: ./work

### 使用方法

コンテナイメージの構築

$ docker-compose build

コンテナの起動(バックグラウンドで)

$ docker-compose up -d

コンテナ内に移動

docker exec -it boiler-plate /bin/sh

### シェル

以下のシェルスクリプトを実行することで、イメージの作成からコンテナに入るまでの処理を実行する

$ sh run.sh

### コンテナの停止と削除

コンテナの停止

$ docker-compose down

コンテナイメージの削除

$ docker image *\<image id\>*


# How to Vim Plug

Plug in のインストール(vimでファイルを開いている状態で)

:GoInstallBinaries

:PlugInstall

:CocInstall coc-rls

## Command

ファイルの実行

:GoRun %

パッケージ全体を実行

:GoRun

ビルドする

:GoBuild

テストする

:GoTest
ドキュメント表示(関数の上にカーソルを合わせて)

:GoDoc

識別子をリネームする(対象の上にカーソルを合わせて)

:GoRename
