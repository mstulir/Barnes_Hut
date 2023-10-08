//Madison Stulir
//Barnes-Hut HW

package main

import (
	"fmt"
	"gifhelper"
	"os"
	"math"
)

func main() {
	//Running the barnes hut algorithm on galaxies gravitational simulation

	//take in user input of what simulation to run
	dataset:=os.Args[1]
	galaxies:=make([]Galaxy,0)
	//set theta (same for all datasets)
	theta:=0.5
	//initialize variables that are different for each input
	numGens:=0
	time:=0.0
	width:=0.0
	scalingFactor:=0.0

	if dataset=="jupiter"{
		var jupiter, io, europa, ganymede, callisto Star

		jupiter.red, jupiter.green, jupiter.blue = 223, 227, 202
		io.red, io.green, io.blue = 249, 249, 165
		europa.red, europa.green, europa.blue = 132, 83, 52
		ganymede.red, ganymede.green, ganymede.blue = 76, 0, 153
		callisto.red, callisto.green, callisto.blue = 0, 153, 76

		//values from wikipedia
		jupiter.mass = 1.898 * math.Pow(10, 27)
		io.mass = 8.9319 * math.Pow(10, 22)
		europa.mass = 4.7998 * math.Pow(10, 22)
		ganymede.mass = 1.4819 * math.Pow(10, 23)
		callisto.mass = 1.0759 * math.Pow(10, 23)
		//values from wikipedia
		jupiter.radius = 71000000
		io.radius = 1821000
		europa.radius = 1569000
		ganymede.radius = 2631000
		callisto.radius = 2410000

		//jupiter in the middle of the universe 4000000000 width
		jupiter.position.x, jupiter.position.y = 2000000000, 2000000000
		//positions are listed relative to jupiter
		io.position.x, io.position.y = 2000000000-421600000, 2000000000
		europa.position.x, europa.position.y = 2000000000, 2000000000+670900000
		ganymede.position.x, ganymede.position.y = 2000000000+1070400000, 2000000000
		callisto.position.x, callisto.position.y = 2000000000, 2000000000-1882700000

		//jupiter is not moving
		jupiter.velocity.x, jupiter.velocity.y = 0, 0
		//wikipedia values
		io.velocity.x, io.velocity.y = 0, -17320
		europa.velocity.x, europa.velocity.y = -13740, 0
		ganymede.velocity.x, ganymede.velocity.y = 0, 10870
		callisto.velocity.x, callisto.velocity.y = 8200, 0


		g := make(Galaxy,5)
		g[0]=&jupiter
		g[1]=&io
		g[2]=&europa
		g[3]=&ganymede
		g[4]=&callisto

		galaxies=append(galaxies,g)

		width = 4000000000.0
		numGens = 500000
		time = 1.0
		scalingFactor = 10.0 // a scaling factor is needed to inflate size of stars when drawn because galaxies are very sparse
	} else if dataset=="galaxy" {
		g0 := InitializeGalaxy(500, 4e21, 7e22, 2e22)
		galaxies = append(galaxies,g0)

		width = 1.0e23
		numGens = 100000
		time = 2e14
		scalingFactor = 1e11 // a scaling factor is needed to inflate size of stars when drawn because galaxies are very sparse
	} else if dataset=="collision" {
		g0 := InitializeGalaxy(500, 4e21, 7e22, 2e22)
		g0=PushGalaxy(g0,0)
		g1 := InitializeGalaxy(500, 4e21, 6e22, 3.5e22) //3 to 3.5
		g1=PushGalaxy(g1,1)

		galaxies =append(galaxies,g0)
		galaxies=append(galaxies,g1)

		width = 1.0e23
		numGens = 300000
		time = 2e14
		scalingFactor = 1e11 // a scaling factor is needed to inflate size of stars when drawn because galaxies are very sparse
	} else {
		panic("Improper dataset given in command line!")
	}

	//perform simulation
	initialUniverse := InitializeUniverse(galaxies, width)
	timePoints := BarnesHut(initialUniverse, numGens, time, theta)

	fmt.Println("Simulation run. Now drawing images.")
	canvasWidth := 1000
	frequency := 1000

	imageList := AnimateSystem(timePoints, canvasWidth, frequency, scalingFactor)

	fmt.Println("Images drawn. Now generating GIF.")
	gifhelper.ImagesToGIF(imageList, "galaxy")
	fmt.Println("GIF drawn.")
}
