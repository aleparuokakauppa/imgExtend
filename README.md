# ImgExtend
Extend the resolution of your images!
This program works by extending the original image in the direction specified, and filling the new space with the color at the center of the edge, where the image is extending.
Very useful for making wallpapers for machines, with differing resolutions of the original image.
Currently only works with PNG images.

# Usage
Build with:
`go build`
Run:
`./main [inputImage.png] [new resolution eg. 1500]`
It might be useful to run `file` on the inputImage to get its resolution, to have context on what the new resolution should be.

# Roadmap
- The program extends only upwards, in the future images could be extended in any direction.
- Maybe non PNG images
