package utils

type ImmutableMap struct {
	data	map[interface{}]interface{}
}

func NewImmutableMap(data	map[interface{}]interface{}) *ImmutableMap {
	return &ImmutableMap{data: data}
}

func (m *ImmutableMap) Get(key interface{}) interface{} {
	value, ok := m.data[key]
	if !ok {
		return nil
	}

	return value
}

func (m *ImmutableMap) GetAll() map[interface{}]interface{} {
	result := make(map[interface{}]interface{}, len(m.data))

	for k, v := range m.data {
    result[k] = v
	}

	return result
}
