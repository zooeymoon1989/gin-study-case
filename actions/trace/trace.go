package trace

import (
	"github.com/gin-gonic/gin"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

const (
	output     = "img/out.png"
	width      = 2048
	height     = 2048
	numWorkers = 8
	cpuPprof   = "cpu.pprof"
)


func Trace(c *gin.Context)  {

	now := time.Now()
	file , err := os.Create(cpuPprof)
	if err != nil {
		c.JSON(200 , err.Error())
		return
	}

	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()
	defer file.Close()

	f , err := os.Create(output)
	if err != nil {
		c.JSON(200 , err.Error())
		return
	}

	img := create(width,height)

	if err = png.Encode(f, img); err != nil {
		log.Fatal(err)
	}

	c.JSON(200 , time.Since(now).Seconds())

}

func create(width , height int) image.Image {
	m := image.NewGray(image.Rect(0,0,width,height))
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			m.Set(i , j ,pixel(i,j,width,height))
		}
	}

	return m
}


// createCol creates one goroutine per column.
func createCol(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			m.Set(i, j, pixel(i, j, width, height))
		}
	}
	return m
}

// pixel returns the color of a Mandelbrot fractal at the given point.
func pixel(i, j, width, height int) color.Color {
	// Play with this constant to increase the complexity of the fractal.
	// In the justforfunc.com video this was set to 4.
	const complexity = 1024

	xi := norm(i, width, -1.0, 2)
	yi := norm(j, height, -1, 1)

	const maxI = 1000
	x, y := 0., 0.

	for i := 0; (x*x+y*y < complexity) && i < maxI; i++ {
		x, y = x*x-y*y+xi, 2*x*y+yi
	}

	return color.Gray{uint8(x)}
}

func norm(x, total int, min, max float64) float64 {
	return (max-min)*float64(x)/float64(total) - max
}
