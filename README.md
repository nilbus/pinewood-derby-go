PinewoodDerby
=============

A web-based race manager and status board application for the Cub Scout [pinewood derby](http://en.wikipedia.org/wiki/Pinewood_derby)

This is a rewrite of https://github.com/nilbus/pinewood-derby, designed to be fast enough to run well on a Raspberry Pi. It is under active development.

![screenshot](http://cl.ly/image/1L3b3g0o0R0F/Screen%20shot%202013-02-03%20at%209.18.25%20PM.png)

Developing
----------

Ensure `$GOPATH` is set. See https://golang.org/doc/code.html

Get the package:

```bash
$ go get github.com/nilbus/pinewood-derby-go
```

Use gin to run the server and reload code live as you make changes:

```bash
$ cd $GOPATH/src/github.com/nilbus/pinewood-derby
$ gin
[gin] listening on port 3000
```

License
-------

This pinewood-derby application is Copyright 2014 Edward Anderson,
and is distributed under the GNU Affero General Public License (see LICENSE).
This license requires that you provide the source code to users of this application.
