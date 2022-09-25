package sample

import (
    "geek-graduate/pb"
    "github.com/google/uuid"
    "math/rand"
)

// 随机键盘布局
func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

// 随机cpu
func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

// 随机cpuName
func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i7-9750H",
			"Core i5-9400F",
			"Core i3-1005G1",
		)
	}
	return randomStringFromSet(
		"Ryzen 7 PRO 2700U",
		"Ryzen 5 PRO 3500U",
		"Ryzen 3 PRO 3200GE",
	)

}

func randomGPUBrand() string {
	return randomStringFromSet("Nvidia", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "Nvidia" {
		return randomStringFromSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1660-Ti",
			"GTX 1070",
		)
	}

	return randomStringFromSet(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
		"RX Vega-56",
	)
}

// 根据给定字符集返回一个随机字符串
func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

// 随机布尔
func randomBool() bool {
	return rand.Intn(2) == 1
}

// 随机范围整数
func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

// 随机范围float64浮点数
func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFLoat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	resulution := &pb.Screen_Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}

	return resulution
}

func randomScreenPanel() pb.Screen_Panel {
	switch rand.Intn(2) {
	case 1:
		return pb.Screen_IPS
	default:
		return pb.Screen_OLED
	}
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "Lenovo")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook Air", "Macbook Pro")
	case "Dell":
		return randomStringFromSet("Latitude", "Vostro", "XPS", "Alienware")
	default:
		return randomStringFromSet("Thinkpad X1", "Thinkpad P1", "Thinkpad P53")
	}
}

func randomID() string {
	return uuid.New().String()
}

func RandomLaptopScore() float64 {
	return float64(randomInt(1, 10))
}
