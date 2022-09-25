package sample

import (
	"geek-graduate/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 生成随机键盘
func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

// 生成随机cpu
func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)

	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 8.0)

	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}

	return cpu
}

// 生成随机gpu
func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)

	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)
	memGB := randomInt(2, 6)

	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: &pb.Memory{
			Value: uint64(memGB),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
	return gpu
}

// 生成随机内存
func NewRAM() *pb.Memory {
	memGB := randomInt(4, 64)
	memory := &pb.Memory{
		Value: uint64(memGB),
		Unit:  pb.Memory_GIGABYTE,
	}
	return memory
}

// 生成ssd硬盘
func NewSSD() *pb.Storage {
	memGB := randomInt(128, 1024)
	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memoery: &pb.Memory{
			Value: uint64(memGB),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
	return ssd
}

// 生成ssd硬盘
func NewHDD() *pb.Storage {
	memTB := randomInt(1, 6)
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memoery: &pb.Memory{
			Value: uint64(memTB),
			Unit:  pb.Memory_TERABYTE,
		},
	}
	return hdd
}

// 生成屏幕
func NewScreen() *pb.Screen {
	screen := &pb.Screen{
		SizeInch:   randomFLoat32(13, 17),
		Resolution: randomResolution(),
		Panel:      randomScreenPanel(),
	}
	return screen
}

// 生成电脑
func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &pb.Laptop{
		Id:       randomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Gpu:      []*pb.GPU{NewGPU()},
		Ram:      NewRAM(),
		Storage:  []*pb.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3500),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdatedAt:   timestamppb.Now(),
	}
	return laptop
}
