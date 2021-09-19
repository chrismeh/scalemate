# scalemate

A tool for generating images for guitar/bass scales.

This app is running at [scalemate.cmeh.dev](https://scalemate.cmeh.dev).

## Usage (CLI)

```shell
$ make build
go build -o=./bin/scalemate-cli ./cmd/cli
go build -o=./bin/scalemate-web ./cmd/web

$ bin/scalemate-cli --help
Usage of bin/scalemate-cli:
  -file string
        Filename for saving the PNG (default "scale.png")
  -frets uint
        Number of frets on the neck (default 12)
  -scale string
        Scale you want to generate (default "A minor")
  -tuning string
        Guitar/bass tuning, notes separated by a whitespace (default "E A D G B E")

```

Example: Draw the A major scale in Drop-C tuning:
```shell
$ bin/scalemate-cli -tuning="C G C F A D" -scale="A major" -file="a-major-in-drop-c.png"
```

This will generate the following image:

![a-major-in-drop-c](https://user-images.githubusercontent.com/32984536/133892891-42cbd796-c6a3-4cb2-a08b-df0fa2f40cfc.png)

## Usage (Web)

```shell
$ make build
go build -o=./bin/scalemate-cli ./cmd/cli
go build -o=./bin/scalemate-web ./cmd/web

$ bin/scalemate-web --help
Usage of bin/scalemate-web:
  -addr string
        TCP address for the server to listen on (default ":8080")
```

Example: Start scalemate server at port 5000:
```shell
$ bin/scalemate-web -addr=":5000"
INFO    2021/09/19 17:15:02 starting application at port :5000
```

Example: Start scalemate server at port 5000, using Docker:
```shell
$ docker build -t scalemate:latest
$ docker run --rm -e ADDR=":5000" -p "5000:5000" scalemate:latest
INFO    2021/09/19 17:15:02 starting application at port :5000
```