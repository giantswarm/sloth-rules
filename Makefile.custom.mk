.PHONY: generate
generate:
	go run main.go

test: install-tools template-chart test-opsrecipes ## Run tests

install-tools:
	scripts/tests/install-tools.sh

template-chart: install-tools ## prepare the helm chart
	scripts/tests/template-chart.sh

test-opsrecipes: install-tools template-chart ## Check if opsrecipes are valid
	scripts/tests/check-opsrecipes.sh

test-ci-opsrecipes: install-tools template-chart ## Check if opsrecipes are valid in CI
	scripts/tests/check-opsrecipes.sh --ci

sloth-validate: install-tools template-chart ## Validate sloth resources
	scripts/tests/bin/sloth validate --input=scripts/tests/output/sloth-rules/templates/

test-clean:
	rm -rf scripts/tests/bin
	rm -rf scripts/tests/output
