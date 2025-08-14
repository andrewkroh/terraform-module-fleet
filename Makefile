TF_DOCS := go run github.com/terraform-docs/terraform-docs@v0.17.0
PACKAGES_DIR ?= ../integrations/packages

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
	rm -f fleet_integration/*/*
	rm -f fleet_input/*/*
	fleet-terraform-generator generate batch --packages-dir "${PACKAGES_DIR}" --out . \
		"aws/cloudtrail/*/aws-s3" \
		"aws/guardduty/guardduty/*" \
		"aws_bedrock/*/invocation/*" \
		"aws_logs/aws_logs/generic/*" \
		"barracuda_cloudgen_firewall/*/*/lumberjack" \
		"barracuda_cloudgen_firewall/*/*/lumberjack" \
		"cloud_security_posture/*/*/*" \
		"crowdstrike/*/fdr/aws-s3" \
		"crowdstrike/*/*/cel" \
		"crowdstrike/*/*/streaming" \
		"entityanalytics_entra_id/*/entity/*" \
		"github/*/issues/httpjson" \
		"google_workspace/*/*/httpjson" \
		"m365_defender/*/event/*" \
		"m365_defender/*/incident/*" \
		"qualys_vmdr/*/*/cel" \
		"sentinel_one/*/*/httpjson" \
		"system/*/application/winlog" \
		"system/*/security/winlog" \
		"system/system/diskio/system_metrics" \
		"system/system/process_summary/system_metrics" \
		"tenable_io/*/vulnerability/*" \
		"ti_abusech/*/*/cel" \
		"ti_crowdstrike/*/*/cel" \
		"ti_recordedfuture/*/*/cel" \
		"ti_threatconnect/*/*/cel" \
		"windows/*/powershell*/winlog" \
		"windows/*/*sysmon*/winlog" \
		"cel/*/*" \
		"log/*/*" \
		"sql/*/*" \
		"winlog/*/*"
	# Ignore shadowing for these packages. See https://github.com/elastic/integrations/issues/6148.
	fleet-terraform-generator generate batch --packages-dir "${PACKAGES_DIR}" --ignore-var-shadow --out . \
		"aws/inspector/*/*" \
		"aws/securityhub/*/*"

.PHONY: install
install:
	cd tools/cmd/fleet-terraform-generator && go install .

.PHONY: fmt
fmt:
	terraform fmt --recursive .

.PHONY: validate
validate:
	cd examples/github && terraform init && terraform validate

.PHONY: docs
docs:
	@for i in $(shell find fleet_* -name '*.tf' -or -name '*.tf.json' -not -path '*/.terraform*' -print0 | xargs -0 -n1 dirname | sort --unique); do \
	  module=$$i; \
	  with_header=$$(test -f "$$module/.readme.md" && echo -n --header-from=".readme.md"); \
	  ${TF_DOCS} markdown table --output-check=${DOCS_CHECK_UPDATE} $$with_header --output-file="README.md" "$$module" || exit 1; \
	done

# docs-check verifies that all README.md files are up to date.
.PHONY: docs-check
docs-check:
	${MAKE} docs DOCS_CHECK_UPDATE=true
