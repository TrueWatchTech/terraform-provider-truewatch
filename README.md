# Terraform Provider: TrueWatch

The TrueWatch Provider provides resources to manage [TrueWatch Cloud](https://en.truewatch.com/) resources.

## Documentation, questions, and discussions

Official documentation on how to use this provider can be found on the [Terraform Registry](https://registry.terraform.io/providers/TrueWatchTech/truewatch/latest/docs).

The remainder of this document will focus on the development aspects of the provider.

The resource supports as follows:

| Resource | Description |
| --- | --- |
| `truewatch_alert_policy` | Alert routing policy, notification targets, aggregation, silence, escalation, permissions, and checker/security rule bindings. |
| `truewatch_alert_policy_notice_date` | Custom alert policy notification date windows. |
| `truewatch_blacklist` | Blacklist configuration. |
| `truewatch_dashboard` | Dashboard configuration. |
| `truewatch_membergroup` | Workspace member groups. |
| `truewatch_monitor` | Monitor/checker rules managed with structured Terraform fields. |
| `truewatch_monitor_json` | Monitor/checker rules managed with raw JSON payloads. |
| `truewatch_mute` | Alert, checker, tag, and custom mute rules. |
| `truewatch_notify_object` | Alert notification objects and notification permissions. |
| `truewatch_pipeline` | Pipeline configuration. |
| `truewatch_role` | Workspace roles. |

The data source supports as follows:

| Data Source | Description |
| --- | --- |
| `truewatch_alert_policy` | Look up an alert policy by UUID or exact name. |
| `truewatch_alert_policy_notice_date` | Look up a custom alert policy notification date by UUID or exact name. |
| `truewatch_members` | List workspace members. |
| `truewatch_monitor` | Look up a monitor/checker by UUID or exact name. |
| `truewatch_monitors` | List monitor/checker rules with filters. |
| `truewatch_mute` | Look up a mute rule by UUID or exact name. |
| `truewatch_notify_object` | Look up a notification object by UUID or exact name. |
| `truewatch_permissions` | List workspace permissions. |

The region supports as follows:

* [x] oregon
* [x] frankfurt
* [x] singapore
* [x] southafrica
* [x] jakarta

If there are more resources you need, create an [issue](https://github.com/TrueWatchTech/terraform-provider-truewatch/issues) for free.

## Compatibility

Compatibility table between this provider, the [Terraform Plugin Protocol](https://www.terraform.io/plugin/how-terraform-works#terraform-plugin-protocol)
version it implements, and Terraform:

| TrueWatch Provider | Terraform Plugin Protocol | Terraform |
|:---------------:|:-------------------------:|:---------:|
|    `>= 0.x`     |            `6`            | `>= 1.0`  |

Details can be found by querying the [Registry API](https://www.terraform.io/internals/provider-registry-protocol#list-available-versions)
that return all the details about which versions are currently available for a particular provider.
[Here](https://registry.terraform.io/v1/providers/TrueWatchTech/truewatch/versions) are the details.

## Requirements

* [Terraform](https://www.terraform.io/downloads)
* [Go](https://go.dev/doc/install) (1.26.2)

## Development

### Generating documentation

This provider uses [terraform-plugin-docs](https://github.com/hashicorp/terraform-plugin-docs/)
to generate documentation and store it in the `docs/` directory.
Once a release is cut, the Terraform Registry will download the documentation from `docs/`
and associate it with the release version. Read more about how this works on the
[official page](https://www.terraform.io/registry/providers/docs).

Use `make docs` to regenerate documentation and `make check-docs` to verify
that the generated documentation is up to date.

### Using a development build

If [running tests and acceptance tests](#testing) aren't enough, it's possible to set up a local terraform configuration
to use a development build of the provider. This can be achieved by leveraging the Terraform CLI
[configuration file development overrides](https://www.terraform.io/cli/config/config-file#development-overrides-for-provider-developers).

First, use `make install` to place a fresh development build of the provider in your
[`${GOBIN}`](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies)
(defaults to `${GOPATH}/bin` or `${HOME}/go/bin` if `${GOPATH}` is not set). Repeat
this every time you make changes to the provider locally.

Then, set up your environment following [these instructions](https://www.terraform.io/plugin/debugging#terraform-cli-development-overrides)
to make your local Terraform CLI use your local build.

### Testing

Run the unit test suite with:

```shell
make test
```

Run acceptance tests against a configured TrueWatch environment with:

```shell
make testacc
```

### Testing GitHub Actions

This project uses [GitHub Actions](https://docs.github.com/en/actions/automating-builds-and-tests) to realize its CI.

Sometimes it might be helpful to reproduce the behavior of those actions locally,
and for this, we use [act](https://github.com/nektos/act). Once installed, you can _simulate_ the actions executed
when opening a PR with:

```shell
# List of workflows for the 'pull_request' action
$ act -l pull_request

# Execute the workflows associated with the `pull_request' action 
$ act pull_request
```

## Releasing

The release process is automated via GitHub Actions, and it's defined in the Workflow
[release.yml](./.github/workflows/release.yml).

Each release is cut by pushing a [semantically versioned](https://semver.org/) tag to the default branch.

## License

[Mozilla Public License v2.0](./LICENSE)
