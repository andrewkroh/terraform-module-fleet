.PHONY: all
all: fmt docs modules validate

.PHONY: modules
modules: fleet-modules terraform-docs

# Generate fleet module. The modules are specified using selectors that are formatted as:
#    {package_name}/{policy_template}/{data_stream}/{input_type} - for integrations
#    {package_name/{policy_template}/{input_type} - for inputs
# Any slashes (/) contained in the attribute values are replaced with underscores. For
# example, if matching the input_type of "prometheus/metrics" use "prometheus_metrics".
.PHONY: fleet-modules
fleet-modules: install
	fleet-terraform-generator generate batch --packages-dir ../integrations/packages --out . \
		"aws/cloudtrail/*/aws-s3" \
		"barracuda_cloudgen_firewall/*/*/lumberjack" \
		"github/*/issues/httpjson" \
		"google_workspace/*/*/httpjson" \
		"m365_defender/*/*/*" \
		"system/*/application/winlog" \
		"system/*/security/winlog" \
		"ti_abusech/*/*/httpjson" \
		"ti_recordedfuture/*/*/httpjson" \
		"windows/*/powershell*/winlog" \
		"windows/*/*sysmon*/winlog" \
		"winlog/*/*/winlog" \
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
	@for i in $(shell find fleet_integration/ fleet_input/ -name module.tf.json); do \
	  module=$$(dirname $$i); \
	  terraform-docs markdown table --output-file="README.md" "$$module" || exit 1; \
	done