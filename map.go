package dataparse

type Map map[any]any

func (m Map) Get(key any) (Value, bool) {
	v, ok := m[key]
	if ok {
		return NewValue(v), true
	}
	return NewValue(nil), false
}

func (m Map) MustGet(key any) Value {
	v, _ := m.Get(key)
	return v
}

func (m Map) Map(key any) (Map, bool) {
	v, ok := m[key]
	if ok {
		if typed, ok := v.(map[any]any); ok {
			return Map(typed), true
		}
	}
	return nil, false
}

func (m Map) MustMap(key any) Map {
	v, ok := m.Map(key)
	if ok {
		return v
	}
	return Map{}
}
