# bfc

A brainfuck compiler written in Go. Right now only supports Linux on the x86-64 architechture.
More targets to come.

## Why?

Just for fun! Also to test my [libtcc go bindings](https://github.com/sorribas/tcc) and build
a proof of concept for them.

## Installing

```sh
curl -L https://github.com/sorribas/bfc/releases/download/v0.1.0/bfc-0.1.0-linux-x86-64.tar.gz | tar -xzf -
mv bfc /usr/bin/ # or another directory in your path
```

## Usage

Run `bfc` to see the usage.
