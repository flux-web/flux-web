# Flux-Web chart

This chart bootstraps a [flux-web](https://github.com/idobry/flux-web) deployment. 

## Configuration

The following table lists the configurable parameters of the flux-web chart and their default values.

| Parameter                  | Description                                | Default  |
| -------------------------- | ------------------------------------------ | ----- |
| `replicaCount`             | Number of pod replicas	                  | `1`     |
| `applicationPort`          | The port your app runs on in its container | `8080`     |
| `image.repository`         | Image Repository.                          | `idobry/flux-web`     |
| `image.tag`                | Image tag                                  | `latest`     |
| `image.pullPolicy`         | Image Pull policy                          | `IfNotPresent`     |
| `imagePullSecrets`         | Specify Image pull secrets   | `nil`     |
| `nameOverride`             | String to partially override flux-web.fullname template with a string (will prepend the release name)   | `nil`     |
| `fullnameOverride`         | String to fully override flux-web.fullname template with a string  | `nil`     |
| `readOnly`                 | String property  to restrict flux-web to read-only mode   | `false`     |
| `namepaces`                | `;` seperated list of namespaces (to show in the navigation bar ) | `nil`     |
| `fluxUrl`                  | fluxd's endpoint     | `http://flux:3030`     |
| `defaultNamespace`         | default namespace to be set as home page | `nil` |
| `environment`              |  A map containing all environment values you wish to set. <br> **Note**: environment variables (the key in KEY: value) must be uppercase and only contain letters,  "_", or numbers and value can be templated | `nil`|
| `service.type`             | K8s service type | `ClusterIP`|
| `service.port`             | K8s service port | `80`|
| `ingress.enabled`          | Optionally enable ingress | `false`|
| `ingress.annotations`      | yaml property. Annotations for ingress resource. | `{}`|
| `ingress.hosts`            | List of ingress hosts | `[]`|
| `resources`                | yaml property to set k8s pod resources like memory, cpu limits  | `nil`|
| `nodeSelector`             | yaml property to set k8s node selector settings|  `nil` |
| `tolerations`              | yaml property to set k8s node toleration settings |  `nil` |
| `affinity`                 | yaml property to set k8s node affinity settings |  `nil` |