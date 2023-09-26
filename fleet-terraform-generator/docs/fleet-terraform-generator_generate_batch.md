## fleet-terraform-generator generate batch

Generate a batch of Terraform modules as specified by glob or config file.

```
fleet-terraform-generator generate batch [flags]
```

### Options

```
      --continue-on-error   Continue on errors related to reading package specifications.
  -h, --help                help for batch
```

### Options inherited from parent commands

```
      --ignore-var-shadow     Ignore variable shadowing errors in Fleet packages.
      --out string            Output path. It creates a new sub-directory named based on the package type, package name, policy template, data stream, and input.
      --packages-dir string   Directory containing Fleet packages.
```

### SEE ALSO

* [fleet-terraform-generator generate](fleet-terraform-generator_generate.md)	 - Generate a Terraform module for managing a Fleet package policy.


