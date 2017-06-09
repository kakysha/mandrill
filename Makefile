# update Go dependencies - flatten the vendor structure for greater consistency
dependency_update:
	glide update --strip-vendor

# download Go dependencies into vendor - flatten the vendor structure
dependency_get:
	glide install --strip-vendor

# run all tests, excluding those from dependencies
test: dependency_get
	go test -v -race -cover $$(go list ./... | grep -v vendor)
