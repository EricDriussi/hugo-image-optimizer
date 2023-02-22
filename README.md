# Hugo Image Optimizer

⚠️ WIP!

> Quick and dirty image optimizer for Hugo based websites
>
> Removes all unused images, converts the rest (png, jpg and gifs) to a compressed .webp format and updates all references in your posts.

## Use Case

I often screw around with my blog, adding and changing images without much attention to size, weight or format.
As a result, I end up with a bunch of unused images lying around and a less-than ideal performance.

This takes care of these issues without me having to worry about any of it.

## Dependencies

- `go`
- [`cwebp`](https://developers.google.com/speed/webp/docs/cwebp)
- [`gif2webp`](https://developers.google.com/speed/webp/docs/gif2webp) (if GIFs need to be handled)

## Install

Either download the [latest binary](https://github.com/EricDriussi/hugo-image-optimizer/releases) and add it to your `$PATH` or clone the repo and install it yourself:

```sh
git clone git@github.com:EricDriussi/hugo-image-optimizer.git optimizer && cd optimizer && make install
```

## Config

Create a `optimizer.ini` file in your Hugo website with the following structure (or copy the one in this repo):

```ini
[dirs]
posts = content/posts/
images = static/images/
images_exclude = whoami donation
```

This tells the optimizer where your posts and images are located, as well as what subdirectories to ignore when optimizing images.

The script looks recursively through the directory tree, so both the posts and images can be in subdirectories.

## Run

`cd` into your website directory and run `optimize --help` to check out the available commands!

Either run them all at once with `optimize` or separately with `optimize [CMD]` for more fine-grained control.
