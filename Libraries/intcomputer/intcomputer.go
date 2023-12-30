package intcomputer

import (
	"errors"
)

type Computer struct {
	Mem    []int
	In     <-chan int
	Out    chan<- int
	ip     int
	rb     int
	status int
}

func (c *Computer) Run() (err error) {

	c.ip, c.rb = 0, 0

	for c.Mem[c.ip] != 99 || c.Mem[c.ip] == 0 {

		opcode := c.Mem[c.ip] % 100
		modes := []int{
			0,
			c.Mem[c.ip] / 100 % 10,
			c.Mem[c.ip] / 1000 % 10,
			c.Mem[c.ip] / 10000 % 10,
		}

		switch opcode {
		case 1:
			err = c.add(opcode, modes)
		case 2:
			err = c.multiply(opcode, modes)
		case 3:
			err = c.input(opcode, modes)
		case 4:
			err = c.output(opcode, modes)
		case 5:
			err = c.jumpTrue(opcode, modes)
		case 6:
			err = c.jumpFalse(opcode, modes)
		case 7:
			err = c.less(opcode, modes)
		case 8:
			err = c.equals(opcode, modes)
		case 9:
			err = c.offset(opcode, modes)
		default:
			err = errors.New("Invalid opcode")
		}

		if err != nil {
			break
		}

	}

	return err

}

func (c *Computer) getParams(opcode int, modes []int, amount int) (params []*int, err error) {

	params = make([]*int, amount+1, amount+1)
	for i := 1; i <= amount; i++ {

		if c.ip+i >= 0 && c.ip+i < len(c.Mem) {

			switch modes[i] {
			case 0:
				if c.Mem[c.ip+i] >= 0 && c.Mem[c.ip+i] < len(c.Mem) {
					params[i] = &c.Mem[c.Mem[c.ip+i]]
				} else if c.Mem[c.ip+i] >= 0 {
					for len(c.Mem) <= c.Mem[c.ip+i] {
						c.Mem = append(c.Mem, 0)
					}
					params[i] = &c.Mem[c.Mem[c.ip+i]]
				} else {
					err = errors.New("Invalid opcode")
				}
			case 1:
				params[i] = &c.Mem[c.ip+i]
			case 2:
				if c.Mem[c.ip+i]+c.rb >= 0 && c.Mem[c.ip+i]+c.rb < len(c.Mem) {
					params[i] = &c.Mem[c.Mem[c.ip+i]+c.rb]
				} else if c.Mem[c.ip+i]+c.rb >= 0 {
					for len(c.Mem) <= c.Mem[c.ip+i]+c.rb {
						c.Mem = append(c.Mem, 0)
					}
					params[i] = &c.Mem[c.Mem[c.ip+i]+c.rb]
				} else {
					err = errors.New("Invalid opcode")
				}
			}

		} else {
			err = errors.New("Invalid opcode")
		}
	}

	return params, err

}

func (c *Computer) add(opcode int, modes []int) (err error) {
	params, err := c.getParams(opcode, modes, 3)
	if err == nil {
		*params[3] = *params[1] + *params[2]
		c.ip += 4
	}
	return err
}

func (c *Computer) multiply(opcode int, modes []int) (err error) {
	params, err := c.getParams(opcode, modes, 3)
	if err == nil {
		*params[3] = *params[1] * *params[2]
		c.ip += 4
	}
	return err
}

func (c *Computer) input(opcode int, modes []int) (err error) {

	params, err := c.getParams(opcode, modes, 1)
	if err == nil {
		*params[1] = <-c.In
		c.ip += 2
	}
	return err

}

func (c *Computer) output(opcode int, modes []int) (err error) {

	params, err := c.getParams(opcode, modes, 1)
	if err == nil {
		c.Out <- *params[1]
		c.ip += 2
	}
	return err

}

func (c *Computer) jumpTrue(opcode int, modes []int) (err error) {

	params, err := c.getParams(opcode, modes, 2)
	if err == nil {
		if *params[1] != 0 {
			c.ip = *params[2]
		} else {
			c.ip += 3
		}
	}
	return err

}

func (c *Computer) jumpFalse(opcode int, modes []int) (err error) {

	params, err := c.getParams(opcode, modes, 2)
	if err == nil {
		if *params[1] == 0 {
			c.ip = *params[2]
		} else {
			c.ip += 3
		}
	}
	return err

}

func (c *Computer) less(opcode int, modes []int) (err error) {

	params, err := c.getParams(opcode, modes, 3)
	if err == nil {
		if *params[1] < *params[2] {
			*params[3] = 1
		} else {
			*params[3] = 0
		}
		c.ip += 4
	}
	return err

}

func (c *Computer) equals(opcode int, modes []int) (err error) {

	params, err := c.getParams(opcode, modes, 3)
	if err == nil {
		if *params[1] == *params[2] {
			*params[3] = 1
		} else {
			*params[3] = 0
		}
		c.ip += 4
	}
	return err

}

func (c *Computer) offset(opcode int, modes []int) (err error) {

	params, err := c.getParams(opcode, modes, 1)
	if err == nil {
		c.rb += *params[1]
		c.ip += 2
	}
	return err

}
