package pythonlike_go

type Map map[interface{}]interface{}

func NewMap() Map {
	return Map{}
}

func (m Map) Length() int {
	return len(m)
}

func (m Map) Keys() []interface{} {
	var keys []interface{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (m *Map) Clear() {
	*m = NewMap()
}

func (m Map) IsKeyExist(key interface{}) bool {
	_, ok := m[key]
	return ok
}

func (m Map) Get(key, dv interface{}) interface{} {
	val, ok := m[key]
	if !ok {
		return dv
	}
	return val
}

func (m Map) Pop(key, dv interface{}) interface{} {
	val, ok := m[key]
	if ok {
		delete(m, key)
		return val
	}
	return dv
}

func (m Map) PopItem() (interface{}, interface{}) {
	for k, v := range m {
		delete(m, k)
		return k, v
	}
	return nil, nil
}

func (m Map) SetDefault(key, value interface{}) interface{} {
	if m.IsKeyExist(key) {
		return m[key]
	}
	m[key] = value
	return value
}

func (m Map) Update(map2 Map) {
	for k, v := range map2 {
		m[k] = v
	}
}

func (m Map) Values() []interface{} {
	var slice []interface{}
	for _, v := range m {
		slice = append(slice, v)
	}
	return slice
}
