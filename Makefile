TF_DOCS := go run github.com/terraform-docs/terraform-docs@v0.17.0

DOCS_CHECK_UPDATE ?= false

.PHONY: all
all: fmt modules docs validate

.PHONY: modules
modules: fleet-modules docs

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
		"m365_defender/*/event/*" \
		"m365_defender/*/incident/*" \
		"system/*/application/winlog" \
		"system/*/security/winlog" \
		"system/system/diskio/system_metrics" \
		"system/system/process_summary/system_metrics" \
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
	@for i in $(shell find fleet_* -name '*.tf' -or -name '*.tf.json' -not -path '*/.terraform*' -print0 | xargs -0 -n1 dirname | sort --unique); do \
	  module=$$i; \
	  with_header=$$(test -f "$$module/.readme.md" && echo -n --header-from=".readme.md"); \
	  ${TF_DOCS} markdown table --output-check=${DOCS_CHECK_UPDATE} $$with_header --output-file="README.md" "$$module" || exit 1; \
	done

# docs-check verifies that all README.md files are up to date.
.PHONY: docs-check
docs-check:
	${MAKE} docs DOCS_CHECK_UPDATE=true