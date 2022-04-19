<p styles="font-size: 20rem" align="center">
    <img styles="margin: 0px" width="350px" src="https://i.giphy.com/media/bdaMVCP6bPIVWc0B3g/giphy.webp" />
</p>
<h1 align="center">
goapprove
</h1>
<h4 align="center">
Go approve a PR
</h4>
<p align="center">
use it carefully
</p>

## Installation

It's easy!

```sh
bash <(curl -sL https://raw.githubusercontent.com/amazingandyyy/goapprove/main/install.sh)
```

## Preparation

- [Generate a Github personal access token](https://github.com/settings/tokens/new?scopes=repo&description=goapprove-cli)
  - [repo] scrope
  - [no expiration]

## Usage

```sh
# get instructions
goapprove -help

goapprove -url https://github.com/amazingandyyy/go-approve/pull/1

goapprove -url https://github.com/amazingandyyy/go-approve/pull/1 -action comment -message "LGTM!!!!!
```

## Development

### Initialization

```sh
git clone https://github.com/amazingandyyy/goapprove.git
cd goapprove
./scripts/setup
```

### Trigger pre-commit hook manually

```sh
pre-commit run --all-files
```

---

![banner](assets/repo-banner.jpg)

## Contribution

- Try out this tool
- Star this repo
- Share with your friend!
- [Create issue](https://github.com/amazingandyyy/goapprove/issues/new)
- Create PRs

## Resources

- [regex101](http://regex101.com)
- wildcard

## LICENSE

[MIT](LICENSE)
