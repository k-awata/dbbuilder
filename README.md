# dbbuilder

dbbuilder outputs a macro for Aveva Administration to set up project database using options in YAML format.

## Installation

If you're using Go:

```bash
go install github.com/k-awata/dbbuilder@latest
```

Otherwise you can download a binary from [Releases](https://github.com/k-awata/dbbuilder/releases).

## Usage

- Output a sample YAML

```bash
dbbuilder -s > sample.yaml
```

- Output a macro to run in E3D

```bash
dbbuilder sample.yaml > export.mac
```

## License

[MIT License](LICENSE)
