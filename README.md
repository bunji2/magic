# magic

Go implementation to pickup magic numbers

## Usage

```
C:\work>magic.exe
Usage: magic.exe file
```

This tool displays positions of magic numbers.

```
C:\work>magic.exe MrFusion.gpjb
0 gif
6943 png
9727 jpg
2791486 gif
2794240 png
2796217 jpg
5578481 gif
5580896 png
5583378 jpg
8366075 gif
8368830 png
8371932 jpg
```

## magic.json

```
{
    "bmp": [[66, 67]],
	"jpg": [[255, 216]],
	"png": [[137, 80, 78, 71]],
    "gif": [[71, 73, 70, 56, 57, 97], [71, 73, 70, 56, 55, 97]]
}
```