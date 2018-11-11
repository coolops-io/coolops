# coolops.io CLI

## Description

`coolops` is the command line interface to help you notify coolops about new builds to be deployed. It provides an user friendly interface for you to define the parameters to be injected in your deployment script and also metadata to be sent with the Slack message.

## Usage

```sh
coolops build:new [build_name] \
        -t [project_access_token] \
        -p [parameter_name]=[parameter_value] \
        -p [parameter_name]=[parameter_value] \
        -m [metadata_name]=[metadata_value]
```

## Install

```sh
curl -L https://github.com/coolops-io/coolops/releases/download/v0.2.0/install.sh | sh
```

## Contribution

1. Fork ([https://github.com/coolops-io/coolops/fork](https://github.com/coolops-io/coolops/fork))
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run `gofmt -s`
6. Create a new Pull Request

