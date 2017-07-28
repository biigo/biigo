package bgoutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectionDiffer(t *testing.T) {
	type aItem struct {
		name string
	}
	type bItem struct {
		key string
	}

	from := []aItem{
		aItem{name: "item1"},
		aItem{name: "item2"},
	}
	to := []bItem{
		bItem{key: "item1"},
		bItem{key: "item3"},
	}

	differ := CollectionDiffer{
		From: func(hander func(interface{}, string)) {
			for _, item := range from {
				hander(item, item.name)
			}
		},
		To: func(hander func(interface{}, string)) {
			for _, item := range to {
				hander(item, item.key)
			}
		},
	}

	added := []string{}
	del := []string{}
	merge := [][]string{}

	differ.Diff(
		func(item interface{}) error {
			added = append(added, item.(aItem).name)
			return nil
		},
		func(item interface{}) error {
			del = append(del, item.(bItem).key)
			return nil
		},
		func(a interface{}, b interface{}) error {
			merge = append(merge, []string{
				a.(aItem).name,
				b.(bItem).key,
			})
			return nil
		},
	)

	assert.Equal(t, []string{"item2"}, added)
	assert.Equal(t, []string{"item3"}, del)
	assert.Equal(t, [][]string{[]string{"item1", "item1"}}, merge)
}
