package allinone

import "encoding/json"

func ToJsonString(v interface{}) (string, error) {
	ret, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(ret), nil
}

func UncheckToJsonString(v interface{}) string {
	ret, err := ToJsonString(v)
	if err != nil {
		panic(err)
	}
	return ret
}

func ToJsonStringPretty(v interface{}) (string, error) {
	ret, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return string(ret), nil
}

func UncheckToJsonStringPretty(v interface{}) string {
	ret, err := ToJsonStringPretty(v)
	if err != nil {
		panic(err)
	}
	return ret
}
