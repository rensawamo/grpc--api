# コンテナ内に作成
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

#  ここで local の スキーマの中に simple_bankというものが形成される
dropdb:
	docker exec -it postgres12 dropdb simple_bank

postgres:
	docker run --name postgres12  -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

.PHONY: createdb dropdb
