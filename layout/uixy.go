package layout

func getItems() map[string][]string {
	//x_y_layer_hidden

	itemsLayout := map[string][]string{
		"HP0": {"421_441_0_0", "452_441_0_0", "483_441_0_0", "512_441_0_0",
			"471_254_1_1", "499_254_1_1", "527_254_1_1", "555_254_1_1", "583_254_1_1"},
		"head":  {"531_0_0_1"},
		"futou": {"413_50_0_1"},
		"body":  {"528_67_0_1"},
	}
	return itemsLayout
}
