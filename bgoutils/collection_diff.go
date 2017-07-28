package bgoutils

type collectionDiffItemEachHandler func(interface{}, string)
type collectionDiffItemEacher func(func(interface{}, string))

// CollectionDiffer 提供两个集合对比功能
type CollectionDiffer struct {
	From collectionDiffItemEacher
	To   collectionDiffItemEacher
}

func (differ CollectionDiffer) Diff(
	addHandler func(interface{}) error,
	delhandler func(interface{}) error,
	mergeHandler func(interface{}, interface{}) error,
) error {
	fromMap := map[string]interface{}{}
	differ.From(func(item interface{}, key string) {
		fromMap[key] = item
	})
	toMap := map[string]interface{}{}
	differ.To(func(item interface{}, key string) {
		toMap[key] = item
	})

	for key, item := range fromMap {
		if toItem, ok := toMap[key]; ok {
			if err := mergeHandler(item, toItem); err != nil {
				return err
			}
		} else {
			if err := addHandler(item); err != nil {
				return err
			}
		}
	}
	for key, toItem := range toMap {
		if _, ok := fromMap[key]; !ok {
			if err := delhandler(toItem); err != nil {
				return err
			}
		}
	}

	return nil
}
