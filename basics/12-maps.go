package basics

import "fmt"

func main() {
	// declaration
	// 1. var mapName map[keyType]valueType
	// 2. var mapName := make(map[keyType]valueType)
	// 3. var mapName := map[keyType]valueType{
	// 	key1: value1,
	// 	key2: value2,
	// }

	mpp := make(map[string]int)
	fmt.Println(mpp)

	// to insert
	mpp["k1"] = 1
	mpp["k2"] = 2

	// to delete
	delete(mpp, "k2")

	// to empty the map
	// clear(mpp)

	fmt.Println(mpp)

	// to check if there is value associated with a key -> myMap["key"] return two things, first is value, second is true/false depending on if its prosent or not
	_, isPrsent := mpp["k1"]
	fmt.Println(isPrsent)

	// there is also nested maps

	// for range loop works on this too
}
