## fleet-terraform-generator generate

List all Fleet packages, policy templates, data streams, and inputs.

```
fleet-terraform-generator generate [flags]
```

### Options

```
  -d, --data-stream string       Data stream name.
  -h, --help                     help for generate
  -i, --input string             Input name.
      --out string               Output path. It creates a new sub-directory named based on the package, policy template, data stream, and input.
  -p, --package string           Package name.
      --packages-dir string      Directory containing Fleet packages.
  -t, --policy-template string   Policy template name.
```

### SEE ALSO

* [fleet-terraform-generator](fleet-terraform-generator.md)	 - Tool for generating Terraform modules for managing Fleet integrations.
* [fleet-terraform-generator generate batch](fleet-terraform-generator_generate_batch.md)	 - Generate a batch of Terraform modules as specified by glob or config file.


