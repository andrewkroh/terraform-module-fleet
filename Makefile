.PHONY: all
all: fmt docs modules validate

.PHONY: modules
modules: fleet-modules terraform-docs

.PHONY: fleet-modules
fleet-modules: install
	fleet-terraform-generator generate batch --packages-dir ../integrations/packages --out fleet_integrations \
		"aws/cloudtrail/*/aws-s3" \
		"barracuda_cloudgen_firewall/*/*/lumberjack" \
		"github/*/issues/httpjson" \
		"google_workspace/*/*/httpjson" \
		"ti_abusech/*/*/httpjson" \
		"ti_recordedfuture/*/*/httpjson" \
		"cel/*/*" \
		"log/*/*" \
		"sql/*/*"

.PHONY: install
install:
	cd fleet-terraform-generator && go install .

.PHONY: fmt
fmt:
	terraform fmt --recursive .

.PHONY: validate
validate:
	cd examples/github && terraform init && terraform validate

.PHONY: docs
docs:
	cd fleet-terraform-generator/docs && go generate .

.PHONY: terraform-docs
terraform-docs:
	go install github.com/terraform-docs/terraform-docs@latest
	terraform-docs markdown table --output-file="README.md" fleet_agent_policy
	terraform-docs markdown table --output-file="README.md" fleet_output
	terraform-docs markdown table --output-file="README.md" fleet_package_policy
	terraform-docs markdown table --output-file="README.md" fleet_server_host
	@for i in $(shell find fleet_integrations/ -name module.tf.json); do \
	  module=$$(dirname $$i); \
	  terraform-docs markdown table --output-file="README.md" "$$module" || exit 1; \
	done