# flux-web

> If you don't know what *GitOps* is then I highly encourage you to read [Weavework blog](https://www.weave.works/technologies/gitops/) and [flux](https://github.com/fluxcd/flux) for the implemention.

flux-web is used for manuall actions against flux api:
* list current workloads
* dynamic namespaces overview
* workload's version promotion or rollback

<img src="flux-web-01.gif"/>

## Setup

Please see the [Weave Flux setup documentation](https://github.com/weaveworks/flux/blob/master/site/standalone/installing.md) for setting up Flux.

To use flux-web, you should deploy a seperate deploymenty, see `deploy/flux-web-deploy.yaml`.

Set the following environment variables in your chosen deployment:

- `FLUX_URL`: fluxd's endpoint
- `DEFAULT_NAMESPACE`: default namespace to show in the home page
- `NAMESPACES`: namespaces list to show in the navigation bar

And then apply the configuration:

```
kubectl apply -f deploy/flux-web-deploy.yaml
```

## Reasoning

When using `flux` your pipelines, or, to be more precise your Continues Deployments are fully automated and that's great.
But, we would like to have some more stable environments with fewer deployments and more controls, for example, production.
That's mean we need to do some manual actions against flux which have his own CLI tool `fluxctl`.

flux-web is intended to be the UI approach to this problem. With flux-web you can see in the browser your workloads per namespace and the available versions for each and with a single click to promote a workload or to perform a rollback.

## Continued Development

Basically a roadmap.

### Coming soon

- filter workloads - Done
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
