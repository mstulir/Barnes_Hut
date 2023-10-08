//Madison Stulir
//Barnes-Hut HW

//Testing
package main

import(
  "testing"
  "fmt"
  "math"
)

//test UpdateAcceleration  (inputStar *Star)UpdateAcceleration(quadTree QuadTree, theta float64)
func TestUpdateAcceleration(t *testing.T) {
  fmt.Println("Testing UpdateAcceleration!")
  type test struct {
    star *Star
    tree QuadTree
    theta float64
    answer *Star
  }
  tests:=make([]test,1)
  tests[0].theta=0.5
  var star1 Star
  star1.position.x=20
  star1.position.y=70
  star1.mass=1
  var star2 Star
  star2.position.x=60
  star2.position.y=80
  star2.mass=1
  var star3 Star
  star3.position.x=56
  star3.position.y=65
  star3.mass=1
  var star4 Star
  star4.position.x=76
  star4.position.y=60
  star4.mass=1
  var star5 Star
  star5.position.x=75
  star5.position.y=20
  star5.mass=1
  //define the answer QuadTree
  var tree QuadTree
	//all of the stars values default to zero which is desired for a dummy star
	var root Star
  root.position.x=57.4
  root.position.y=59
  root.mass=5
	//fmt.Println(root)
	var vx Node
	//need to set the nodes width to the entire universe and the x,y of the bottom left corner are -width and -width
	vx.star=&root
	vx.sector.x=0
	vx.sector.y=0
	vx.sector.width=100
	//give the root node 4 children (other nodes) and assign their quadrants (should contain nil stars)
	vx.AddChildren()
	tree.root=&vx
  tree.root.children[0].star=&star1
  tree.root.children[3].star=&star5
  //add children to node[1]
  var dummy Star
  dummy.position.x=64
  dummy.position.y=68.33
  dummy.mass=3
  tree.root.children[1].star=&dummy
  tree.root.children[1].AddChildren()
  tree.root.children[1].children[0].star=&star2
  tree.root.children[1].children[2].star=&star3
  tree.root.children[1].children[3].star=&star4

  tests[0].tree=tree

  tests[0].star=&star1

  var answerStar Star
  answerStar.position.x=20
  answerStar.position.y=70
  answerStar.mass=1
  answerStar.acceleration.x=1.1737E-13
  answerStar.acceleration.y=-9.18E-15

  tests[0].answer=&answerStar


  for i, test := range tests {
		tests[i].star.UpdateAcceleration(test.tree,test.theta)
    var numDigits uint = 17
    var outcome OrderedPair
    outcome.x=roundFloat(tests[i].star.acceleration.x,numDigits)
    outcome.y=roundFloat(tests[i].star.acceleration.y,numDigits)
		if outcome != test.answer.acceleration {
			t.Errorf("Error! For input test dataset %d, your code gives %v,%v and the correct acceleration is %v,%v", i, outcome.x, outcome.y, test.answer.acceleration.x,test.answer.acceleration.y)
		} else {
			fmt.Println("Correct! The acceleration is", test.answer)
		}
	}



}

//roundFloat is used to round the output of function calls to the number of digits in the answers input to the function
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

//test UpdatePosition
func TestUpdatePosition(t *testing.T) {
  fmt.Println("Testing UpdatePosition!")
  type test struct {
    star *Star
    timeStep float64
    //width float64
    answer OrderedPair
  }
  tests:=make([]test,1)
  var star Star
  //assign test values hard coded
  //test a normal case in the middle of the board
  star.position.x=100
  star.position.y=1000
  star.velocity.x=1
  star.velocity.y=0.5
  star.acceleration.x=0.2
  star.acceleration.y=0.7
  tests[0].timeStep=1.0
  //tests[1].width=2000
  tests[0].answer.x=101.1
  tests[0].answer.y=1000.85
  tests[0].star=&star

  for i, test := range tests {
		tests[i].star.UpdatePosition(test.timeStep)
    var numDigits uint = 4
    var outcome OrderedPair
    outcome.x=roundFloat(tests[i].star.position.x,numDigits)
    outcome.y=roundFloat(tests[i].star.position.y,numDigits)
		if outcome != test.answer {
			t.Errorf("Error! For input test dataset %d, your code gives %v,%v and the correct position is %v,%v", i, outcome.x, outcome.y, test.answer.x,test.answer.y)
		} else {
			fmt.Println("Correct! The position is", test.answer)
		}
	}
}

//test UpdateVelocity
func TestUpdateVelocity(t *testing.T) {
  fmt.Println("Testing UpdateVelocity!")
  type test struct {
    star *Star
    timeStep float64
    answer OrderedPair
  }
  tests:=make([]test,1)
  var star Star
  star.velocity.x=.3
  star.velocity.y=.4
  star.acceleration.x=.1
  star.acceleration.y=.5

  tests[0].timeStep=1
  tests[0].answer.x=0.4
  tests[0].answer.y=0.9
  tests[0].star=&star

  for i, test := range tests {
		tests[i].star.UpdateVelocity(test.timeStep)
    var outcome OrderedPair
    var numDigits uint = 1
    outcome.x=roundFloat(tests[i].star.velocity.x,numDigits)
    outcome.y=roundFloat(tests[i].star.velocity.y,numDigits)
		if outcome != test.answer {
			t.Errorf("Error! For input test dataset %d, your code gives %v,%v and the correct velocity is %v,%v", i, outcome.x,outcome.y, test.answer.x,test.answer.y)
		} else {
			fmt.Println("Correct! The velocity is", test.answer)
		}
	}
}

//test Distance
func TestDistance(t *testing.T) {
  fmt.Println("Testing Distance!")
  type test struct {
    p1 OrderedPair
    p2 OrderedPair
    answer float64
  }
  tests:=make([]test,2)
  tests[0].p1.x=3.0
  tests[0].p1.y=4.0
  tests[0].p2.x=32.0
  tests[0].p2.y=57.0
  tests[0].answer=60.4152


  tests[1].p1.x=100.0
  tests[1].p1.y=500.0
  tests[1].p2.x=103.0
  tests[1].p2.y=401.0
  tests[1].answer=99.0454

  for i, test := range tests {
		outcome := Distance(test.p1,test.p2)
    var numDigits uint = 4
    outcome=roundFloat(outcome,numDigits)
		if outcome != test.answer {
			t.Errorf("Error! For input test dataset %d, your code gives %v, and the correct distance is %v", i, outcome, test.answer)
		} else {
			fmt.Println("Correct! The distance is", test.answer)
		}
	}
}


//test ComputeForce
//func ComputeForce(b1,b2 *Star) OrderedPair { -- uses their positions and masses
func TestComputeForce(t *testing.T) {
  fmt.Println("Testing ComputeForce!")
  type test struct {
    star1 *Star
    star2 *Star
    answer OrderedPair
  }
  tests:=make([]test,1)
  var star1 Star
  var star2 Star
  star1.position.x=100.0
  star1.position.y=500.0
  star1.mass=4.0
  star2.position.x=103.0
  star2.position.y=401.0
  star2.mass=3.0

  tests[0].star1=&star1
  tests[0].star2=&star2

  tests[0].answer.x=2.5E-15
  tests[0].answer.y=-8.16E-14

  for i, test := range tests {
    var outcome OrderedPair
    outcome=ComputeForce(tests[i].star1,tests[i].star2)
    var numDigits uint = 16
    outcome.x=roundFloat(outcome.x,numDigits)
    outcome.y=roundFloat(outcome.y,numDigits)
		if outcome != test.answer {
			t.Errorf("Error! For input test dataset %d, your code gives %v,%v and the correct force is %v,%v", i, outcome.x,outcome.y, test.answer.x,test.answer.y)
		} else {
			fmt.Println("Correct! The force is", test.answer)
		}
	}



}

//test AddChildren
func TestAddChildren(t *testing.T) {
  fmt.Println("Testing AddChildren!")
  type test struct {
    node *Node
    answer int
  }
  tests:=make([]test,1)
  var node Node
  tests[0].node=&node

  tests[0].answer=4

  for i, test := range tests {
    tests[i].node.AddChildren()
    outcome:=len(tests[i].node.children)
		if outcome != test.answer {
			t.Errorf("Error! For input test dataset %d, your code gives %d and the correct number of children is %d", i, outcome, test.answer)
		} else {
			fmt.Println("Correct! The number of children is", test.answer)
		}

  }

}


//test CheckPosition
func TestCheckPosition(t *testing.T) {
  fmt.Println("Testing CheckPosition!")
  type test struct {
    node *Node
    star *Star
    answer int
  }
  tests:=make([]test,2)
  var node Node
  var star Star
  node.sector.x=0
  node.sector.y=0
  node.sector.width=20
  //give the node 4 children
  node.AddChildren()

  star.position.x=11
  star.position.y=11

  tests[0].node=&node
  tests[0].star=&star

  tests[0].answer=1

  var node2 Node
  var star2 Star
  node2.sector.x=0
  node2.sector.y=0
  node2.sector.width=20
  //give the node 4 children
  node2.AddChildren()

  star2.position.x=10
  star2.position.y=10

  tests[1].node=&node2
  tests[1].star=&star2

  tests[1].answer=0
  for i, test := range tests {
    outcome:=CheckPosition(tests[i].star, tests[i].node)
		if outcome != test.answer {
			t.Errorf("Error! For input test dataset %d, your code gives %d and the correct quadrant is %d", i, outcome, test.answer)
		} else {
			fmt.Println("Correct! The quadrant is", test.answer)
		}
  }
}

//RecursiveFindStarPosition
func TestRecursiveFindStarPosition(t *testing.T) {
  fmt.Println("Testing RecursiveFindStarPosition!")
  type test struct {
    node *Node
    star *Star
    answer *Node
  }
  tests:=make([]test,1)

  var root Node
  root.sector.x=0.0
  root.sector.y=0.0
  root.sector.width=20.0

  var rootStar Star
  rootStar.position.x=10.0
  rootStar.position.y=15.0
  rootStar.mass=2.0

  root.star=&rootStar
  root.AddChildren()

  var A Star
  A.mass=1.0
  A.position.x=6.0
  A.position.y=11.0
  root.children[0].star=&A

  var B Star
  B.mass=1.0
  B.position.x=14.0
  B.position.y=19.0
  root.children[1].star=&B


  var C Star
  C.mass=1.0
  C.position.x=12.0
  C.position.y=13.0

  tests[0].node=&root
  tests[0].star=&C

  //defining the QuadTree of the answer
  var rootAns Node
  rootAns.sector.x=0.0
  rootAns.sector.y=0.0
  rootAns.sector.width=20.0

  var rootStarAns Star
  rootStarAns.position.x=10.67
  rootStarAns.position.y=14.33
  rootStarAns.mass=3.0

  rootAns.star=&rootStarAns
  rootAns.AddChildren()

  var dummyNode Node


  var dummy Star
  dummy.mass=2.0
  dummy.position.x=13.0
  dummy.position.y=16.0

  dummyNode.star=&dummy

  dummyNode.AddChildren()
  dummyNode.children[0].star=&B
  dummyNode.children[2].star=&C

  rootAns.children[0].star=&A
  rootAns.children[1]=&dummyNode

  tests[0].answer=&rootAns
  //is the mass of the root updated correctly?
  for i := range tests {
    tests[i].node.RecursiveFindStarPosition(tests[i].star)
    outcomeMass:=tests[i].node.star.mass
    answerMass:=tests[i].answer.star.mass
		if outcomeMass != answerMass {
			t.Errorf("Error! For input test dataset %d, your code gives %v and the mass of root is %v", i, outcomeMass, answerMass)
		} else {
			fmt.Println("Correct! The mass of the root is", answerMass)
		}
    //is the x position of the root updated correctly?
    outcomeXPos:=tests[i].node.star.position.x
    answerXPos:=tests[i].answer.star.position.x
		if outcomeMass != answerMass {
			t.Errorf("Error! For input test dataset %d, your code gives %v and the x position of root is %v", i, outcomeXPos, answerXPos)
		} else {
			fmt.Println("Correct! The x position of the root is", answerXPos)
		}

    //is the new node made correctly?
    outcomeChildren:=len(tests[i].node.children[1].children)
    answerChildren:=len(tests[i].answer.children[1].children)
    if outcomeChildren!=answerChildren {
      t.Errorf("Error! For input test dataset %d, your code gives %d and the new dummy node is %d", i, outcomeChildren, answerChildren)
    } else {
      fmt.Println("Correct! The number of children is", answerChildren)
    }
    //is the star put in the right spot?
    outcomeCPosition:=tests[i].node.children[1].children[2].star
    answerCPosition:=tests[i].answer.children[1].children[2].star
    fmt.Println("C",outcomeCPosition, answerCPosition)
    if outcomeCPosition!=answerCPosition{
      t.Errorf("Error! For input test dataset %d, your code gives star C and the number of children is star 2", i)

    } else {
      fmt.Println("Correct! The star in that position is", answerCPosition)
    }

    //is the star B moved correctly in the tree?
    outcomeBPosition:=tests[i].node.children[1].children[0].star
    answerBPosition:=tests[i].answer.children[1].children[0].star
    fmt.Println("B",outcomeBPosition, answerBPosition)
    if outcomeBPosition!=answerBPosition{
      t.Errorf("Error! For input test dataset %d, your code gives star B and the number of children is star 2", i)
      fmt.Println(outcomeBPosition, answerBPosition)
    } else {
      fmt.Println("Correct! The star in the B position is", answerBPosition)
    }

    //is the dummy node the correct mass and center of gravity
    outcomedummyPosition:=tests[i].node.children[1].star
    answerdummyPosition:=tests[i].answer.children[1].star
    if outcomedummyPosition.mass!=answerdummyPosition.mass || outcomedummyPosition.position.x!=answerdummyPosition.position.x || outcomedummyPosition.position.y!=answerdummyPosition.position.y {
      fmt.Println("dummy",outcomedummyPosition, answerdummyPosition)
      t.Errorf("Error! For input test dataset %d, your code gives star dummy and the number of children is star 2", i)
    } else {
      fmt.Println("Correct! The star in that dummy position is", answerdummyPosition)
    }

  }
}


//test UpdatePositionAndMass
func TestUpdatePositionAndMass(t *testing.T) {
  fmt.Println("Testing UpdatePositionAndMass!")
  type test struct {
    node *Node
    star *Star
    answer *Node
  }
  //takes in a node and a star which has a mass and a position
  // the nodes star should be updated to contain the new star in its position and mass
  tests:=make([]test,1)
  var node Node
  var star Star
  star.position.x=10
  star.position.y=15
  star.mass=2

  node.star=&star
  tests[0].node=&node

  var addedStar Star
  addedStar.position.x=12
  addedStar.position.y=13
  addedStar.mass=1
  tests[0].star=&addedStar

  var answer Node
  var answerStar Star
  answerStar.position.x=10.67
  answerStar.position.y=14.33
  answerStar.mass=3

  answer.star=&answerStar
  tests[0].answer=&answer

  for i := range tests {
    tests[i].node.UpdatePositionAndMass(tests[i].star)
    outcomeMass:=tests[i].node.star.mass
    answerMass:=tests[i].answer.star.mass
		if outcomeMass != answerMass {
			t.Errorf("Error! For input test dataset %d, your code gives %v and the mass of root is %v", i, outcomeMass, answerMass)
		} else {
			fmt.Println("Correct! The mass of the root is", answerMass)
		}
    outcomeX:=roundFloat(tests[i].node.star.position.x,2)
    answerX:=tests[i].answer.star.position.x
    if outcomeX != answerX {
			t.Errorf("Error! For input test dataset %d, your code gives %v and the X position of root is %v", i, outcomeX, answerX)
		} else {
			fmt.Println("Correct! The X position is", answerX)
		}
    outcomeY:=tests[i].node.star.position.y
    answerY:=tests[i].answer.star.position.y
    if outcomeMass != answerMass {
			t.Errorf("Error! For input test dataset %d, your code gives %v and the X position of root is %v", i, outcomeY, answerY)
		} else {
			fmt.Println("Correct! The X position of the root is", answerY)
		}
  }

}

//MakeQuadTree
func TestMakeQuadTree(t *testing.T) {
  fmt.Println("Testing MakeQuadTree!")
  type test struct {
    universe *Universe
    answer QuadTree
  }
  tests:=make([]test,1)

  var star1 Star
  star1.position.x=20
  star1.position.y=70
  star1.mass=1
  var star2 Star
  star2.position.x=60
  star2.position.y=80
  star2.mass=1
  var star3 Star
  star3.position.x=56
  star3.position.y=65
  star3.mass=1
  var star4 Star
  star4.position.x=76
  star4.position.y=60
  star4.mass=1
  var star5 Star
  star5.position.x=75
  star5.position.y=20
  star5.mass=1
  //add all the stars to the universe
  var universe Universe
  universe.width=100
  universe.stars = make([]*Star, 0)
  universe.stars=append(universe.stars,&star1)
  universe.stars=append(universe.stars,&star2)
  universe.stars=append(universe.stars,&star3)
  universe.stars=append(universe.stars,&star4)
  universe.stars=append(universe.stars,&star5)
  tests[0].universe=&universe

  //define the answer QuadTree
  var tree QuadTree
	//all of the stars values default to zero which is desired for a dummy star
	var root Star
  root.position.x=57.4
  root.position.y=59
  root.mass=5
	//fmt.Println(root)
	var vx Node
	//need to set the nodes width to the entire universe and the x,y of the bottom left corner are -width and -width
	vx.star=&root
	vx.sector.x=0
	vx.sector.y=0
	vx.sector.width=universe.width
	//give the root node 4 children (other nodes) and assign their quadrants (should contain nil stars)
	vx.AddChildren()
	tree.root=&vx
  tree.root.children[0].star=&star1
  tree.root.children[3].star=&star5
  //add children to node[1]
  var dummy Star
  dummy.position.x=64
  dummy.position.y=68.33
  dummy.mass=3
  tree.root.children[1].star=&dummy
  tree.root.children[1].AddChildren()
  tree.root.children[1].children[0].star=&star2
  tree.root.children[1].children[2].star=&star3
  tree.root.children[1].children[3].star=&star4

  tests[0].answer=tree
  for i := range tests {
    outcome:=MakeQuadTree(tests[i].universe)
    //check the root nodes position and masses
    outcomeMass:=outcome.root.star.mass
    answerMass:=tests[i].answer.root.star.mass
		if outcomeMass != answerMass {
			t.Errorf("Error! For input test dataset %d, your code gives %v and the mass of root is %v", i, outcomeMass, answerMass)
		} else {
			fmt.Println("Correct! The mass of the root is", answerMass)
		}
    outcomePosX:=outcome.root.star.position.x
    answerPosX:=tests[i].answer.root.star.position.x
    if outcomePosX != answerPosX {
      t.Errorf("Error! For input test dataset %d, your code gives %v and the x position of root is %v", i, outcomePosX, answerPosX)
    } else {
      fmt.Println("Correct! The mass of the root is", answerPosX)
    }
    outcomePosY:=outcome.root.star.position.y
    answerPosY:=tests[i].answer.root.star.position.y
    if outcomePosY != answerPosY {
      t.Errorf("Error! For input test dataset %d, your code gives %v and the y position of root is %v", i, outcomePosY, answerPosY)
    } else {
      fmt.Println("Correct! The mass of the root is", answerPosY)
    }

    //check the internal dummy node postion and mass
    outcomeMass=outcome.root.children[1].star.mass
    answerMass=tests[i].answer.root.children[1].star.mass
		if outcomeMass != answerMass {
			t.Errorf("Error! For input test dataset %d, your code gives %v and the mass of dummy is %v", i, outcomeMass, answerMass)
		} else {
			fmt.Println("Correct! The mass of the dummy is", answerMass)
		}
    outcomePosX=roundFloat(outcome.root.children[1].star.position.x,2)
    answerPosX=tests[i].answer.root.children[1].star.position.x
    if outcomePosX != answerPosX {
      t.Errorf("Error! For input test dataset %d, your code gives %v and the x position of dummy is %v", i, outcomePosX, answerPosX)
    } else {
      fmt.Println("Correct! The mass of the dummy is", answerPosX)
    }
    outcomePosY=roundFloat(outcome.root.children[1].star.position.y,2)
    answerPosY=roundFloat(tests[i].answer.root.children[1].star.position.y,2)
    if outcomePosY != answerPosY {
      t.Errorf("Error! For input test dataset %d, your code gives %v and the y position of dummy is %v", i, outcomePosY, answerPosY)
    } else {
      fmt.Println("Correct! The mass of the dummy is", answerPosY)
    }

    //check the other quadrant masses and positions -- 0,2,3 (0 will be star 1, 3 will be star 5)
    outcomeMass=outcome.root.children[0].star.mass
    answerMass=tests[i].answer.root.children[0].star.mass
		if outcomeMass != answerMass {
			t.Errorf("Error! For input test dataset %d, your code gives %v and the mass of quad 0 is %v", i, outcomeMass, answerMass)
		} else {
			fmt.Println("Correct! The mass of quad 0 is", answerMass)
		}
    outcomePosX=roundFloat(outcome.root.children[0].star.position.x,2)
    answerPosX=tests[i].answer.root.children[0].star.position.x
    if outcomePosX != answerPosX {
      t.Errorf("Error! For input test dataset %d, your code gives %v and the x position of quad 0 is %v", i, outcomePosX, answerPosX)
    } else {
      fmt.Println("Correct! The mass of quad 0 is", answerPosX)
    }
    outcomePosY=roundFloat(outcome.root.children[0].star.position.y,2)
    answerPosY=roundFloat(tests[i].answer.root.children[0].star.position.y,2)
    if outcomePosY != answerPosY {
      t.Errorf("Error! For input test dataset %d, your code gives %v and the y position of quad 0 is %v", i, outcomePosY, answerPosY)
    } else {
      fmt.Println("Correct! The mass of quad 0 is", answerPosY)
    }

    outcomeMass=outcome.root.children[3].star.mass
    answerMass=tests[i].answer.root.children[3].star.mass
		if outcomeMass != answerMass {
			t.Errorf("Error! For input test dataset %d, your code gives %v and the mass of quad 3 is %v", i, outcomeMass, answerMass)
		} else {
			fmt.Println("Correct! The mass of quad 3 is", answerMass)
		}
    outcomePosX=roundFloat(outcome.root.children[3].star.position.x,2)
    answerPosX=tests[i].answer.root.children[3].star.position.x
    if outcomePosX != answerPosX {
      t.Errorf("Error! For input test dataset %d, your code gives %v and the x position of quad 3 is %v", i, outcomePosX, answerPosX)
    } else {
      fmt.Println("Correct! The mass of quad 3 is", answerPosX)
    }
    outcomePosY=roundFloat(outcome.root.children[3].star.position.y,2)
    answerPosY=roundFloat(tests[i].answer.root.children[3].star.position.y,2)
    if outcomePosY != answerPosY {
      t.Errorf("Error! For input test dataset %d, your code gives %v and the y position of quad 3 is %v", i, outcomePosY, answerPosY)
    } else {
      fmt.Println("Correct! The mass of quad 3 is", answerPosY)
    }

    //see what the 4 children of the tree root look like --  do they look correct (is quad 2 nil?)
    fmt.Println("root children",outcome.root.children[0].star, outcome.root.children[1].star,outcome.root.children[2].star,outcome.root.children[3].star)

    //check that the children occupy their correct quadrants within the dummy
    fmt.Println("dummy children",outcome.root.children[1].children[0].star, outcome.root.children[1].children[1].star,outcome.root.children[1].children[2].star,outcome.root.children[1].children[3].star)
  }
}


// test deepcopy -- takes in a universe and creates a universe
func TestDeepCopy(t *testing.T) {
  fmt.Println("Testing DeepCopy!")
  type test struct {
    universe *Universe
    answer *Universe
  }
  tests:=make([]test,1)
  var u Universe
  u.width=20

  //make stars to put into the universe []*Star
  var star1 Star
  var star2 Star
  star1.position.x=100.0
  star1.position.y=500.0
  star1.mass=4.0
  star2.position.x=103.0
  star2.position.y=401.0
  star2.mass=3.0
  u.stars=make([]*Star,0)

  u.stars=append(u.stars,&star1)
  u.stars=append(u.stars,&star2)
  tests[0].universe=&u

  var u2 Universe
  u2.width=20

  var star3 Star
  var star4 Star
  star3.position.x=100.0
  star3.position.y=500.0
  star3.mass=4.0
  star4.position.x=103.0
  star4.position.y=401.0
  star4.mass=3.0
  u2.stars=make([]*Star,0)

  u2.stars=append(u2.stars,&star3)
  u2.stars=append(u2.stars,&star4)
  tests[0].answer=&u2
  for i := range tests {
    outcome:=DeepCopy(tests[i].universe)
		if outcome.width != tests[i].answer.width {
			t.Errorf("Error! For input test dataset %d, your code gives %v and the width of the universe is %v", i, outcome.width, tests[i].answer.width)
		} else {
			fmt.Println("Correct! The mass of the root is", tests[i].answer.width)
		}
    if outcome.stars[0].position.x == tests[i].answer.stars[0].position.x && outcome.stars[0].position.y == tests[i].answer.stars[0].position.y && outcome.stars[0].mass == tests[i].answer.stars[0].mass{
      fmt.Println("Correct! The star 0 is", tests[i].answer.stars[0].position.x,tests[i].answer.stars[0].position.y,tests[i].answer.stars[0].mass)

    } else {
      t.Errorf("Error! For input test dataset %d, wrong star information for star 0", i)
    }
    if outcome.stars[1].position.x == tests[i].answer.stars[1].position.x && outcome.stars[1].position.y == tests[i].answer.stars[1].position.y && outcome.stars[1].mass == tests[i].answer.stars[1].mass{
      fmt.Println("Correct! The star 1 is", tests[i].answer.stars[1].position.x,tests[i].answer.stars[1].position.y,tests[i].answer.stars[1].mass)

    } else {
      t.Errorf("Error! For input test dataset %d, wrong star information for star 1", i)
    }
    }

  }
