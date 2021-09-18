# scalemate

A simple CLI tool (as of now) for generating graphical representations of guitar/bass scales.

## Usage

```
$ make build            
go build -o=./bin/scalemate ./cmd/scalemate

$ bin/scalemate --help
Usage of bin/scalemate:
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
```bash
$ bin/scalemate -tuning="C G C F A D" -scale="A major" -file="a-major-in-drop-c.png"
```

This will generate the following image:

![a-major-in-drop-c](https://user-images.githubusercontent.com/32984536/133892891-42cbd796-c6a3-4cb2-a08b-df0fa2f40cfc.png)
