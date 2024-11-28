COVERAGE_DIR ?= .coverage

# cp from: https://github.com/yyle88/rese/blob/90d3ef2ac448adb2576117bb026d9778220d600e/Makefile#L4
test:
	@-rm -r $(COVERAGE_DIR)
	@mkdir $(COVERAGE_DIR)
	make test-with-flags TEST_FLAGS='-v -race -covermode atomic -coverprofile $$(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m'

test-with-flags:
	@go test $(TEST_FLAGS) ./...
