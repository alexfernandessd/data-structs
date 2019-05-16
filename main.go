package main

import (
	"fmt"
	"strconv"
)

func main() {
	str, err := convert(13)
	fmt.Println(str, err)
}

type pilha struct {
	arr []int
	top int
	max int
}

func newPilha(max int) *pilha {
	return &pilha{
		arr: make([]int, max),
		max: max,
		top: -1,
	}
}

func convert(n int) (string, error) {
	s := newPilha(10)
	for n > 0 {
		fmt.Println(n % 2)
		err := s.push(n % 2)
		if err != nil {
			return "", err
		}
		n = n / 2
	}
	str := ""
	for !s.isEmpty() {
		i, err := s.pop()
		if err != nil {
			return "", err
		}
		str += strconv.Itoa(i)
	}
	return str, nil
}

func (p *pilha) peek() (int, error) {
	if p.isEmpty() {
		return -1, fmt.Errorf("stack is empty")
	}
	return p.arr[p.top], nil
}

func (p *pilha) push(push int) error {
	if p.isFull() {
		return fmt.Errorf("stack is full")
	}
	p.top++
	p.arr[p.top] = push
	return nil
}

func (p *pilha) pop() (int, error) {
	if p.isEmpty() {
		return -1, fmt.Errorf("stack is empty")
	}
	element := p.arr[p.top]
	p.top--
	return element, nil
}

func (p *pilha) isEmpty() bool {
	return p.top == -1
}

func (p *pilha) isFull() bool {
	return p.top == p.max-1
}
