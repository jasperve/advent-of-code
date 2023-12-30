package intcomputerday23

import (
	"errors"
)

type Base struct {
	Mem       []int
	In        <-chan int
	Out       chan<- int
	Ip        int
	Rb        int
	InCounter int
}

type Computer interface {
	Run() error
	SetMem(mem []int)
	IncrementInCounter()
	ResetInCounter()
	GetInCount() int
	Add(opcode int, modes []int) error
	Multiply(opcode int, modes []int) error
	Input(opcode int, modes []int) error
	Output(opcode int, modes []int) error
	JumpTrue(opcode int, modes []int) error
	JumpFalse(opcode int, modes []int) error
	Less(opcode int, modes []int) error
	Equals(opcode int, modes []int) error
	Offset(opcode int, modes []int) error
}

func (b *Base) IncrementInCounter() {
	b.InCounter++
}

func (b *Base) ResetInCounter() {
	b.InCounter = 0
}

func (b *Base) GetInCount() int {
	return b.InCounter
}

func (b *Base) SetMem(mem []int) {
	b.Mem = mem
}

func (b *Base) Run() (err error) {

	b.Ip, b.Rb = 0, 0

	for b.Mem[b.Ip] != 99 || b.Mem[b.Ip] == 0 {

		opcode := b.Mem[b.Ip] % 100
		modes := []int{
			0,
			b.Mem[b.Ip] / 100 % 10,
			b.Mem[b.Ip] / 1000 % 10,
			b.Mem[b.Ip] / 10000 % 10,
		}

		switch opcode {
		case 1:
			err = b.Add(opcode, modes)
		case 2:
			err = b.Multiply(opcode, modes)
		case 3:
			err = b.Input(opcode, modes)
		case 4:
			err = b.Output(opcode, modes)
		case 5:
			err = b.JumpTrue(opcode, modes)
		case 6:
			err = b.JumpFalse(opcode, modes)
		case 7:
			err = b.Less(opcode, modes)
		case 8:
			err = b.Equals(opcode, modes)
		case 9:
			err = b.Offset(opcode, modes)
		default:
			err = errors.New("Invalid opcode")
		}

		if err != nil {
			break
		}

	}

	return err

}

func (b *Base) GetParams(opcode int, modes []int, amount int) (params []*int, err error) {

	params = make([]*int, amount+1, amount+1)
	for i := 1; i <= amount; i++ {

		if b.Ip+i >= 0 && b.Ip+i < len(b.Mem) {

			switch modes[i] {
			case 0:
				if b.Mem[b.Ip+i] >= 0 && b.Mem[b.Ip+i] < len(b.Mem) {
					params[i] = &b.Mem[b.Mem[b.Ip+i]]
				} else if b.Mem[b.Ip+i] >= 0 {
					for len(b.Mem) <= b.Mem[b.Ip+i] {
						b.Mem = append(b.Mem, 0)
					}
					params[i] = &b.Mem[b.Mem[b.Ip+i]]
				} else {
					err = errors.New("Invalid opcode")
				}
			case 1:
				params[i] = &b.Mem[b.Ip+i]
			case 2:
				if b.Mem[b.Ip+i]+b.Rb >= 0 && b.Mem[b.Ip+i]+b.Rb < len(b.Mem) {
					params[i] = &b.Mem[b.Mem[b.Ip+i]+b.Rb]
				} else if b.Mem[b.Ip+i]+b.Rb >= 0 {
					for len(b.Mem) <= b.Mem[b.Ip+i]+b.Rb {
						b.Mem = append(b.Mem, 0)
					}
					params[i] = &b.Mem[b.Mem[b.Ip+i]+b.Rb]
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

func (b *Base) Add(opcode int, modes []int) (err error) {
	params, err := b.GetParams(opcode, modes, 3)
	if err == nil {
		*params[3] = *params[1] + *params[2]
		b.Ip += 4
	}
	return err
}

func (b *Base) Multiply(opcode int, modes []int) (err error) {
	params, err := b.GetParams(opcode, modes, 3)
	if err == nil {
		*params[3] = *params[1] * *params[2]
		b.Ip += 4
	}
	return err
}

func (b *Base) Input(opcode int, modes []int) (err error) {

	params, err := b.GetParams(opcode, modes, 1)
	if err == nil {
		*params[1] = <-b.In
		b.Ip += 2
	}
	return err

}

func (b *Base) Output(opcode int, modes []int) (err error) {

	params, err := b.GetParams(opcode, modes, 1)
	if err == nil {
		b.Out <- *params[1]
		b.Ip += 2
	}
	return err

}

func (b *Base) JumpTrue(opcode int, modes []int) (err error) {

	params, err := b.GetParams(opcode, modes, 2)
	if err == nil {
		if *params[1] != 0 {
			b.Ip = *params[2]
		} else {
			b.Ip += 3
		}
	}
	return err

}

func (b *Base) JumpFalse(opcode int, modes []int) (err error) {

	params, err := b.GetParams(opcode, modes, 2)
	if err == nil {
		if *params[1] == 0 {
			b.Ip = *params[2]
		} else {
			b.Ip += 3
		}
	}
	return err

}

func (b *Base) Less(opcode int, modes []int) (err error) {

	params, err := b.GetParams(opcode, modes, 3)
	if err == nil {
		if *params[1] < *params[2] {
			*params[3] = 1
		} else {
			*params[3] = 0
		}
		b.Ip += 4
	}
	return err

}

func (b *Base) Equals(opcode int, modes []int) (err error) {

	params, err := b.GetParams(opcode, modes, 3)
	if err == nil {
		if *params[1] == *params[2] {
			*params[3] = 1
		} else {
			*params[3] = 0
		}
		b.Ip += 4
	}
	return err

}

func (b *Base) Offset(opcode int, modes []int) (err error) {

	params, err := b.GetParams(opcode, modes, 1)
	if err == nil {
		b.Rb += *params[1]
		b.Ip += 2
	}
	return err

}
