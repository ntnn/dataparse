package dataparse

type Map map[any]any

func (m Map) Get(keys ...any) (Value, bool) {
	for _, key := range keys {
		if v, ok := m[key]; ok {
			return NewValue(v), true
		}
	}
	return NewValue(nil), false
}

func (m Map) MustGet(keys ...any) Value {
	v, _ := m.Get(keys...)
	return v
}

func (m Map) Map(keys ...any) (Map, bool) {
	for _, key := range keys {
		if v, ok := m[key]; ok {
			if typed, ok := v.(map[any]any); ok {
				return Map(typed), true
			}
		}
	}
	return Map{}, false
}

func (m Map) MustMap(keys ...any) Map {
	v, _ := m.Map(keys...)
	return v
}
