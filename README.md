# Gopherpen

Gopherpen is a project template that will let you easily get started with GopherJS
for building a web app. It includes some simple HTML, CSS, and Go code for the frontend.
Make some changes, and refresh in browser to see results. When there are errors in your
frontend Go code, they will show up in the browser console.

Once you're done making changes, you can easily create a fully self-contained static
production binary.

Installation
------------

Get the source code for gopherpen and all dependencies, both for production and development modes:

```bash
go get -u -d github.com/gopherjs/gopherpen/...
go get -u -d -tags=dev github.com/gopherjs/gopherpen/...
go get -u -d -tags=generate github.com/gopherjs/gopherpen/...
```

Building
--------

### Development Build

Accesses assets from disk directly:

```bash
go build -tags=dev
```

### Production Build

All assets are statically embedded in the binary, so it can run standalone in any folder:

```bash
go generate
go build
```

License
-------

-	[MIT License](http://opensource.org/licenses/mit-license.php)
