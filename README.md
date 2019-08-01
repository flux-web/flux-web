# Flux-web

> If you don't know what *GitOps* is then I highly encourage you to read [Weavework blog](https://www.weave.works/technologies/gitops/) and [Flux](https://github.com/fluxcd/flux) for the implemention.

Flux-web is used for manuall actions against flux api:
* List current workload
* Workload's version promotion or rollback
* Dynamic namespaces overview

<a href="https://youtu.be/FTHxTS-TV5U" target="_blank"><img src="http://img.youtube.com/vi/YOUTUBE_VIDEO_ID_HERE/0.jpg" alt="Flux-web Demo" width="240" height="180" border="10" /></a>

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
