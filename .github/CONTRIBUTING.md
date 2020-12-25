# Contributing

Hi! Thank you for considering contributing to Fay. You'll
find below useful information about how to contribute to the Fay project.

## Contributing code

### Install from sources

1. Install and run the latest version of Docker
2. Verify your Go version (>= 1.15)
3. Fork this repository
4. Clone it outside of your `GOPATH` (we're using Go modules)

### Working with git

1. Create your feature branch (`git checkout -b my-new-feature`)
2. Commit your changes (`git commit -am 'Add some feature'`)
3. Push to the branch (`git push origin my-new-feature`)
4. Create a new pull request

### Building

```bash
./scripts/build.sh
```

### Testing

```bash
./scripts/test.sh
./scripts/coverage.sh
```

### Update swagger

```bash
./scripts/swagger.sh
```

## Development

1. Install `npm install -g serve`.
2. Serve example template `serve -l 3001 ./example`.
3. Run server `go run cmd/fay/main.go`.
4. Open swagger [endpoint](http://localhost:3000/swagger/index.html).

## Reporting bugs and feature request

Your issue or feature request may already be reported!
Please search on the [issue tracker](../../../issues) before creating one.

If you do not find any relevant issue or feature request, feel free to
add a new one!

## Additional resources

TDB.