# Flux-web

> If you don't know what *GitOps* is then I highly encourage you to read [Weavework blog](https://www.weave.works/technologies/gitops/) and [Flux](https://github.com/fluxcd/flux) for the implemention.

Flux-web is used for manuall actions against flux api:
* List current workload
* Workload's version promotion or rollback
* Dynamic namespaces overview

<div style='position:relative; padding-bottom:calc(62.40% + 44px)'><iframe src='https://gfycat.com/ifr/EquatorialAcclaimedJaeger' frameborder='0' scrolling='no' width='100%' height='100%' style='position:absolute;top:0;left:0;' allowfullscreen></iframe></div>

## Usage

```sh
$ git url <path_to_file> [remote]
```

## Example

```sh
$ git clone git@github.com:maorfr/git-url.git && cd git-url
Cloning into 'git-url'...

$ git url README.md
https://github.com/maorfr/git-url/blob/master/README.md
```

# Install

1. Clone this repository.
2. Copy `git-url` to your $PATH.
3. Add git alias - `git config --global alias.url '!git-url'`.
