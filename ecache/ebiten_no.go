//go:build gtxt

package ecache

import "image"

// Same as etxt.GlyphMask.
type GlyphMask = *image.Alpha

const constMaskSizeFactor = 56

func GlyphMaskByteSize(mask GlyphMask) uint32 {
	if mask == nil {
		return constMaskSizeFactor
	}
	w, h := mask.Rect.Dx(), mask.Rect.Dy()
	return maskDimsByteSize(w, h)
}

func maskDimsByteSize(width, height int) uint32 {
	return uint32(width*height) + constMaskSizeFactor
}
