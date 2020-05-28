.PHONY: generate-client
generate-client:
	@rm -rf client docs go.* .gitignore README.md
	docker run --rm \
		-v ${PWD}:/local \
		--user $$(id -u):$$(id -g) \
		openapitools/openapi-generator-cli generate \
		-i https://api.rekki.com/swagger/doc.json \
		-g go \
		--git-user-id rekki \
		--git-repo-id supplier-client-go-generated \
		--package-name=client \
		-o /local/client \
		-p enumClassPrefix=true \
		-p withGoCodegenComment=true
	@cp constants.go.tpl client/constants.go
	@mv client/go.* .
	@mv client/docs .
	@mv client/.gitignore .
	@mv client/README.md .
	@rm client/.travis.yml

.PHONY: validate-client
validate-client:
	#-i https://api.rekki.com/swagger/doc.json
	@docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli validate \
		-i https://api.rekki.com/swagger/doc.json
