//Madison Stulir
//Barnes-Hut HW

package main

import(
	//"fmt"
	"math"
)

//BarnesHut is our highest level function.
//Input: initial Universe object, a number of generations, and a time interval.
//Output: collection of Universe objects corresponding to updating the system
//over indicated number of generations every given time interval.
func BarnesHut(initialUniverse *Universe, numGens int, time, theta float64) []*Universe {
	timePoints := make([]*Universe, numGens+1)

	// Your code goes here. Use subroutines! :)
	timePoints[0]=initialUniverse
	for i:=1;i<numGens+1;i++{
		//fmt.Println("generation",i)
		timePoints[i]=UpdateUniverse(timePoints[i-1], time, theta)
	}
	return timePoints
}

//Update Universe updates the positions of all stars in the universe by 1 timestep
//input: a universe, a timestep for the generation and a theta value representing s/d
//output: a new universe updated by the timestep
func UpdateUniverse(universe *Universe, time, theta float64) *Universe{
	//copy the universe for the next gen
	newUniverse:=DeepCopy(universe)
	//make the quad tree
	quadTree:=MakeQuadTree(universe)

	//determine number of stars in universe
	numStars:=len(universe.stars)
	//range through all the stars and update their position velocity and acceleration to that star in new universe
	for i:=0;i<numStars;i++{
		newUniverse.stars[i].UpdateAcceleration(quadTree,theta)
		newUniverse.stars[i].UpdateVelocity(time)
		newUniverse.stars[i].UpdatePosition(time)
	}
	return newUniverse
}

//DeepCopy creates a copy of all elements of the universe to new location in the memory for the next generation to update the fields
//Input: a pointer to a universe
//Output: a pointer to a universe (a copy with new memory locations from the first universe)
func DeepCopy(universe *Universe) *Universe {
	//copy all elements of the universe
	var newUniverse Universe

	numStars:=len(universe.stars)
	newUniverse.width = universe.width
	newUniverse.stars = make([]*Star, 0)

	//initialize Stars
	for i:=0;i<numStars;i++{
		var s Star
		newUniverse.stars=append(newUniverse.stars,&s)
	}

	//for all the stars in the universe, make a copy at a new address location and all of their characteristics
	for i:=0;i<numStars;i++{
		newUniverse.stars[i].position.x=universe.stars[i].position.x
		newUniverse.stars[i].position.y=universe.stars[i].position.y
		newUniverse.stars[i].acceleration.x=universe.stars[i].acceleration.x
		newUniverse.stars[i].acceleration.y=universe.stars[i].acceleration.y
		newUniverse.stars[i].velocity.x=universe.stars[i].velocity.x
		newUniverse.stars[i].velocity.y=universe.stars[i].velocity.y
		newUniverse.stars[i].radius=universe.stars[i].radius
		newUniverse.stars[i].mass=universe.stars[i].mass
		newUniverse.stars[i].red=universe.stars[i].red
		newUniverse.stars[i].green=universe.stars[i].green
		newUniverse.stars[i].blue=universe.stars[i].blue
	}
	return &newUniverse
}

//MakeQuadTree generates a quadtree of all the stars in the universe
//Input: a universe containing a slice of stars and a width of the universe
//Output: a quadtree which is a *Node to the root node (all other nodes are children beneath it)
func MakeQuadTree(universe *Universe) QuadTree {
	//initialize quad tree and add in the universe width and make the center 0,0
	//the tree is a pointer to the root node (it will have no star (dummy star), 4 possible children and no quadrant?)
	var t QuadTree
	//all of the stars values default to zero which is desired for a dummy star
	var root Star

	var vx Node
	//need to set the nodes width to the entire universe and the x,y of the bottom left corner (0,0)
	vx.star=&root
	vx.sector.x=0
	vx.sector.y=0
	vx.sector.width=universe.width
	//give the root node 4 children (other nodes) and assign their quadrants (should contain nil stars)
	vx.AddChildren()

	t.root=&vx

	numStars:=len(universe.stars)
	//loop through the stars and add them to the tree
	for i:=0;i<numStars;i++{
		t.root.RecursiveFindStarPosition(universe.stars[i])
	}

	return t
}

//RecursiveFindStarPosition adds the current star to a QuadTree by recursively searching through nodes starting at the root and placing it into the correct quadrant until the node is unoccupied by another star
//Input: a pointer to a node, inputNode, and a pointer to a star
//Output: the QuadTree, updated to contain the star (it is a pointer so there is no real output of the function, but this is the result)
func (inputNode *Node) RecursiveFindStarPosition(inputStar *Star) {
	//check the quadrant of the star
	quad:=CheckPosition(inputStar,inputNode)
	//if the star is out of the universe, we will ignore it
	if quad==57{
	//find the position of the star within the quadTree
	} else {
		//base case: the star can be added to the location because the position is nil
		if inputNode.children[quad].star==nil{
			//update the center of gravity and mass of the inputNode, since we know that the star is below it (would this be only if it is not a nil position)
			inputNode.UpdatePositionAndMass(inputStar)
			//assign the star to that location in the tree
			inputNode.children[quad].star=inputStar
		} else {
			//inductive case: there is something besides nil in that location

			//does the subnode it falls within already have children? -- check the length of the children characteristic
			if len(inputNode.children[quad].children)>0{
				// yes: run recursive function
				//update the center of gravity and mass of the inputNode, since we know that the star is below it (would this be only if it is not a nil position)
				inputNode.UpdatePositionAndMass(inputStar)
				//call the function again! it is within this quadrant, but lower in the tree since there is already a node here
				inputNode.children[quad].RecursiveFindStarPosition(inputStar)
			} else { // the subnode does not have children but it does currently have a star in that position
				// no: need to add children
				inputNode.children[quad].AddChildren()
				//grab the pointer to the node for the current occupant star
				currentStar:=inputNode.children[quad].star
				//make a dummy star and assign it to for the parent location
				var dummy Star
				inputNode.children[quad].star=&dummy

				//update the center of gravity and mass of the inputNode, since we know that the star is below it (would this be only if it is not a nil position)
				inputNode.UpdatePositionAndMass(inputStar)

				//need to make sure we take the star occupying this location and add it back into the map before we call the recursive function
				currentStarNode:=CheckPosition(currentStar,inputNode.children[quad])

				//if the star is outside of the universe, we will ignore it
				if currentStarNode==57{
					//do not add star to tree
				} else {
					// add the current star and update the position and mass
					inputNode.children[quad].UpdatePositionAndMass(currentStar)
					inputNode.children[quad].children[currentStarNode].star=currentStar
					//call function again to add the inputStar to the tree within the subnode
					inputNode.children[quad].RecursiveFindStarPosition(inputStar)
				}
			}
		}
	}
}

//UpdatePositionAndMass takes in a node and a star and computes the new mass and position of the node's star incorporating the mass and position of the inputStar
//Input: a pointer to a node inputNode and a pointer to a star inputStar
//Output: there is no output - the nodes position and mass of its dummy star are updated to include the inputStars mass and position
func (inputNode *Node)UpdatePositionAndMass(inputStar *Star) {
	//update position of the the input node's star
	numeratorx:=inputNode.star.position.x*inputNode.star.mass
	weightedXPosition:=inputStar.position.x*inputStar.mass
	numeratory:=inputNode.star.position.y*inputNode.star.mass
	weightedYPosition:=inputStar.position.y*inputStar.mass
	//update mass of inputNodes star
	inputNode.star.mass+=inputStar.mass
	inputNode.star.position.x=(numeratorx+weightedXPosition)/inputNode.star.mass
	inputNode.star.position.y=(numeratory+weightedYPosition)/inputNode.star.mass
}

//CheckPosition determine from an input node and star which qudrant of the node the star belongs to
//Input: a star and a node
//Output: an integer between 0 and 3 representing which quadrant the star belongs to
func CheckPosition(star *Star, node *Node) int {
	//loop through the 4 children and see which quad it goes into
	for i:=0;i<4;i++{
		if star.position.x>=node.children[i].sector.x && star.position.x<=(node.children[i].sector.x+node.children[i].sector.width) && star.position.y>=node.children[i].sector.y && star.position.y<=(node.children[i].sector.y+node.children[i].sector.width) {
			return i
		}
	}
	//the star is not within a quadrant -- 57 is a random number that we will use in the other function to know to exclude this star from the tree
	return 57
}

//AddChildren gives the current node 4 node children and sets their quadrants as 1/4 of their parent input node
//Input: a pointer to a node
//Output: the pointer is updated to contain 4 children nodes with nil stars and quadrants specified
func (vx *Node) AddChildren() {
	vx.children=make([]*Node,0)
	//initialize 4 child Nodes
	for i:=0;i<4;i++{
		var n Node
		vx.children=append(vx.children,&n)
	}
	//assign the quadrants for each of the 4 nodes
	//node0=NW, node1=NE, node2=SW,node3=SE
	vx.children[0].sector.x=vx.sector.x
	vx.children[0].sector.y=vx.sector.y+(vx.sector.width/2)
	vx.children[0].sector.width=vx.sector.width/2
	vx.children[1].sector.x=vx.sector.x+(vx.sector.width/2)
	vx.children[1].sector.y=vx.sector.y+(vx.sector.width/2)
	vx.children[1].sector.width=vx.sector.width/2
	vx.children[2].sector.x=vx.sector.x
	vx.children[2].sector.y=vx.sector.y
	vx.children[2].sector.width=vx.sector.width/2
	vx.children[3].sector.x=vx.sector.x+(vx.sector.width/2)
	vx.children[3].sector.y=vx.sector.y
	vx.children[3].sector.width=vx.sector.width/2
}

//UpdateAcceleration determines the new acceleration of a star based on the gravitational forces from other stars in the universe
//Input: a pointer to the star we are updating, the quadTree of the universe, and theta representing the distance for which a force is relevant
//Output: the ordered pair corresponding to the updated acceleration is updated for the pointer star
func (inputStar *Star)UpdateAcceleration(quadTree QuadTree, theta float64) {
	var accel OrderedPair
	var force OrderedPair
	//compute forces from the 4 children of the root node
	force0:=ComputeNetForce(inputStar, quadTree.root.children[0],theta)
	force1:=ComputeNetForce(inputStar, quadTree.root.children[1],theta)
	force2:=ComputeNetForce(inputStar, quadTree.root.children[2],theta)
	force3:=ComputeNetForce(inputStar, quadTree.root.children[3],theta)
	//add the 4 forces together componentwise
	force.x=0+force0.x+force1.x+force2.x+force3.x
	force.y=0+force0.y+force1.y+force2.y+force3.y
	accel.x=force.x/inputStar.mass
	accel.y=force.y/inputStar.mass
	//update the acceleration
	inputStar.acceleration=accel
}

//ComputeNetForce is a recursive function that determines the forces from all stars beneath a given node in the quad tree on the inputStar
//Input: an inputStar pointer to a star, a pointer to a node within the quadTree, and theta the statistic used to determine whether to traverse the tree or compute the force of all the bodies beneath the node as one force
//Output: an ordered pair representing the netforce from that node and below
func ComputeNetForce(inputStar *Star, node *Node, theta float64) OrderedPair {
	var netForce OrderedPair
	// the force is zero if the node has no star or the star is the star itself
	if node.star==nil || node.star==inputStar {
		netForce.x=0
		netForce.y=0
		return netForce
	} else { // there is a star to calculate the force from
		//s is the width of the ancestor
		s:= node.sector.width
		//d is the distance from the inPut star to the star we are checking
		d:=Distance(inputStar.position,node.star.position)
		statistic:=s/d
		//inductiveStep --  we want to explore this node deeper since statistic is greater than or equal to theta
		if statistic>=theta{
			//if it has children, compute the force from the 4 children separately
			if len(node.children)>0{
				force0:=ComputeNetForce(inputStar,node.children[0],theta)
				force1:=ComputeNetForce(inputStar,node.children[1],theta)
				force2:=ComputeNetForce(inputStar,node.children[2],theta)
				force3:=ComputeNetForce(inputStar,node.children[3],theta)
				netForce.x=0+force0.x+force1.x+force2.x+force3.x
				netForce.y=0+force0.y+force1.y+force2.y+force3.y
			} else {
				//if it has no children, compute the Force
				netForce=ComputeForce(inputStar,node.star)
			}
			return netForce
		} else {
			//basecase
			//if s/d is less than theta, compute the force gtom all stars under this star as one force and add it to netForce and return that value
			netForce:=ComputeForce(inputStar,node.star)
			return netForce
		}
	}
}


//ComputeForce determines the force between 2 star objects based on their postions and masses, as well as the gravitational constant G
//Input: two star objects b1 and b2
//Output: The force due to gravity (as a vector) acting on b1 subject to b2
func ComputeForce(b1,b2 *Star) OrderedPair {
	var force OrderedPair
	dist:=Distance(b1.position,b2.position)
	if dist==0.0{ // they are in the same spot so there is no force
		force.x=0.0
		force.y=0.0
	} else { // they are in different locations
		F:=G*b1.mass*b2.mass/(dist*dist)
		deltaX:=b2.position.x-b1.position.x
		deltaY:=b2.position.y-b1.position.y
		force.x=F*deltaX/dist //equal to cos theta
		force.y=F*deltaY/dist //equal to sin theta
	}

	return force
}

//Distance takes two position ordered pairs and it returns the distance between these two points in 2-D space.
//Input: 2 orderedpairs representing positions
//Output: a float64 of the distance between the points
func Distance(p1, p2 OrderedPair) float64 {
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

//UpdateVelocity changes the velocity of the star based on the new acceleration and previous velocity
//Input: a pointer to a star and a time step (float64)
//Output: the orederedpair corresponding to the velocity of this object after a single time step, using the stars's current acceleration
func (b *Star) UpdateVelocity(time float64) {
	var vel OrderedPair
	//new velocity is current velocity + acceleration *time
	vel.x=b.velocity.x + b.acceleration.x*time
	vel.y=b.velocity.y + b.acceleration.y*time

	b.velocity=vel
}
//UpdatePosition changes the position of the star based on its previous position and updated acceleration and velocity
//Input: a pointer to a star and a timestep float64
//Output: the orderedpair corresponding to the updated position of the star after a single time step using the stars current acceleration and velocity
func (b *Star)UpdatePosition(time float64) {
	var pos OrderedPair

	pos.x=0.5*b.acceleration.x*(time*time) +b.velocity.x*time +b.position.x
	pos.y=0.5*b.acceleration.y*(time*time) +b.velocity.y*time +b.position.y
	//fmt.Println("oldpos,newpos",b.position,pos)
	b.position=pos
}
