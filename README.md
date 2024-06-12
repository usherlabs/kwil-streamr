# Kwil-Streamr

Kwil-Streamr is a collection of Kwil extensions that allows databases to natively sync with Streamr streams. Using these extensions, network operators can configure a Kwil network's validators to sync Streamr data into a database, without the need for a centralized oracle.

## Sponsors

A huge thank you to the [PowerPod](<https://www.powerpod.pro/>) team for sponsoring this extension. PowerPod is a revolutionary DePIN network building decentralized electric vehicle charging infrastructure. They are currently shipping their first product, Pulse, which can be found [here](<https://pulse.powerpod.pro/>), and are also working on several non-ev devices as well. [Follow their journey to monetize every electron on X](<https://x.com/PowerPod_People>).

![PowerPod Logo](./assets/powerpod_logo.png)

## Getting Started

Choose one of the links below to get started:

- [Tutorial](./docs/tutorial.md)
- [Extension Documentation](./docs/extensions.md)
- [Required Streamr Configuration](./docs/streamr.md)

## Build

To build from source, ensure you have Go 1.22+ installed, and run:

```shell
make build
```

## In Your Own Kwil Network

To use the Streamr extension in your own Kwil network, import the extensions found in [`extensions/`](./extensions/) into your own Kwil binary.