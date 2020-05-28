---
layout: tour

title: Raster
order: 6

---

The `canvas.Raster` is like an image but draws exactly one spot
for each pixel on the screen. This means that as a user interface
scales or the image resizes more pixels will be requested to fill
the space. To do this we use a `Generator` function as illustrated in
this example - it will be used to return the colour of each pixel.

The generator functions can be pixel based (as in this example where we 
generate a new random colour for each pixel) or based on full images.
Generating complete images (with `canvas.NewRaster()` is more efficient
but sometimes controlling pixels directly is more convenient.

If your pixel data is stored in an image you can load it through the
`NewRasterFromImage()` function which will load the image to display
pixel perfect on screen.
