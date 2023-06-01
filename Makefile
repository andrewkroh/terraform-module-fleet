.PHONY: all
all: fmt docs modules validate

.PHONY: modules
modules: install
	fleet-terraform-generator generate batch --packages-dir ../integrations/packages --out fleet_integrations \
		"aws/cloudtrail/*/aws-s3" \
		"barracuda_cloudgen_firewall/*/*/lumberjack" \
		"github/*/issues/httpjson" \
		"google_workspace/*/*/httpjson" \
		"ti_abusech/*/*/httpjson" \
		"ti_recordedfuture/*/*/httpjson"

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