package emoji

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/omer5985/Go-Emoji-Utils/utils"
	"github.com/omer5985/Go-Emoji-Utils/data"
)

// Emoji - Struct representing Emoji
type Emoji struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Descriptor string `json:"descriptor"`
}

// Emojis - Map of Emoji Runes as Hex keys to their description
var Emojis map[string]Emoji

// Unmarshal the emoji JSON into the Emojis map
func init() {
	// Work out where we are in relation to the caller
	json.Unmarshal([]byte(data.EmojiData), &Emojis)
}

// LookupEmoji - Lookup a single emoji definition
func LookupEmoji(emojiString string) (emoji Emoji, err error) {

	hexKey := utils.StringToHexKey(emojiString)

	// If we have a definition for this string we'll return it,
	// else we'll return an error
	if e, ok := Emojis[hexKey]; ok {
		emoji = e
	} else {
		err = fmt.Errorf("No record for \"%s\" could be found", emojiString)
	}

	return emoji, err
}

// LookupEmojis - Lookup definitions for each emoji in the input
func LookupEmojis(emoji []string) (matches []interface{}) {
	for _, emoji := range emoji {
		if match, err := LookupEmoji(emoji); err == nil {
			matches = append(matches, match)
		} else {
			matches = append(matches, err)
		}
	}

	return
}

// RemoveAll - Remove all emoji
func RemoveAll(input string) string {

	// Find all the emojis in this string
	matches := FindAll(input)

	for _, item := range matches {
		emo := item.Match.(Emoji)
		rs := []rune(emo.Value)
		for _, r := range rs {
			input = strings.ReplaceAll(input, string([]rune{r}), "")
		}
	}

	// Remove and trim and left over whitespace
	return strings.TrimSpace(strings.Join(strings.Fields(input), " "))
	//return input
}
