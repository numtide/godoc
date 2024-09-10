# Godoc

<!-- prettier-ignore -->
> [!NOTE]
> **Status: alpha**

`godoc` is a custom documentation generator for [Go] projects. Currently focused on [Markdown], it may expand into other areas such as JSON schemas.

<!-- prettier-ignore -->
> [!WARNING]
Heavily biased towards [Numtide] use cases and projects.

## Quick Start

```console
❯ nix run .# -- --help
Custom Go doc generation

Usage:
  godoc [source directory] [flags]

Flags:
  -c, --clean        clean output directory before writing
  -h, --help         help for godoc
  -o, --out string   output directory

```

## Contributing

Contributions are always welcome!

## License

This software is provided free under the [MIT] license.

---

This project is supported by [Numtide](https://numtide.com/).

![Numtide Logo](https://codahosted.io/docs/6FCIMTRM0p/blobs/bl-sgSunaXYWX/077f3f9d7d76d6a228a937afa0658292584dedb5b852a8ca370b6c61dabb7872b7f617e603f1793928dc5410c74b3e77af21a89e435fa71a681a868d21fd1f599dd10a647dd855e14043979f1df7956f67c3260c0442e24b34662307204b83ea34de929d)

We’re a team of independent freelancers that love open source.
We help our customers make their project lifecycles more efficient by:

-   Providing and supporting useful tools such as this one.
-   Building and deploying infrastructure, and offering dedicated DevOps support.
-   Building their in-house Nix skills, and integrating Nix with their workflows.
-   Developing additional features and tools.
-   Carrying out custom research and development.

[Contact us](https://numtide.com/contact) if you have a project in mind,
or if you need help with any of our supported tools, including this one.

We'd love to hear from you.

[MIT]: ./LICENSE
[Numtide]: https://numtide.com
[Markdown]: https://www.markdownguide.org/
[Go]: https://go.dev/
