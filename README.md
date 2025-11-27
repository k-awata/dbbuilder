# dbbuilder

dbbuilder outputs a macro for Aveva Administration to set up project database using options in JSON format.

## Installation

If you're using Go:

```bash
go install github.com/k-awata/dbbuilder@latest
```

Otherwise you can download a binary from [Releases](https://github.com/k-awata/dbbuilder/releases).

## Usage

- Output a sample JSON

```bash
dbbuilder -s > sample.json
```

- Output a macro to run in E3D

```bash
dbbuilder sample.json > export.mac
```

## License

[MIT License](LICENSE)
