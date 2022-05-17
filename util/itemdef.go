package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Items map[int]Item

type Item struct {
	ItemID      int    `json:"itemID"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Category    string `json:"category"`
	GameValue   int    `json:"gameValue"`
}

func LoadItemDef(file string) (Items, error) {

	// Open our jsonFile
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	// we initialize our Item array
	var items []Item

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &items)

	// Let's now build a map of the items so we can easily access items by their itemid/index
	itemDefMap := map[int]Item{}
	for _, v := range items {
		itemDefMap[v.ItemID] = v
	}

	return itemDefMap, nil
}
