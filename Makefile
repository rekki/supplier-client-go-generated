.PHONY: generate-client
generate-client:
	#-i https://api.rekki.com/swagger/doc.json
	@docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
		-i https://external-supplier-api.feat.eu-west-2.rekki.com/swagger/doc.json \
		-g go \
		--git-user-id rekki \
		--git-repo-id supplier-client-go-generated \
		--package-name=client \
		-o /local/client \
		-p enumClassPrefix=true \
		-p withGoCodegenComment=true
	@sudo chown -R dima:dima .
	@rm -rf docs
	@mv client/go.* .
	@mv client/docs .
	@mv client/.gitignore .
	@mv client/README.md .
	@rm client/.travis.yml

.PHONY: validate-client
validate-client:
	#-i https://api.rekki.com/swagger/doc.json
	@docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli validate \
		-i https://external-supplier-api.feat.eu-west-2.rekki.com/swagger/doc.json
