# Gopherpen

This is a template project that will let you easily get started with GopherJS. It includes some simple HTML, CSS, and Go code for the frontend. Make some changes, and refresh in browser to see results.

Once you're done making changes, you can create a static production binary that has all assets built in, and can be deployed to any server.

## Installation

Run this to get gopherpen and all dependencies, both for development and production modes.

```
go get -u github.com/gopherjs/gopherpen
go get -u -tags=dev github.com/gopherjs/gopherpen
```

## Building

### Development Build

Accesses assets from disk directly.

```
go build -tags=dev
```

### Production Build

All assets are statically embedded in the binary, so it can run standalone in any folder.

```
go generate
go build
```

### `go generate` Dependencies

To run `go generate`, you'll also need:

```
go get -u github.com/shurcooL/vfsgen
```

License
-------

- [MIT License](http://opensource.org/licenses/mit-license.php)
