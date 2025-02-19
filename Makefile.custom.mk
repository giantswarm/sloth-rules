.PHONY: generate
generate:
	go run main.go

test: install-tools template-chart test-runbooks ## Run tests

install-tools:
	scripts/tests/install-tools.sh

template-chart: install-tools ## prepare the helm chart
	scripts/tests/template-chart.sh

test-runbooks: install-tools template-chart ## Check if runbooks are valid
	scripts/tests/check-runbooks.sh

test-ci-runbooks: install-tools template-chart ## Check if runbooks are valid in CI
	scripts/tests/check-runbooks.sh --ci

sloth-validate: install-tools template-chart ## Validate sloth resources
	scripts/tests/bin/sloth validate --input=scripts/tests/output/sloth-rules/templates/

test-clean:
	rm -rf scripts/tests/bin
	rm -rf scripts/tests/output
