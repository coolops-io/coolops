# CoolOps.io CLI

## Description

`coolops` is the command line interface to help you notify CoolOps about new builds to be deployed. It provides an user friendly interface for you to define the parameters to be injected in your deployment script and also metadata to be sent with the Slack message.

## Usage

```sh
coolops build:notify [build_name] -t [project_access_token] -p [parameter_name]=[parameter_value] -p [parameter_name]=[parameter_value] -m [metadata_name]=[metadata_value]
```

## Install

```sh
curl -L https://github.com/coolopsio/coolops/archive/master.tar.gz | tar xvz
cd coolops-* && make install
```

## Contribution

1. Fork ([https://github.com/coolopsio/coolops/fork](https://github.com/coolopsio/coolops/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run `gofmt -s`
1. Create a new Pull Request

