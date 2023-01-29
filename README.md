# Hugo Image Optimizer

⚠️ WIP!

> Quick and dirty image optimizer for Hugo based websites
>
> Removes all unused images, converts the rest (png, jpg and gifs) to a compressed .webp format and updates all references in your posts.

## Dependencies

- `go`
- `gif2webp`
- `cwebp`

## Install

Either download the [latest binary](https://github.com/EricDriussi/hugo-image-optimizer/releases) and `go env -w GOBIN=$HOME/.local/bin && go install` or clone the repo and build it yourself:

```sh
git clone git@github.com:EricDriussi/hugo-image-optimizer.git optimizer && cd optimizer && go build && go env -w GOBIN=$HOME/.local/bin && go install
```

## Config

Create a `optimizer.ini` file in your Hugo website with the following structure (or copy the one in this repo):

```ini
[dirs]
posts = content/posts/
images = static/images/
images_exclude = whoami donation
```

<!--TODO. add compression config-->

This tells the optimizer where your posts and images are located, as well as what subdirectories to ignore when optimizing images.

The script looks recursively through the directory tree, so both the posts and images can be in subdirectories.

## Run

`cd` into your website directory and run `optimize --help` to check out the available commands!

Either run them all at once with `optimize` or separately with `optimize [CMD]` for more fine-grained control.
