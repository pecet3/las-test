run:
	go run ./cmd

rm:
	rm -rf store.db 


reset_db:
	rm -rf ./database/*
	touch ./database/store.db
	sqlite3 ./database/store.db < ./data/migrations/0001_initial.up.sql