package interfaces

type UIInterface interface {
	AddItemToBag(mousex, mousey int, itemName string) bool
	DelItemFromBag(imageX, imageY int)
	JudgeCanToEquip(mousex, mousey int, itemName string) bool
	ClearTempBag() string
	AddItemToBagByHand(x, y int, itemName string)
}
