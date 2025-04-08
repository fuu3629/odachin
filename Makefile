gen-service:
	cd apps/service && buf generate proto

gen-web:
	cd apps/web && npm run gen:services

gen-all: gen-service gen-web