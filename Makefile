COVERAGE_DIR ?= .coverage.out

# cp from: https://github.com/yyle88/rese/blob/90d3ef2ac448adb2576117bb026d9778220d600e/Makefile#L4
test:
	@if [ -d $(COVERAGE_DIR) ]; then rm -r $(COVERAGE_DIR); fi
	@mkdir $(COVERAGE_DIR)
	make test-with-flags TEST_FLAGS='-v -race -covermode atomic -coverprofile $$(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m'

test-with-flags:
	@go test $(TEST_FLAGS) ./...
