package solution

import "github.com/kyokomi/emoji"

// GetMessage return greeting string
func GetMessage() string {
	return emoji.Sprint("Hello :world_map:!")
}
