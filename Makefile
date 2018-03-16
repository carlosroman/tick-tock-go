ifndef TEST_RESULTS
	TEST_RESULTS := 'target'
endif

.PHONY: test

test-report-dir:
	mkdir -p ${TEST_RESULTS}

test: test-report-dir
	go test \
		-race -v  \
		./...
