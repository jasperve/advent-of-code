package transpiler

import (
	"errors"
	"fmt"
)

type Transpiler struct {
	Mem map[int]*int
	In <-chan int
	Out chan<- int
	ip int
	rb int
}

func (c Transpiler) Run() (err error) {

	c.ip, c.rb = 0, 0

	for *c.Mem[c.ip] != 99 || *c.Mem[c.ip] == 0 {

		fmt.Printf("ip: %v rb: %v --- ", c.ip, c.rb)

		opcode := *c.Mem[c.ip] % 100
		modes := []int{
			0,
			*c.Mem[c.ip] / 100 % 10,
			*c.Mem[c.ip] / 1000 % 10,
			*c.Mem[c.ip] / 10000 % 10,
		}

		switch opcode {
		case 1: err = c.add(opcode, modes)
		case 2: err = c.multiply(opcode, modes)
		case 3: err = c.input(opcode, modes)
		case 4: err = c.output(opcode, modes)
		case 5: err = c.jumpTrue(opcode, modes)
		case 6: err = c.jumpFalse(opcode, modes)
		case 7: err = c.less(opcode, modes)
		case 8: err = c.equals(opcode, modes)
		case 9: err = c.offset(opcode, modes)
		default: err = errors.New("Invalid opcode")
		}

		if err != nil {
			break
		}

	}

	return err

}

func (c *Transpiler) getParams(opcode int, modes []int, amount int) (params []*int, indexes []int, err error) {

	params = make([]*int, amount+1, amount+1)
	indexes = make([]int, amount+1, amount+1)
	for i := 1; i <= amount; i++ {
		
		if iv, ok := c.Mem[c.ip+i]; ok {
		
			switch modes[i] {
			case 0:	
			 	if pv, ok := c.Mem[*iv]; ok {
					params[i] = pv
					indexes[i] = *iv
				} else if *iv > -1 {
					c.Mem[*iv] = new(int)
					params[i] = c.Mem[*iv]
					indexes[i] = *iv
				} else {
					err = errors.New("Invalid opcode")
				}
			case 1: 
				params[i] = iv
				indexes[i] = c.ip+i
			case 2:
				if rv, ok := c.Mem[*iv + c.rb]; ok { 
					params[i] = rv
					indexes[i] = *iv + c.rb
				} else if *iv + c.rb > -1 {
					c.Mem[*iv + c.rb] = new(int)
					params[i] = c.Mem[*iv + c.rb]
					indexes[i] = *iv + c.rb
				} else {
					err = errors.New("Invalid opcode")
				}
			}

		} else {
			err = errors.New("Invalid opcode")
		}
	}

	return params, indexes, err

}

func (c *Transpiler) add(opcode int, modes []int) (err error) {
	params, indexes, err := c.getParams(opcode, modes, 3)
	if err == nil {
		fmt.Printf("%v (mem[%v]) = %v (mem[%v]) + %v (mem[%v])\n", *params[1] + *params[2], indexes[3], *params[1], indexes[1], *params[2], indexes[2])
		*params[3] = *params[1] + *params[2]
		c.ip += 4
	}
	return err
}

func (c *Transpiler) multiply(opcode int, modes []int) (err error) {
	params, indexes, err := c.getParams(opcode, modes, 3)
	if err == nil {
		fmt.Printf("%v (mem[%v]) = %v (mem[%v]) * %v (mem[%v])\n", *params[1] * *params[2], indexes[3], *params[1], indexes[1], *params[2], indexes[2])
		*params[3] = *params[1] * *params[2]
		c.ip += 4
	}
	return err
}

func (c *Transpiler) input(opcode int, modes []int) (err error) {

	params, indexes, err := c.getParams(opcode, modes, 1)
	if err == nil {
		in := <- c.In
		fmt.Printf("%v (mem[%v]) <- %v (input)\n", in, indexes[1], in)
		*params[1] = in
		c.ip += 2
	}
	return err

}

func (c *Transpiler) output(opcode int, modes []int) (err error) {

	params, indexes, err := c.getParams(opcode, modes, 1)
	if err == nil {
		c.Out <- *params[1]
		fmt.Printf("%v (mem[%v]) -> output (%v)\n", *params[1], indexes[1], *params[1])
		c.ip += 2
	}
	return err

}

func (c *Transpiler) jumpTrue(opcode int, modes []int) (err error) {

	params, indexes, err := c.getParams(opcode, modes, 2)
	if err == nil {
		if *params[1] != 0 {
			fmt.Printf("If %v (mem[%v]) != 0 -> instruction pointer jumps to %v (mem[%v])\n", *params[1], indexes[1], *params[2], indexes[2])
			if *params[2] < c.ip {
				fmt.Printf("---------------- LOOP DETECTED (JUMPING BACK TO %v) ----------------\n", *params[2])
			}
			c.ip = *params[2]
		} else { 
			fmt.Printf("If %v (mem[%v]) == 0 -> CONTINUE || ELSE JUMP TO %v mem[%v]\n", *params[1], indexes[1], *params[2], indexes[2])
			c.ip += 3
		}
	}
	return err
	
}

func (c *Transpiler) jumpFalse(opcode int, modes []int) (err error) {

	params, indexes, err := c.getParams(opcode, modes, 2)
	if err == nil {
		if *params[1] == 0 {
			fmt.Printf("If %v (mem[%v]) == 0 -> instruction pointer jumps to %v (mem[%v])\n", *params[1], indexes[1], *params[2], indexes[2])
			if *params[2] < c.ip {
				fmt.Printf("---------------- LOOP DETECTED (JUMPING BACK TO %v) ----------------\n", *params[2])
			}
			c.ip = *params[2]
		} else {
			fmt.Printf("If %v (mem[%v]) != 0 -> CONTINUE || ELSE JUMP TO %v mem[%v]\n", *params[1], indexes[1], *params[2], indexes[2])
			c.ip += 3
		}
	}
	return err

}

func (c *Transpiler) less(opcode int, modes []int) (err error) {

	params, indexes, err := c.getParams(opcode, modes, 3)
	if err == nil {
		if *params[1] < *params[2] {
			fmt.Printf("If %v (mem[%v]) < %v (mem[%v]) -> mem[%v] = 1 || else mem[%v] = 0\n", *params[1], indexes[1], *params[2], indexes[2], indexes[3], indexes[3])
			*params[3] = 1
		} else {
			fmt.Printf("If %v (mem[%v]) >= %v (mem[%v]) -> mem[%v] = 0 || else mem[%v] = 1\n", *params[1], indexes[1], *params[2], indexes[2], indexes[3], indexes[3])
			*params[3] = 0
		}
		c.ip += 4
	}
	return err 

}

func (c *Transpiler) equals(opcode int, modes []int) (err error) {

	params, indexes, err := c.getParams(opcode, modes, 3)
	if err == nil {
		if *params[1] == *params[2] {
			fmt.Printf("If %v (mem[%v]) == %v (mem[%v]) -> mem[%v] = 1 || else mem[%v] = 0\n", *params[1], indexes[1], *params[2], indexes[2], indexes[3], indexes[3])
			*params[3] = 1
		} else {
			fmt.Printf("If %v (mem[%v]) != %v (mem[%v]) -> mem[%v] = 0 || else mem[%v] = 1\n", *params[1], indexes[1], *params[2], indexes[2], indexes[3], indexes[3])
			*params[3] = 0
		}
		c.ip += 4
	}
	return err

}

func (c *Transpiler) offset(opcode int, modes []int) (err error) {

	params, indexes, err := c.getParams(opcode, modes, 1)
	if err == nil {
		fmt.Printf("Change relative offset by %v (mem[%v])\n", *params[1], indexes[1])
		c.rb += *params[1]
		c.ip += 2
	}
	return err

}
