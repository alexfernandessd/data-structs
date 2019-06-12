package main

import (
	"fmt"
	"strconv"
)

type pilha struct {
	pilha []int
	top   int
}

func newPilha() *pilha {
	return &pilha{
		pilha: []int{},
		top:   -1,
	}
}

func convert(n int) string {
	p := newPilha()
	for n > 0 {
		p.push(n % 2)
		n = n / 2
	}
	str := ""

	for _, n := range p.pilha {
		str += strconv.Itoa(n)
	}
	return str
}

func main() {
	p := newPilha()
	for i := 1; i < 11; i++ {
		p.push(i)
	}
	fmt.Println("Simple push", p.pilha)
	str := convert(13)
	fmt.Println("Convert number to binary", str)
	p.pushSpecific(5, 910)
	fmt.Println("Push with specific index", p.pilha)

	p.popSpecific(5)
	fmt.Println("Pop specific", p.pilha)
}

func (p *pilha) push(element int) {
	p.pilha = append(p.pilha, element)
}

func (p *pilha) pushSpecific(i, element int) {
	for lastIndex := len(p.pilha) - 1; lastIndex >= i; lastIndex-- {
		if lastIndex == len(p.pilha)-1 {
			p.pilha = append(p.pilha, p.pilha[lastIndex])
		} else {
			p.pilha[lastIndex+1] = p.pilha[lastIndex]
		}
	}
	p.pilha[i] = element
}

func (p *pilha) popSpecific(i int) {
	for j := 0; j < len(p.pilha)-1; j++ {
		if j < i {
			continue
		}
		p.pilha[j] = p.pilha[j+1]
	}
	p.pilha = p.pilha[:len(p.pilha)-1]
}

func (p *pilha) pop() int {
	item := p.pilha[len(p.pilha)-1]
	p.pilha = p.pilha[:len(p.pilha)-1]
	return item
}

func (p *pilha) peek() int {
	return p.pilha[len(p.pilha)-1]
}

// func (p *pilha) isFull() bool {
// 	return len(p.pilha) == p.top
// }

func (p *pilha) isEmpty() bool {
	return len(p.pilha) == 0
}
