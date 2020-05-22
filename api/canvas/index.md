---
layout: page
tags: [api]
title: Fyne API canvas
---

# canvas
--
    import "fyne.io/fyne/canvas"

Package canvas contains all of the primitive CanvasObjects that make up a Fyne
### GUI

## Usage

#### func  Refresh

```go
func Refresh(obj fyne.CanvasObject)
```
Refresh instructs the containing canvas to refresh the specified obj.

#### type Circle

```go
type Circle struct {
	Position1 fyne.Position // The current top-left position of the Circle
	Position2 fyne.Position // The current bottomright position of the Circle
	Hidden    bool          // Is this circle currently hidden

	FillColor   color.Color // The circle fill colour
	StrokeColor color.Color // The circle stroke colour
	StrokeWidth float32     // The stroke width of the circle
}
```

Circle describes a coloured circle primitive in a Fyne canvas

#### func  NewCircle

```go
func NewCircle(color color.Color) *Circle
```
NewCircle returns a new Circle instance

#### func (*Circle) Hide

```go
func (l *Circle) Hide()
```
Hide will set this circle to not be visible

#### func (*Circle) MinSize

```go
func (l *Circle) MinSize() fyne.Size
```
MinSize for a Circle simply returns Size{1, 1} as there is no explicit content

#### func (*Circle) Move

```go
func (l *Circle) Move(pos fyne.Position)
```
Move the circle object to a new position, relative to its parent / canvas

#### func (*Circle) Position

```go
func (l *Circle) Position() fyne.Position
```
Position gets the current top-left position of this circle object, relative to
its parent / canvas

#### func (*Circle) Refresh

```go
func (l *Circle) Refresh()
```
Refresh causes this object to be redrawn in it's current state

#### func (*Circle) Resize

```go
func (l *Circle) Resize(size fyne.Size)
```
Resize sets a new bottom-right position for the circle object

#### func (*Circle) Show

```go
func (l *Circle) Show()
```
Show will set this circle to be visible

#### func (*Circle) Size

```go
func (l *Circle) Size() fyne.Size
```
Size returns the current size of bounding box for this circle object

#### func (*Circle) Visible

```go
func (l *Circle) Visible() bool
```
Visible returns true if this circle is visible, false otherwise

#### type Image

```go
type Image struct {

	// one of the following sources will provide our image data
	File     string        // Load the image from a file
	Resource fyne.Resource // Load the image from an in-memory resource
	Image    image.Image   // Specify a loaded image to use in this canvas object

	Translucency float64   // Set a translucency value > 0.0 to fade the image
	FillMode     ImageFill // Specify how the image should scale to fill or fit
}
```

Image describes a drawable image area that can render in a Fyne canvas The image
may be a vector or a bitmap representation and it will fill the area. The fill
mode can be changed by setting FillMode to a different ImageFill.

#### func  NewImageFromFile

```go
func NewImageFromFile(file string) *Image
```
NewImageFromFile creates a new image from a local file. Images returned from
this method will scale to fit the canvas object. The method for scaling can be
set using the Fill field.

#### func  NewImageFromImage

```go
func NewImageFromImage(img image.Image) *Image
```
NewImageFromImage returns a new Image instance that is rendered from the Go
image.Image passed in. Images returned from this method will scale to fit the
canvas object. The method for scaling can be set using the Fill field.

#### func  NewImageFromResource

```go
func NewImageFromResource(res fyne.Resource) *Image
```
NewImageFromResource creates a new image by loading the specified resource.
Images returned from this method will scale to fit the canvas object. The method
for scaling can be set using the Fill field.

#### func (*Image) Alpha

```go
func (i *Image) Alpha() float64
```
Alpha is a convenience function that returns the alpha value for an image based
on its Translucency value. The result is 1.0 - Translucency.

#### func (*Image) Hide

```go
func (r *Image) Hide()
```
Hide will set this object to not be visible

#### func (*Image) MinSize

```go
func (r *Image) MinSize() fyne.Size
```
MinSize returns the specified minimum size, if set, or {1, 1} otherwise

#### func (*Image) Move

```go
func (r *Image) Move(pos fyne.Position)
```
Move the rectangle object to a new position, relative to its parent / canvas

#### func (*Image) Position

```go
func (r *Image) Position() fyne.Position
```
CurrentPosition gets the current position of this rectangle object, relative to
its parent / canvas

#### func (*Image) Refresh

```go
func (i *Image) Refresh()
```
Refresh causes this object to be redrawn in it's current state

#### func (*Image) Resize

```go
func (r *Image) Resize(size fyne.Size)
```
Resize sets a new size for the rectangle object

#### func (*Image) SetMinSize

```go
func (r *Image) SetMinSize(size fyne.Size)
```
SetMinSize specifies the smallest size this object should be

#### func (*Image) Show

```go
func (r *Image) Show()
```
Show will set this object to be visible

#### func (*Image) Size

```go
func (r *Image) Size() fyne.Size
```
CurrentSize returns the current size of this rectangle object

#### func (*Image) Visible

```go
func (r *Image) Visible() bool
```
IsVisible returns true if this object is visible, false otherwise

#### type ImageFill

```go
type ImageFill int
```

ImageFill defines the different type of ways an image can stretch to fill its
space.

```go
const (
	// ImageFillStretch will scale the image to match the Size() values.
	// This is the default and does not maintain aspect ratio.
	ImageFillStretch ImageFill = iota
	// ImageFillContain makes the image fit within the object Size(),
	// centrally and maintaining aspect ratio.
	// There may be transparent sections top and bottom or left and right.
	ImageFillContain //(Fit)
	// ImageFillOriginal ensures that the container grows to the pixel dimensions
	// required to fit the original image. The aspect of the image will be maintained so,
	// as with ImageFillContain there may be transparent areas around the image.
	// Note that the minSize may be smaller than the image dimensions if scale > 1.
	ImageFillOriginal
)
```

#### type Line

```go
type Line struct {
	Position1 fyne.Position // The current top-left position of the Line
	Position2 fyne.Position // The current bottomright position of the Line
	Hidden    bool          // Is this Line currently hidden

	StrokeColor color.Color // The line stroke colour
	StrokeWidth float32     // The stroke width of the line
}
```

Line describes a coloured line primitive in a Fyne canvas. Lines are special as
they can have a negative width or height to indicate an inverse slope (i.e.
slope up vs down).

#### func  NewLine

```go
func NewLine(color color.Color) *Line
```
NewLine returns a new Line instance

#### func (*Line) Hide

```go
func (l *Line) Hide()
```
Hide will set this line to not be visible

#### func (*Line) MinSize

```go
func (l *Line) MinSize() fyne.Size
```
MinSize for a Line simply returns Size{1, 1} as there is no explicit content

#### func (*Line) Move

```go
func (l *Line) Move(pos fyne.Position)
```
Move the line object to a new position, relative to its parent / canvas

#### func (*Line) Position

```go
func (l *Line) Position() fyne.Position
```
Position gets the current top-left position of this line object, relative to its
parent / canvas

#### func (*Line) Refresh

```go
func (l *Line) Refresh()
```
Refresh causes this object to be redrawn in it's current state

#### func (*Line) Resize

```go
func (l *Line) Resize(size fyne.Size)
```
Resize sets a new bottom-right position for the line object

#### func (*Line) Show

```go
func (l *Line) Show()
```
Show will set this line to be visible

#### func (*Line) Size

```go
func (l *Line) Size() fyne.Size
```
Size returns the current size of bounding box for this line object

#### func (*Line) Visible

```go
func (l *Line) Visible() bool
```
Visible returns true if this line// Show will set this circle to be visible is
visible, false otherwise

#### type LinearGradient

```go
type LinearGradient struct {
	StartColor color.Color // The beginning color of the gradient
	EndColor   color.Color // The end color of the gradient
	Angle      float64     // The angle of the gradient (0/180 for vertical; 90/270 for horizontal)
}
```

LinearGradient defines a Gradient travelling straight at a given angle. The only
supported values for the angle are `0.0` (vertical) and `90.0` (horizontal),
currently.

#### func  NewHorizontalGradient

```go
func NewHorizontalGradient(start, end color.Color) *LinearGradient
```
NewHorizontalGradient creates a new horizontally travelling linear gradient. The
start color will be at the left of the gradient and the end color will be at the
right.

#### func  NewLinearGradient

```go
func NewLinearGradient(start, end color.Color, angle float64) *LinearGradient
```
NewLinearGradient creates a linear gradient at a the specified angle. The angle
parameter is the degree angle along which the gradient is calculated. A
NewHorizontalGradient uses 270 degrees and NewVerticalGradient is 0 degrees.

#### func  NewVerticalGradient

```go
func NewVerticalGradient(start color.Color, end color.Color) *LinearGradient
```
NewVerticalGradient creates a new vertically travelling linear gradient. The
start color will be at the top of the gradient and the end color will be at the
bottom.

#### func (*LinearGradient) Generate

```go
func (g *LinearGradient) Generate(iw, ih int) image.Image
```
Generate calculates an image of the gradient with the specified width and
height.

#### func (*LinearGradient) Hide

```go
func (r *LinearGradient) Hide()
```
Hide will set this object to not be visible

#### func (*LinearGradient) MinSize

```go
func (r *LinearGradient) MinSize() fyne.Size
```
MinSize returns the specified minimum size, if set, or {1, 1} otherwise

#### func (*LinearGradient) Move

```go
func (r *LinearGradient) Move(pos fyne.Position)
```
Move the rectangle object to a new position, relative to its parent / canvas

#### func (*LinearGradient) Position

```go
func (r *LinearGradient) Position() fyne.Position
```
CurrentPosition gets the current position of this rectangle object, relative to
its parent / canvas

#### func (*LinearGradient) Refresh

```go
func (g *LinearGradient) Refresh()
```
Refresh causes this object to be redrawn in it's current state

#### func (*LinearGradient) Resize

```go
func (r *LinearGradient) Resize(size fyne.Size)
```
Resize sets a new size for the rectangle object

#### func (*LinearGradient) SetMinSize

```go
func (r *LinearGradient) SetMinSize(size fyne.Size)
```
SetMinSize specifies the smallest size this object should be

#### func (*LinearGradient) Show

```go
func (r *LinearGradient) Show()
```
Show will set this object to be visible

#### func (*LinearGradient) Size

```go
func (r *LinearGradient) Size() fyne.Size
```
CurrentSize returns the current size of this rectangle object

#### func (*LinearGradient) Visible

```go
func (r *LinearGradient) Visible() bool
```
IsVisible returns true if this object is visible, false otherwise

#### type RadialGradient

```go
type RadialGradient struct {
	StartColor color.Color // The beginning color of the gradient
	EndColor   color.Color // The end color of the gradient
	// The offset of the center for generation of the gradient.
	// This is not a DP measure but relates to the width/height.
	// A value of 0.5 would move the center by the half width/height.
	CenterOffsetX, CenterOffsetY float64
}
```

RadialGradient defines a Gradient travelling radially from a center point
outward.

#### func  NewRadialGradient

```go
func NewRadialGradient(start, end color.Color) *RadialGradient
```
NewRadialGradient creates a new radial gradient.

#### func (*RadialGradient) Generate

```go
func (g *RadialGradient) Generate(iw, ih int) image.Image
```
Generate calculates an image of the gradient with the specified width and
height.

#### func (*RadialGradient) Hide

```go
func (r *RadialGradient) Hide()
```
Hide will set this object to not be visible

#### func (*RadialGradient) MinSize

```go
func (r *RadialGradient) MinSize() fyne.Size
```
MinSize returns the specified minimum size, if set, or {1, 1} otherwise

#### func (*RadialGradient) Move

```go
func (r *RadialGradient) Move(pos fyne.Position)
```
Move the rectangle object to a new position, relative to its parent / canvas

#### func (*RadialGradient) Position

```go
func (r *RadialGradient) Position() fyne.Position
```
CurrentPosition gets the current position of this rectangle object, relative to
its parent / canvas

#### func (*RadialGradient) Refresh

```go
func (g *RadialGradient) Refresh()
```
Refresh causes this object to be redrawn in it's current state

#### func (*RadialGradient) Resize

```go
func (r *RadialGradient) Resize(size fyne.Size)
```
Resize sets a new size for the rectangle object

#### func (*RadialGradient) SetMinSize

```go
func (r *RadialGradient) SetMinSize(size fyne.Size)
```
SetMinSize specifies the smallest size this object should be

#### func (*RadialGradient) Show

```go
func (r *RadialGradient) Show()
```
Show will set this object to be visible

#### func (*RadialGradient) Size

```go
func (r *RadialGradient) Size() fyne.Size
```
CurrentSize returns the current size of this rectangle object

#### func (*RadialGradient) Visible

```go
func (r *RadialGradient) Visible() bool
```
IsVisible returns true if this object is visible, false otherwise

#### type Raster

```go
type Raster struct {
	Generator func(w, h int) image.Image // Render the raster image from code

	Translucency float64 // Set a translucency value > 0.0 to fade the raster
}
```

Raster describes a raster image area that can render in a Fyne canvas

#### func  NewRaster

```go
func NewRaster(generate func(w, h int) image.Image) *Raster
```
NewRaster returns a new Image instance that is rendered dynamically using the
specified generate function. Images returned from this method should draw
dynamically to fill the width and height parameters passed to pixelColor.

#### func  NewRasterFromImage

```go
func NewRasterFromImage(img image.Image) *Raster
```
NewRasterFromImage returns a new Raster instance that is rendered from the Go
image.Image passed in. Rasters returned from this method will map pixel for
pixel to the screen starting img.Bounds().Min pixels from the top left of the
canvas object. Truncates rather than scales the image. If smaller than the
target space, the image will be padded with zero-pixels to the target size.

#### func  NewRasterWithPixels

```go
func NewRasterWithPixels(pixelColor func(x, y, w, h int) color.Color) *Raster
```
NewRasterWithPixels returns a new Image instance that is rendered dynamically by
iterating over the specified pixelColor function for each x, y pixel. Images
returned from this method should draw dynamically to fill the width and height
parameters passed to pixelColor.

#### func (*Raster) Alpha

```go
func (r *Raster) Alpha() float64
```
Alpha is a convenience function that returns the alpha value for a raster based
on its Translucency value. The result is 1.0 - Translucency.

#### func (*Raster) Hide

```go
func (r *Raster) Hide()
```
Hide will set this object to not be visible

#### func (*Raster) MinSize

```go
func (r *Raster) MinSize() fyne.Size
```
MinSize returns the specified minimum size, if set, or {1, 1} otherwise

#### func (*Raster) Move

```go
func (r *Raster) Move(pos fyne.Position)
```
Move the rectangle object to a new position, relative to its parent / canvas

#### func (*Raster) Position

```go
func (r *Raster) Position() fyne.Position
```
CurrentPosition gets the current position of this rectangle object, relative to
its parent / canvas

#### func (*Raster) Refresh

```go
func (r *Raster) Refresh()
```
Refresh causes this object to be redrawn in it's current state

#### func (*Raster) Resize

```go
func (r *Raster) Resize(size fyne.Size)
```
Resize sets a new size for the rectangle object

#### func (*Raster) SetMinSize

```go
func (r *Raster) SetMinSize(size fyne.Size)
```
SetMinSize specifies the smallest size this object should be

#### func (*Raster) Show

```go
func (r *Raster) Show()
```
Show will set this object to be visible

#### func (*Raster) Size

```go
func (r *Raster) Size() fyne.Size
```
CurrentSize returns the current size of this rectangle object

#### func (*Raster) Visible

```go
func (r *Raster) Visible() bool
```
IsVisible returns true if this object is visible, false otherwise

#### type Rectangle

```go
type Rectangle struct {
	FillColor   color.Color // The rectangle fill colour
	StrokeColor color.Color // The rectangle stroke colour
	StrokeWidth float32     // The stroke width of the rectangle
}
```

Rectangle describes a coloured rectangle primitive in a Fyne canvas

#### func  NewRectangle

```go
func NewRectangle(color color.Color) *Rectangle
```
NewRectangle returns a new Rectangle instance

#### func (*Rectangle) Hide

```go
func (r *Rectangle) Hide()
```
Hide will set this object to not be visible

#### func (*Rectangle) MinSize

```go
func (r *Rectangle) MinSize() fyne.Size
```
MinSize returns the specified minimum size, if set, or {1, 1} otherwise

#### func (*Rectangle) Move

```go
func (r *Rectangle) Move(pos fyne.Position)
```
Move the rectangle object to a new position, relative to its parent / canvas

#### func (*Rectangle) Position

```go
func (r *Rectangle) Position() fyne.Position
```
CurrentPosition gets the current position of this rectangle object, relative to
its parent / canvas

#### func (*Rectangle) Refresh

```go
func (r *Rectangle) Refresh()
```
Refresh causes this object to be redrawn in it's current state

#### func (*Rectangle) Resize

```go
func (r *Rectangle) Resize(size fyne.Size)
```
Resize sets a new size for the rectangle object

#### func (*Rectangle) SetMinSize

```go
func (r *Rectangle) SetMinSize(size fyne.Size)
```
SetMinSize specifies the smallest size this object should be

#### func (*Rectangle) Show

```go
func (r *Rectangle) Show()
```
Show will set this object to be visible

#### func (*Rectangle) Size

```go
func (r *Rectangle) Size() fyne.Size
```
CurrentSize returns the current size of this rectangle object

#### func (*Rectangle) Visible

```go
func (r *Rectangle) Visible() bool
```
IsVisible returns true if this object is visible, false otherwise

#### type Text

```go
type Text struct {
	Alignment fyne.TextAlign // The alignment of the text content

	Color     color.Color    // The main text draw colour
	Text      string         // The string content of this Text
	TextSize  int            // Size of the text - if the Canvas scale is 1.0 this will be equivalent to point size
	TextStyle fyne.TextStyle // The style of the text content
}
```

Text describes a text primitive in a Fyne canvas. A text object can have a style
set which will apply to the whole string. No formatting or text parsing will be
performed

#### func  NewText

```go
func NewText(text string, color color.Color) *Text
```
NewText returns a new Text implementation

#### func (*Text) Hide

```go
func (r *Text) Hide()
```
Hide will set this object to not be visible

#### func (*Text) MinSize

```go
func (t *Text) MinSize() fyne.Size
```
MinSize returns the minimum size of this text object based on its font size and
content. This is normally determined by the render implementation.

#### func (*Text) Move

```go
func (r *Text) Move(pos fyne.Position)
```
Move the rectangle object to a new position, relative to its parent / canvas

#### func (*Text) Position

```go
func (r *Text) Position() fyne.Position
```
CurrentPosition gets the current position of this rectangle object, relative to
its parent / canvas

#### func (*Text) Refresh

```go
func (t *Text) Refresh()
```
Refresh causes this object to be redrawn in it's current state

#### func (*Text) Resize

```go
func (r *Text) Resize(size fyne.Size)
```
Resize sets a new size for the rectangle object

#### func (*Text) SetMinSize

```go
func (r *Text) SetMinSize(size fyne.Size)
```
SetMinSize specifies the smallest size this object should be

#### func (*Text) Show

```go
func (r *Text) Show()
```
Show will set this object to be visible

#### func (*Text) Size

```go
func (r *Text) Size() fyne.Size
```
CurrentSize returns the current size of this rectangle object

#### func (*Text) Visible

```go
func (r *Text) Visible() bool
```
IsVisible returns true if this object is visible, false otherwise
