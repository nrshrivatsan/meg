//go version go1.3.3 darwin/amd64

package main

import (
     "fmt"
     "bufio"
     "image"
     "image/color"
     "image/draw"
     "image/jpeg"
     "math/rand"
     "time"
     "os"
)


var dim int = 100;

var width int = 800;
var height int = 800;

var x0 int= 5
var y0 int= 5
var x1 int= 10
var y1 int= 10
var increment int = 5;

var prime color.RGBA = color.RGBA{0, 255, 255, 0}
var composite color.RGBA = color.RGBA{255, 255, 255, 255}

var img *image.RGBA = image.NewRGBA(image.Rect(0, 0, width, height))

func writeTofile(){
     if imgFileJpeg, err := os.Create("img.jpeg"); err != nil {
         fmt.Println("Jpeg error: ", err)
     } else {
         defer imgFileJpeg.Close()
         jpeg.Encode(bufio.NewWriter(imgFileJpeg), img,&jpeg.Options{jpeg.DefaultQuality})
     }
}

func paintAll( arr *[][]int){
    x0 = 0        
    y0 = 0 

    x1 = increment
    y1 = increment

    array := *arr
    for i1 := 0; i1 < len(array) ; i1++ {
        
        //Relocating the horizontal cursor to the starting of the image
        x0 = 0
        x1 = increment

        for i2 := 0 ; i2 < len(array[i1]) ; i2++{

            value := array[i1][i2]

            if( value == 0 ){
                //Asynchronously painting the image
                go draw.Draw(img, image.Rect(x0, y0, x1, y1), &image.Uniform{composite},image.ZP, draw.Src)
            }else{
                //Asynchronously painting the image
                go draw.Draw(img, image.Rect(x0, y0, x1, y1), &image.Uniform{prime},image.ZP, draw.Src)
            }        

                //Navigating the horizontal cursors to the left
                x0 += increment;
                x1 += increment;
            }

        //Navigating the vertical cursors to one step below the current position   
        y1 += increment;
        y0 += increment;   
     }

     writeTofile();
}

func main() {

    
    array := make([][]int,dim,dim);
    rand.Seed( time.Now().UnixNano())
    fmt.Println(len(array))
    fmt.Println(len(array[0]))
     for j:=0 ; j< len(array); j++ {
        array[j] = make([]int,dim);
        for k:=0 ; k<len(array[j]); k++ {
            if( time.Now().UnixNano() %2 ==0 ){
                array[j][k] = 0
            }else{
                array[j][k] = 1
            }
        }
     }
     
     //Asynch painting of the image based on the array values
     //1 => prime
     //0 => composite   
     paintAll(&array)
     

     writeTofile();

    
}
