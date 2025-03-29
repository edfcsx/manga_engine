package keyboard

import (
	"github.com/veandco/go-sdl2/sdl"
)

var keyMap = map[int]bool{}

func RegisterKeyPressed(code int) {
	keyMap[code] = true
}

func RegisterKeyReleased(code int) {
	keyMap[code] = false
}

func IsKeyPressed(code int) bool {
	return keyMap[code]
}

func IsAnyKeyPressed(codes []int) bool {
	for _, code := range codes {
		if IsKeyPressed(code) {
			return true
		}
	}

	return false
}

func IsAllKeysPressed(codes []int) bool {
	for _, code := range codes {
		if !IsKeyPressed(code) {
			return false
		}
	}

	return true
}

func GetKeyCode(key string) int {
	code := sdl.GetKeyFromName(key)
	return int(code)
}
