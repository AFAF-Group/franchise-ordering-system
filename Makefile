swagger-doc:
	@swag init --parseDependency --parseInternal --parseDepth 1 -g cmd/api/main.go --dir ./ -o docs/swagger --exclude router,views
swagger-fmt:
	@swag fmt --generalInfo cmd/api/main.go --dir ./ --exclude router,views