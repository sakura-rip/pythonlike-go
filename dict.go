package pythonlike_go

//List is python like object dictionary in Golang
type Dict map[interface{}]interface{}

func NewDict() Dict {
	return Dict{}
}

//Return the number of items in the dictionary d.
func (d Dict) Length() int {
	return len(d)
}

//Return a new view of the dictionary’s keys
func (d Dict) Keys() []interface{} {
	var keys []interface{}
	for k := range d {
		keys = append(keys, k)
	}
	return keys
}

//Remove all items from the dictionary
func (d *Dict) Clear() {
	*d = NewDict()
}

//Check is the key exist
func (d Dict) IsKeyExist(key interface{}) bool {
	_, ok := d[key]
	return ok
}

//Return the value for key if key is in the dictionary, else default.
//TODO:dont require dv
func (d Dict) Get(key, dv interface{}) interface{} {
	val, ok := d[key]
	if !ok {
		return dv
	}
	return val
}

//If key is in the dictionary, remove it and return its value, else return default.
//TODO:dont require dv
func (d Dict) Pop(key, dv interface{}) interface{} {
	val, ok := d[key]
	if ok {
		delete(d, key)
		return val
	}
	return dv
}

//Remove and return a (key, value) pair from the dictionary.
func (d Dict) PopItem() (interface{}, interface{}) {
	for k, v := range d {
		delete(d, k)
		return k, v
	}
	return nil, nil
}

//If key is in the dictionary, return its value.
//If not, insert key with a value of default and return default. default defaults to Nil.
//TODO:dont require value
func (d Dict) SetDefault(key, value interface{}) interface{} {
	if d.IsKeyExist(key) {
		return d[key]
	}
	d[key] = value
	return value
}

//Update the dictionary with the key/value pairs from other, overwriting existing keys.
//update() accepts either another dictionary object
func (d Dict) Update(map2 Dict) {
	for k, v := range map2 {
		d[k] = v
	}
}

//Return a new view of the dictionary’s values
func (d Dict) Values() []interface{} {
	var slice []interface{}
	for _, v := range d {
		slice = append(slice, v)
	}
	return slice
}
