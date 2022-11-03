# Tabify

Converts spaces to tabs in all files within a folder, recursively.

The target file extensions are specified in the [main.go](main.go) file, and can be modified to suit your needs.

## Usage

Build the executable:

    go build -ldflags "-s -w"

Run in the given folder:

    ./tabify ~/foo/my-project

## License

Licensed under [MIT license](https://opensource.org/licenses/MIT), see [LICENSE.md](LICENSE.md) for details.
