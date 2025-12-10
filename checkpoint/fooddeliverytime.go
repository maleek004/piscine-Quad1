package checkpoint

import "strings"

type food struct {
	preptime int
}

// i have been given a struct (a class) called food with a single field 'preptime' of int data type
// i also know that i have to create a map (a dictionary / a list of key:value pairs) of string key and food (the struct we created) values
//// which is defned by this line `menu := map[string]food{...}`
// my first question was that why can't we just use a value of int type for the map and define them like this for example "bueger": 15, "chips":10...
//// and avoid having to create the struct all together???
// my second question was that can struct hold multiple ields of different data types ??
// I also notice that as soon as we find an item that does not exist in our menu , the whole proram stops instead of skipping and trying to calculate the preptime for the rest
// I also noticed that using `f, exists` when checking for an item in a map is very clean and idiomatic because you can just do `if exists` or `!exists`
//// later on which is super clean, i guess it could be even better like this `key , exists := my_map[item]` , correct me if this is wrong

// ANSWERS!!!!!
// the struct was choosen as the values of the struct for scalability and expressiveness purpose, think about adding 'cost' or 'popularity score' for each
//// item in the map later on , we will not have to rebuild it , just add them to the struct and you are good to go...
//// on the expressiveness side , when we write burger.preptime or burger.cost, it is immediately clear which 'attribute' of the struct's item we are retrieveing
//// I also observed that having a map of struct values looks like creating a nested dictionary in python
// About `f,exists`, yes , i can even write it like `data, exists` or `found, exists` etc
// when a field you are requesting is empty in a map of structs, you get the zero value of that field's data type
// when declaring the items within a map, you need to end the list with a trailing comma, unlike in python where there is no need for that??

func FoodDeliveryTime(order string) int {

	menu := map[string]food{
		"burger":  {preptime: 15},
		"chips":   {preptime: 10},
		"nuggets": {preptime: 12},
	}
	totalTime := 0
	items := strings.Split(order, " ")
	for _, item := range items {
		f, exists := menu[item]
		if !exists {
			return 404
		}
		totalTime += f.preptime

	}
	return totalTime

}
