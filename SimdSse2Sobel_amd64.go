//+build !noasm
//+build !appengine

package gocvsimd

import "unsafe"

//go:noescape
func _SimdSse2SobelDx(src unsafe.Pointer, srcStride, srcWidth, srcHeight uint64, dst unsafe.Pointer, dstStride uint64)

//go:noescape
func _SimdSse2SobelDy(src unsafe.Pointer, srcStride, srcWidth, srcHeight uint64, dst unsafe.Pointer, dstStride uint64)

//go:noescape
func _SimdSse2ContourAnchors(src unsafe.Pointer, srcStride, srcWidth, srcHeight uint64, step uint64, threshold int16, dst unsafe.Pointer, dstStride uint64)

//
func SimdSse2SobelDx(src, dst View) {

	_SimdSse2SobelDx(src.GetData(), uint64(src.GetStride()), uint64(src.GetWidth()), uint64(src.GetHeight()), dst.GetData(), uint64(dst.GetStride()))
}

//
func SimdSse2SobelDy(src, dst View) {

	_SimdSse2SobelDy(src.GetData(), uint64(src.GetStride()), uint64(src.GetWidth()), uint64(src.GetHeight()), dst.GetData(), uint64(dst.GetStride()))
}
