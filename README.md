# flux-web
[![HitCount](http://hits.dwyl.io/idobry/flux-web.svg)](http://hits.dwyl.io/idobry/flux-web) [![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/dwyl/esta/issues) [![Go Report Card](https://goreportcard.com/badge/github.com/idobry/flux-web)](https://goreportcard.com/report/github.com/idobry/flux-web) [![Build Status](https://travis-ci.com/ybaruchel/flux-web.svg?branch=master)](https://travis-ci.com/ybaruchel/flux-web)

> If you don't know what *GitOps* is then I highly encourage you to read [Weavework blog](https://www.weave.works/technologies/gitops/) and [flux](https://github.com/fluxcd/flux) for the implemention.

flux-web is used for manuall actions against flux api:
* list current workloads
* dynamic namespaces overview
* workload's version promotion or rollback

<img src="flux-web-01.gif"/>

## Setup

Please see the [Weave Flux setup documentation](https://github.com/weaveworks/flux/blob/master/site/standalone/installing.md) for setting up Flux.

To use flux-web, you should deploy a seperate deployment, see `deploy/flux-web-deploy.yaml`.

Set the following environment variables in your chosen deployment:

- `FLUX_URL`: fluxd's endpoint
- `DEFAULT_NAMESPACE`: default namespace to be set as home page
- `NAMESPACES`: namespaces list to show in the navigation bar
- `READ_ONLY`: restric flux-web to read-only mode

And then apply the configuration:

```
kubectl apply -f deploy/flux-web-deploy.yaml
```

## Reasoning

When using `flux` your pipelines, or, to be more precise your Continues Deployments are fully automated and that's great.
But, sometimes we would like to have more stable environments with fewer deployments and with more control, for example, production.
That's mean we need to do some manual actions against `flux` which have his own CLI tool `fluxctl`.

flux-web is intended to be the UI approach to this problem. With flux-web you can view at our workloads per namespace and the available versions for each and with a single click we can promote a workload or to perform a rollback.

## Continued Development

Basically a roadmap.

### Coming soon

- filter workloads - Done
- scroll over workload's versions
- features you'd like to see?

### Maybe in the future, if people want it

- user access and authentication

### Probably in the future

- select multiple workloads
- use socket

## Built With

* [beego](https://beego.me/) - Web framework
* [go](https://golang.org/) - Programing language
* [docker](https://www.docker.com/) - Packaged with docker


## Contributing

Code contributions are very welcome. If you are interested in helping make flux-web great then feel free!

## Authors

* **Ido Braunstain** - *Initial work*
