# goapprove

Go approve a PR.

# Installation

```sh
bash <(curl -sL https://raw.githubusercontent.com/amazingandyyy/goapprove/main/install.sh)
```

## Preparation

- [Generate a Github personal access token](https://github.com/settings/tokens/new?scopes=repo&description=goapprove-cli)
  - [repo] scrope
  - [no expiration]

## Usage

```
goapprove -help
goapprove -url https://github.com/amazingandyyy/go-approve/pull/1
goapprove -url https://github.com/amazingandyyy/go-approve/pull/1 -action comment -message "LGTM!"
```

## LICENSE

MIT
