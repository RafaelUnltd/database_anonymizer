run:
	docker compose up -d

seed-input-db:
	psql 'postgres://root:root@localhost:5445/anonymize-input' -f ./scripts/init_db/dump.sql