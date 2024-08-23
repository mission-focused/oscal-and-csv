package common

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// FlattenJSON flattens a nested JSON object into a string.
func FlattenJSON(data map[string]interface{}, prefix string) string {
	var flatStr []string

	for key, value := range data {
		flatKey := key
		if prefix != "" {
			flatKey = prefix + "." + key
		}

		switch v := value.(type) {
		case map[string]interface{}:
			flatStr = append(flatStr, FlattenJSON(v, flatKey))
		case []interface{}:
			for i, item := range v {
				arrayKey := fmt.Sprintf("%s[%d]", flatKey, i)
				if nestedMap, ok := item.(map[string]interface{}); ok {
					flatStr = append(flatStr, FlattenJSON(nestedMap, arrayKey))
				} else {
					flatStr = append(flatStr, fmt.Sprintf("%s=%v", arrayKey, item))
				}
			}
		default:
			flatStr = append(flatStr, fmt.Sprintf("%s=%v", flatKey, v))
		}
	}

	return strings.Join(flatStr, "|")
}

// UnflattenJSON converts a flattened string back to a nested JSON object.
func UnflattenJSON(flatStr string) map[string]interface{} {
	data := make(map[string]interface{})
	entries := strings.Split(flatStr, "|")

	for _, entry := range entries {
		parts := strings.SplitN(entry, "=", 2)
		key, value := parts[0], parts[1]

		setValue(data, key, value)
	}

	return data
}

// setValue sets a value in the nested map based on the flattened key.
func setValue(data map[string]interface{}, key string, value interface{}) {
	if strings.Contains(key, "[") {
		// Handle array key like "user.phones[0]"
		re := regexp.MustCompile(`^([^\[]+)\[(\d+)\]`)
		matches := re.FindStringSubmatch(key)

		if len(matches) == 3 {
			arrayKey := matches[1]
			index, _ := strconv.Atoi(matches[2])

			// Ensure the key exists and is a slice
			if _, ok := data[arrayKey]; !ok {
				data[arrayKey] = make([]interface{}, index+1)
			}

			arr := data[arrayKey].([]interface{})
			if len(arr) <= index {
				newArr := make([]interface{}, index+1)
				copy(newArr, arr)
				arr = newArr
			}

			if strings.Contains(key[len(matches[0]):], ".") {
				nestedMap := make(map[string]interface{})
				if arr[index] != nil {
					nestedMap = arr[index].(map[string]interface{})
				}
				setValue(nestedMap, key[len(matches[0])+1:], value)
				arr[index] = nestedMap
			} else {
				arr[index] = value
			}

			data[arrayKey] = arr
		}
	} else if strings.Contains(key, ".") {
		// Handle nested key like "user.address.city"
		parts := strings.SplitN(key, ".", 2)
		nestedKey := parts[0]
		rest := parts[1]

		if _, ok := data[nestedKey]; !ok {
			data[nestedKey] = make(map[string]interface{})
		}

		nestedMap := data[nestedKey].(map[string]interface{})
		setValue(nestedMap, rest, value)
	} else {
		// Handle simple key
		data[key] = value
	}
}
