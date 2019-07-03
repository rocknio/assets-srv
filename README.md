# Asset Service

This is the Asset service

Generated with

```
micro new github.com/rocknio/micro-assets/assets-srv --namespace=rz.assets --alias=asset --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: rz.assets.srv.asset
- Type: srv
- Alias: asset

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./asset-srv
```

Build a docker image
```
make docker
```