#!/bin/bash

echo "コンテナイメージの作成"
docker-compose build | echo $?

echo "コンテナのバックグラウンドでの実行"
docker-compose up -d | echo $?

echo "コンテナに接続"
docker-compose exec boiler-plate ash

exit 0
