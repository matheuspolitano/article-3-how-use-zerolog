package utils

type ListKeys []string

func GetKeys(mapIn map[string]string) ListKeys {
	var listKeys = make([]string, 0)
	for key := range mapIn {
		listKeys = append(listKeys, key)
	}
	return listKeys
}
