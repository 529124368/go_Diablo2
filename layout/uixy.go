package layout

func getItems() map[string][]string {
	//x_y_layer_hidden

	itemsLayout := map[string][]string{
		"HP0": {"421_441_0_3", "452_441_0_3", "483_441_0_3", "512_441_0_3",
			"471_254_1_1", "499_254_1_1", "527_254_1_1", "555_254_1_1", "583_254_1_1"},
		"book":   {"615_254_1_1"},
		"head-3": {"531_0_0_1"},
		"sword":  {"413_50_0_1"},
		"body-3": {"528_70_0_1"},
		"dun":    {"645_67_0_1"},
	}
	return itemsLayout
}
