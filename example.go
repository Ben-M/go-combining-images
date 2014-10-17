// Here's an example of downloading images, resizing them and combining them into a single file.
// There's absolutely no error handling or concurrency, so let's add that during the Hackathon.

package main

import (
    "github.com/nfnt/resize"
    "image/jpeg"
    "image/gif"
    "image/png"
    "image/draw"
    "image"
    "net/http"
    "os"
)

func main() {

  image.RegisterFormat("jpg", "jp?g", jpeg.Decode, jpeg.DecodeConfig) 
  image.RegisterFormat("gif", "gif", gif.Decode, gif.DecodeConfig) 


  left_image_original := getImageFromUrl("http://www.mathx.net/image/simplifying_numerical_fractions.gif")
  right_image_original := getImageFromUrl("http://www.trexglobal.com/property-management/wp-content/uploads/2010/02/buying-a-cheap-foreclosure-real-estate-property.jpg")

  left_image_resized := resizeForHorizontalStrip(left_image_original)
  right_image_resized := resizeForHorizontalStrip(right_image_original)

  combined_images := combineImages(left_image_resized, right_image_resized)

  saveImageToFile(combined_images, "test_canvas.png")
}

func getImageFromUrl(url string) image.Image {

  response, _ := http.Get(url)
  img, _, _:= image.Decode(response.Body)
  return img
}

func resizeForHorizontalStrip(image image.Image) image.Image{
  return resize.Resize(0, 150, image, resize.NearestNeighbor)  
}

func combineImages(left_image image.Image, right_image image.Image) image.Image {
  
  left_image_width := getWidth(left_image)
  right_image_width := getWidth(right_image)
  total_width := left_image_width + right_image_width

  canvas := image.NewRGBA(image.Rect(0, 0, total_width, 150))

  drawImageAtPosition(canvas, left_image, image.Point{0,0})
  drawImageAtPosition(canvas, right_image, image.Point{left_image_width,0})
  
  return canvas
}

func getWidth(image image.Image) int{
  return image.Bounds().Max.X
}

func drawImageAtPosition(canvas draw.Image, image_to_draw image.Image, position image.Point) {
  sr := image.Rect(0, 0, getWidth(image_to_draw), 150)
  r := image.Rectangle{position, position.Add(sr.Size())}
  
  draw.Draw(canvas, r, image_to_draw, sr.Min, draw.Src)
  return
}

func saveImageToFile(image image.Image, filename string) {
    out, _ := os.Create(filename)
    png.Encode(out, image)
    out.Close()  
}