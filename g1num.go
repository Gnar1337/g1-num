package g1num

import (
	"fmt"
	"sync"
)

type g1num uint64

func (g1 *g1num) AmIPrime() bool {
	facts := make([]g1num, 0)
	for i := 1; i <= g1.GetInt()/2; i++ {
		if g1.GetInt()%i == 0 {
			facts = append(facts, g1num(i))
		}
	}
	if len(facts) == 0 {
		return true
	} else {
		return false
	}
}

type factors []g1num
type mults []g1num

type g1 struct {
	me          g1num
	isprime     bool
	ants        factors
	kids        mults
	primes      []g1num
	palindromes []g1num
	versitility uint64
	ig1num      g1num
	wg          *sync.WaitGroup
}

func InitG1(num, v int) *g1 {
	g1 := &g1{
		me:          CreateG1Num(num),
		isprime:     false,
		ants:        factors{},
		kids:        mults{},
		primes:      []g1num{},
		palindromes: []g1num{},
		ig1num:      g1num(num * 37037 * 3),
		versitility: uint64(v),
		wg:          &sync.WaitGroup{},
	}
	// g1.wg.Add()
	// wg.Add(2)
	fmt.Println("getting factors")
	g1.wg.Add(1)
	go g1.GenerateFactors()
	g1.wg.Add(1)
	go g1.GenerateMultiples()
	g1.wg.Wait()
	// fmt.Print(g1.palindromes)
	return g1
}

func CreateG1Num(num int) g1num {
	return g1num(num)
}

func (g *g1num) Printme() {
	fmt.Printf("g1-num = %d", *g)
}

func (g *g1num) GetInt() int {
	return int(*g)

}

func (g *g1) GetG1() int {
	return g.me.GetInt()
}

func (g *g1) GetAnts() []g1num {
	a := []g1num(g.ants)
	return a
}
func (g *g1) GetKids() []g1num {
	a := []g1num(g.kids)
	return a
}

func (g *g1) ShowAnts() string {
	return fmt.Sprintf("Ancestors: %v", g.ants)
}
func (g *g1) ShowKids() string {
	return fmt.Sprintf("kids: %v", g.kids)
}

func (g *g1) ShowPalindromes() string {
	return fmt.Sprintf("Palindromes: %v", g.palindromes)
}
func (g *g1) ShowIG1Num() string {
	return fmt.Sprintf("Improved Number: %v", g.ig1num)
}

func (g *g1) GenerateFactors() {
	// g.wg.Add(1)
	facts := make([]g1num, 0)
	for i := 1; i <= g.GetG1()/2; i++ {
		if g.GetG1()%i == 0 {
			facts = append(facts, g1num(i))
			if isPalindrome(i) {
				ant := g1num(i)
				g.palindromes = append(g.palindromes, ant)
			}

		}
	}
	facts = append(facts, g1num(g.GetG1()))
	g.ants = facts
	if len(facts) == 2 {
		g.isprime = true
	}
	g.wg.Done()
}

func (g *g1) GenerateMultiples() {
	// g.wg.Add(1)
	mults := make([]g1num, 0)
	for i := 2; i <= g.GetG1(); i++ {
		kid := g1num(i * g.GetG1())
		mults = append(mults, kid)
		if isPalindrome(int(kid)) {
			g.palindromes = append(g.palindromes, kid)
		}

	}
	g.kids = mults
	g.wg.Done()
}

func isPalindrome(number int) bool {
	original := number
	reversed := 0

	// Reverse the number
	for number > 0 {
		remainder := number % 10
		reversed = reversed*10 + remainder
		number = number / 10
	}

	// Compare the original number with the reversed number
	return original == reversed
}
