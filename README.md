# Hugo Image Optimizer

> Quick and dirty image optimizer for Hugo based websites
>
> Removes all unused images, converts the rest (png, jpg and gifs) to a compressed .webp format and updates all references in your posts.

## Use Case

I often screw around with my blog, adding and changing images without much attention to size, weight or format.
As a result, I end up with a bunch of unused images lying around and a less-than ideal performance.

This takes care of these issues without me having to worry about it.

## Dependencies

- `go`
- [`cwebp`](https://developers.google.com/speed/webp/docs/cwebp)
- [`gif2webp`](https://developers.google.com/speed/webp/docs/gif2webp) (if GIFs need to be handled)

## Install

Either download the [latest binary](https://github.com/EricDriussi/hugo-image-optimizer/releases) and add it to your `$PATH` or install it directly using `go`:

```sh
go install github.com/EricDriussi/hugo-image-optimizer@latest
```

## Config

A `optimizer.ini` config file is expected and should look like this:

```ini
[dirs]
posts = content/posts/
images = static/images/
images_exclude = whoami donation

[compression]
quality = 50
```

This tells the optimizer where your posts and images are located (can be in subdirectories), what subdirectories to ignore when optimizing images, and how much compression to apply (0-100, 100 for max compression).

Add one to your website repo or in `~/.config/`. Specify a different path with `--config`.

If no config file is provided, it will default to the values above.

## Run

Either run it from your website directory or use the `--website-path` flag to specify the desired path.

Run `optimize --help` to check out the available commands and flags.

You can run them all at once with `optimize` or separately with `optimize [CMD]` for more fine-grained control.
