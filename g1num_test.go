package g1num_test

import (
	"testing"

	g1num "github.com/Gnar1337/g1-num"
)

func TestG1Create(t *testing.T) {
	g1 := g1num.InitG1(36, 37)
	// g1.ShowAnts()
	if g1.GetG1() == 36 {
	} else {
		t.Fail()
	}
	// fmt.Print(g1.ShowAnts())
	// fmt.Print(g1.ShowKids())
}

func TestAntLength(t *testing.T) {
	g1 := g1num.InitG1(36, 37)
	// g1.ShowAnts()
	if len(g1.GetAnts()) == 9 {
	} else {
		t.Fail()
	}
	// fmt.Print(g1.GetAnts())
	// fmt.Print(g1.ShowKids())
}

func TestAKidLength(t *testing.T) {
	g1 := g1num.InitG1(36, 37)
	// g1.ShowAnts()
	if len(g1.GetKids()) == 35 {
	} else {
		t.Fail()
	}
	// fmt.Print(g1.GetKids())
	// fmt.Print(g1.ShowKids())
}
