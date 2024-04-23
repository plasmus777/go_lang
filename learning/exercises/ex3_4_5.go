package exercises

import (
	"fmt"

	"github.com/plasmus777/go_lang/learning/exercises/model"
)

// Create a model for item purchases containing: the date of ocurrence, the market's name and the bought items list.
// Have the model on a specific "model" package and create a function to initialize a model and return its related struct using pointers.
func Ex3_4_5() {
	// purchase, error := model.InitializePurchase("")
	purchase, error := model.InitializePurchase("Basic Market")
	if error != nil {
		fmt.Println("Error:", error)
	} else {
		purchase.AddItem2("Apple (1 pound)", 3, 1.29)
		purchase.AddItem2("Pear (1 pound)", 2, 1.34)
		purchase.AddItem2("Car Battery Charger", 1, 20.59)
		purchase.AddItem2("Indoor Fly Trap", 1, 14.91)

		fmt.Println("Market:", purchase.Market)
		fmt.Println("Time of the purchase:", purchase.PurchaseTime)
		fmt.Println("-------- Bought Items --------")
		purchase.ListItems()
	}
}
