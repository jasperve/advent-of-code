package graph

import "reflect"

type dictType map[string]map[interface{}]map[int]interface{}
type dictItemType interface {
	GetID() int
}

func (d dictType) remove(key string, value interface{}, n dictItemType) {
	delete(d[key][value], n.GetID())
}

func (d dictType) add(key string, value interface{}, n dictItemType) {
	if _, ok := d[key]; !ok {
		d[key] = make(map[interface{}]map[int]interface{})
	}
	if _, ok := d[key][value]; !ok {
		d[key][value] = make(map[int]interface{})
	}
	d[key][value][n.GetID()] = n
}

func (d dictType) get(key string, value interface{}) map[int]interface{} {

	if reflect.TypeOf(value).Kind() == reflect.Func {
		combined := map[int]interface{}{}
		for k, v := range d[key] {
			if value.(func(interface{})bool)(k) {
				for ki, vi := range (v) {
					combined[ki] = vi
				}
			}
		}
		return combined
	}

	if v, ok := d[key][value]; ok {
		return v
	}
 	return nil
}
