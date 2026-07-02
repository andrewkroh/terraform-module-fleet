<!-- BEGIN_TF_DOCS -->
# Fleet Package Policy Module

This module installs the Fleet package's assets, creates a package policy, and
associates that package policy to an existing agent policy.

## Package Version

The `package_version` parameter controls which version of the package is
installed. It supports two modes.

### Pinned Version (Default)

Set `package_version` to a semantic version (e.g. `1.2.3`) to install exactly
that version. This is deterministic: the same configuration always produces the
same result, and upgrades only happen when the pinned version is changed.

Note that Kibana Fleet rejects installing a pinned version when a newer version
exists in the package registry, returning HTTP 400 "out-of-date and cannot be
installed or updated" unless `force = true` is used. This means pinned versions
can break over time as newer package versions are published to the registry.

### Latest Version

Set `package_version = "latest"` to resolve the latest available version of the
package from the package registry at plan time. Resolution uses the
`elasticstack_fleet_integration` data source, so the concrete version is
determined during `terraform plan` and recorded in state. Each new release of
the package in the registry shows up as an in-place upgrade of the package and
package policy on the next plan.

This mode avoids the "out-of-date" rejection described above without resorting
to `force = true` and its destructive reinstallation behavior. The trade-off is
that plans are no longer fully deterministic because the resolved version
depends on the registry contents at plan time.

### Prerelease Versions

The `prerelease` parameter controls whether prerelease package versions are
considered when resolving `package_version = "latest"`. Packages with versions
below 1.0.0 are treated as prerelease by semver, so `prerelease = true` is
required to resolve "latest" for those packages. It has no effect when
`package_version` is pinned.

## Force Reinstallation Parameter

### Overview

The `force` parameter controls whether Fleet performs a complete reinstallation
of a package even when it's already installed at the requested version. This
parameter has significant implications for package assets and should be used
with full understanding of its impacts.

### What `force = true` Does

When `force = true`, the Fleet API performs the following actions:

#### 1. Bypasses "Already Installed" Check

Normally, if a package version is already installed with status `installed`,
Fleet immediately returns success without taking any action. With
`force = true`, this check is skipped and the full installation process executes
again.

#### 2. Deletes and Recreates Kibana Assets

The installation process explicitly removes old Kibana saved objects before
reinstalling:

- **Affected Assets:** Dashboards, visualizations, searches, and other saved
  objects associated with the package
- **Impact:** Any manual customizations made to these assets (if not cloned) are
  **permanently lost**
- **Result:** Assets are reset to the package's default configuration

#### 3. Aggressively Handles Transforms

Transform handling is particularly destructive with `force = true`:

- **Action:** Existing transforms are deleted
- **Impact:** Transform destination indices are often deleted as well to ensure
  a clean state
- **Data Loss Risk:** Any data accumulated in transform destination indices will
  be **permanently lost**
- **Recovery:** Transforms are recreated from the package specification, but
  historical data is not restored

#### 4. Overwrites Ingest Pipelines and Index Templates

- **Action:** Elasticsearch assets are re-created using "put" operations
- **Impact:** Existing ingest pipelines and index templates with the same names
  are overwritten
- **Result:** Pipelines and templates revert to package defaults

#### 5. Overrides Concurrency Locks

If a package installation is stuck in the `installing` state:

- **Normal Behavior:** Fleet throws a `ConcurrentInstallOperationError` to
  prevent conflicts
- **With `force = true`:** The lock is ignored and installation restarts from
  the beginning (`CREATE_RESTART_INSTALLATION` state)
- **Use Case:** Recovering from failed or hung installations

#### 6. Allows Restricted Operations

`force = true` bypasses several Fleet restrictions:

- **Downgrades:** Allows installing older versions of packages (normally
  blocked)
- **Out-of-date Versions:** Permits installing versions that aren't the latest
- **Agentless Packages:** Allows installing agentless packages even when the
  agentless feature is disabled in Kibana configuration

### When to Use `force = true`

#### Recommended Use Cases

1. **IaC-Managed Infrastructure** (Default)
   - Terraform is the source of truth for all configuration
   - Declarative state should always match the actual state
   - Repeatability and consistency are prioritized
   - Manual customizations are not expected or allowed

2. **Package Version Changes**
   - Reinstalling the same version to reset state
   - Downgrading to an older package version
   - Ensuring all assets match a specific package version exactly

3. **Development and Testing**
   - Iterating on package development with frequent reinstalls
   - Testing package installations in non-production environments
   - CI/CD pipelines that need deterministic, repeatable installations

4. **Asset Reset/Recovery**
   - Restoring customized Kibana assets to package defaults
   - Recovering from corrupted or misconfigured package assets
   - Fixing inconsistencies between package specification and installed assets

5. **Stuck Installation Recovery**
   - Package stuck in `installing` state from a previous failed operation
   - Need to bypass concurrency locks to restart installation

### When to Use `force = false`

#### Recommended Use Cases

1. **Production with Customized Assets**
   - Dashboards or visualizations have been manually customized
   - Customizations need to be preserved across package updates
   - Users have invested time in tuning visualizations for their specific use cases

2. **Environments with Accumulated Transform Data**
   - Transforms have been running and accumulating valuable data
   - Transform destination indices contain historical data that shouldn't be lost
   - Data continuity is critical for analytics or reporting

3. **Normal Package Updates**
   - Upgrading to a new package version (Fleet handles this without `force`)
   - Package is already at the desired version and functioning correctly
   - No need to trigger destructive re-installation

4. **Cautious Production Changes**
   - Production environments where stability is paramount
   - Changes should be incremental and non-destructive
   - Rollback capability is important

### Technical Details of `force=true`

1. **Kibana Assets:** `deleteKibanaSavedObjectsAssets()` removes all saved objects associated with the package
2. **Transforms:** Matching transforms are stopped, deleted, and their destination indices are removed
3. **Pipelines & Templates:** Overwritten via Elasticsearch PUT operations (no explicit deletion needed)
4. **Recreation:** All assets are reinstalled from the package archive after cleanup

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_elasticstack"></a> [elasticstack](#requirement\_elasticstack) | >= 0.11.0 |
| <a name="requirement_restapi"></a> [restapi](#requirement\_restapi) | >= 1.18.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_elasticstack"></a> [elasticstack](#provider\_elasticstack) | >= 0.11.0 |
| <a name="provider_restapi"></a> [restapi](#provider\_restapi) | >= 1.18.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [elasticstack_fleet_integration.assets](https://registry.terraform.io/providers/elastic/elasticstack/latest/docs/resources/fleet_integration) | resource |
| [restapi_object.package_policy](https://registry.terraform.io/providers/Mastercard/restapi/latest/docs/resources/object) | resource |
| [elasticstack_fleet_integration.latest](https://registry.terraform.io/providers/elastic/elasticstack/latest/docs/data-sources/fleet_integration) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_agent_policy_id"></a> [agent\_policy\_id](#input\_agent\_policy\_id) | ID of the agent policy to add the package policy to. | `string` | n/a | yes |
| <a name="input_all_data_streams"></a> [all\_data\_streams](#input\_all\_data\_streams) | List of all data streams associated to the input type in the policy template. This is necessary to disable all data streams except the one being used. | `list(string)` | `[]` | no |
| <a name="input_all_policy_template_inputs"></a> [all\_policy\_template\_inputs](#input\_all\_policy\_template\_inputs) | List of all policy template and input type combos. This is necessary to disable all inputs except the one being used. | `list(string)` | `[]` | no |
| <a name="input_data_stream"></a> [data\_stream](#input\_data\_stream) | Name of the data\_stream within the integration (e.g. "log"). | `any` | n/a | yes |
| <a name="input_data_stream_variables_json"></a> [data\_stream\_variables\_json](#input\_data\_stream\_variables\_json) | JSON encoded data stream specific variables. | `string` | `"{}"` | no |
| <a name="input_description"></a> [description](#input\_description) | Description to apply to the package policy. | `string` | `""` | no |
| <a name="input_force"></a> [force](#input\_force) | Force reinstallation of the package even if already installed. When true, bypasses "already installed"<br>checks and triggers complete re-installation. This deletes and recreates Kibana assets (dashboards,<br>visualizations), removes transforms and their destination indices, and overwrites ingest pipelines and<br>templates. Use with caution. See README for complete details and use cases. | `bool` | `false` | no |
| <a name="input_input_type"></a> [input\_type](#input\_input\_type) | Input type. | `string` | n/a | yes |
| <a name="input_input_variables_json"></a> [input\_variables\_json](#input\_input\_variables\_json) | JSON encoded input level variables. | `string` | `"{}"` | no |
| <a name="input_namespace"></a> [namespace](#input\_namespace) | Namespace to apply to data streams. | `string` | `"default"` | no |
| <a name="input_package_name"></a> [package\_name](#input\_package\_name) | Name of the Fleet package to install (e.g. ti\_recordedfuture). | `string` | n/a | yes |
| <a name="input_package_policy_name"></a> [package\_policy\_name](#input\_package\_policy\_name) | Name of the package policy. Defaults to "{package\_name}-{namespace}." | `any` | `null` | no |
| <a name="input_package_variables_json"></a> [package\_variables\_json](#input\_package\_variables\_json) | JSON encoded package level variables that are shared by each stream. | `string` | `"{}"` | no |
| <a name="input_package_version"></a> [package\_version](#input\_package\_version) | Package version. Use "latest" to resolve the latest available version from the package registry at plan time. | `string` | n/a | yes |
| <a name="input_policy_template"></a> [policy\_template](#input\_policy\_template) | Name of the policy template (see the integration's manifest.yml). | `string` | n/a | yes |
| <a name="input_prerelease"></a> [prerelease](#input\_prerelease) | Include prerelease package versions when resolving package\_version = "latest".<br>Required for packages whose versions are below 1.0.0 (semver treats 0.x as<br>prerelease). Has no effect when package\_version is pinned. | `bool` | `false` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | Package policy ID |
<!-- END_TF_DOCS -->