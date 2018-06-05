// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gi

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"

	"github.com/goki/ki"
	"github.com/goki/ki/kit"
)

// bitmap contains various bitmap-related elements, including the Bitmap node
// for showing bitmaps, and image processing utilities

// Bitmap is a Viewport2D that is optimized to render a static bitmap image --
// it expects to be a terminal node and does NOT call rendering etc on its
// children.  It is particularly useful for overlays in drag-n-drop uses --
// can grab the image of another vp and show that
type Bitmap struct {
	Viewport2D
}

var KiT_Bitmap = kit.Types.AddType(&Bitmap{}, BitmapProps)

var BitmapProps = ki.Props{
	"background-color": &Prefs.BackgroundColor,
}

// GrabRenderFrom grabs the rendered image from given node -- copies the
// vpBBox from parent viewport of node (or from viewport directly if node is a
// viewport) -- returns success
func (bm *Bitmap) GrabRenderFrom(gii Node2D) bool {
	gi := gii.AsNode2D()
	givp := gii.AsViewport2D()
	if givp != nil && givp.Pixels != nil {
		sz := givp.Pixels.Bounds().Size()
		bm.Resize(sz)
		draw.Draw(bm.Pixels, bm.Pixels.Bounds(), givp.Pixels, image.ZP, draw.Src)
		return true
	}
	givp = gi.Viewport
	if givp == nil || givp.Pixels == nil {
		log.Printf("Bitmap GrabRenderFrom could not grab from node, viewport or pixels nil: %v\n", gi.PathUnique())
		return false
	}
	if gi.VpBBox.Empty() {
		return false // offscreen -- can happen -- no warning -- just check rval
	}
	sz := gi.VpBBox.Size()
	bm.Resize(sz)
	draw.Draw(bm.Pixels, bm.Pixels.Bounds(), givp.Pixels, gi.VpBBox.Min, draw.Src)
	// todo: option to make it semi-transparent
	return true
}

func (bm *Bitmap) Render2D() {
	if bm.PushBounds() {
		bm.DrawIntoParent(bm.Viewport)
		bm.PopBounds()
	}
}

//////////////////////////////////////////////////////////////////////////////////
//  Image IO

func LoadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	im, _, err := image.Decode(file)
	return im, err
}

func LoadPNG(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return png.Decode(file)
}

func SavePNG(path string, im image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, im)
}

//////////////////////////////////////////////////////////////////////////////////
//  Image Manipulations

// ImageClearer makes an image more transparent -- pct is amount to alter
// alpha transparency factor by -- 100 = fully transparent, 0 = no change --
// affects the image itself -- make a copy of you want to keept the original
func ImageClearer(im *image.RGBA, pct float32) {
	pct = InRange32(pct, 0, 100.0)
	fact := pct / 100.0
	sz := im.Bounds().Size()
	for y := 0; y < sz.Y; y++ {
		for x := 0; x < sz.X; x++ {
			f32 := NRGBAf32Model.Convert(im.At(x, y)).(NRGBAf32)
			f32.A -= f32.A * fact
			im.Set(x, y, f32)
		}
	}
}