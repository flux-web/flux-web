# Flux-web

> If you don't know what *GitOps* is then I highly encourage you to read [Weavework blog](https://www.weave.works/technologies/gitops/) and [Flux](https://github.com/fluxcd/flux) for the implemention.

Flux-web is used for manuall actions against flux api:
* List current workloads
* Dynamic namespaces overview
* Workload's version promotion or rollback

<img src="flux-web-01.gif"/>

# Setup

Please see the [Weave Flux setup documentation](https://github.com/weaveworks/flux/blob/master/site/standalone/installing.md) for setting up Flux.

To use Flux-web, you should deploy flux-web a seperate deploymenty, see `deploy/flux-web-deploy.yaml`.

Set the following environment variables in your chosen deployment:

* `FLUX_URL`: endpoint for the fluxd.
* `DEFAULT_NAMESPACE`: default namespace to show in the home page.
* `NAMESPACES`: namespaces list to show in the navigation bar.

And then apply the configuration:

```
kubectl apply -f deploy/flux-web-deploy.yaml
```
