# MaxMind OpenAPI specifications

This repository holds [OpenAPI 3.1](https://spec.openapis.org/oas/v3.1.2)
descriptions of the MaxMind public web services.

| Product                        | Bundled spec                               | Documentation                                                       |
| ------------------------------ | ------------------------------------------ | ------------------------------------------------------------------- |
| GeoIP and GeoLite web services | [`bundled/geoip.yaml`](bundled/geoip.yaml) | [dev.maxmind.com](https://dev.maxmind.com/geoip/docs/web-services/) |

Each file in `bundled/` is self-contained. Use it with API tools and code
generators. The files in `specs/`, `components/`, and `examples/` are the
sources. Do not edit `bundled/` by hand.

The [developer documentation](https://dev.maxmind.com/) is the reference for
behavior that OpenAPI cannot describe. MaxMind also publishes official
[client libraries](https://dev.maxmind.com/geoip/docs/web-services/#official-client-apis).

## Compatibility

MaxMind can add response fields, error codes, and enum values to response fields
without a new API version. Your client must ignore fields that it does not know.

## Versioning

The `info.version` of each spec follows
[Semantic Versioning](https://semver.org/) and tracks changes to the spec. The
API version, for example `v2.1`, is in the server URL.

## Go

The `github.com/maxmind/openapi` Go package embeds the bundled specs:

```go
spec, err := openapi.Bundled.ReadFile("bundled/geoip.yaml")
```

## Development

The tools are pinned with [mise](https://mise.jdx.dev/).

```sh
mise install
pnpm install
pnpm run bundle      # regenerate bundled/
precious lint --all  # lint the specs and format the files
```

## License

This software is Copyright (c) 2026 by MaxMind, Inc.

This is free software, licensed under the
[Apache License, Version 2.0](LICENSE-APACHE) or the [MIT License](LICENSE-MIT),
at your option.
