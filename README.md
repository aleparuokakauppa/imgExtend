# ImgExtend
Extend the resolution of your images!
This program works by extending the original image in the direction specified, and filling the new space with the color at the center of the edge, where the image is extending.
Very useful for making wallpapers for machines, with differing resolutions of the original image.
Currently only works with PNG images.

# Usage
Run:
```bash
go run main.go [inputImage.png] [new resolution eg. 1500] [horizontal offset to fillcolor (optional)]
```

It might be useful to run `file` (or similar) on the inputImage to get its resolution, to have context on what the new resolution should be.

# Roadmap
- The program extends only upwards, in the future images could be extended in any direction.
- Maybe non PNG images

# Example images
Img source: Ripple Effects KAI DAHMS, UNSPLASH

Original image:
![Kai_Dahms_Ripples.png](https://github.com/aleparuokakauppa/imgExtend/blob/main/images/Kai_Dahms_Ripples.png?raw=true)

Extended with `./main images/Kai_Dahms_Ripples.png 1500`:
![Kai_Dahms_Ripples_extended.png](https://github.com/aleparuokakauppa/imgExtend/blob/main/images/Kai_Dahms_Ripples_extended.png?raw=true)
