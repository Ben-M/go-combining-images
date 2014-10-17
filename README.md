Go Pre-Hackathon Image Example
==============================

An example of downloading and combining images in Go.

## What does it do?

* Downloads two images of unknown filetype and size.
* Resizes them to the same height.
* Combines them into a single image.
* Saves them as a PNG.

## What doesn't it do?

* Combine more than two images.
* Error handling.
* Concurrency.

## Getting Started

After installing Go and setting up your [GOPATH](http://golang.org/doc/code.html#GOPATH), clone this repository and install the dependencies:

```
$ go get github.com/nfnt/resize
```

Now you are ready to run the example with `go run example.go`

## Our dependencies

* [net/http](http://golang.org/pkg/net/http/) to download images.
* [resize](https://github.com/nfnt/resize) for resizing.
* [image/draw](http://golang.org/pkg/image/draw/) for combining images. You can learn more about how that works [here](http://blog.golang.org/go-imagedraw-package).
* [image](http://golang.org/pkg/image/), [image/jpeg](http://golang.org/pkg/image/jpeg/), [image/png](http://golang.org/pkg/image/png/) and [image/gif](http://golang.org/pkg/image/gif/) for image import and export.