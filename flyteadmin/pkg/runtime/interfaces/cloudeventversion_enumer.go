// Code generated by "enumer -type=CloudEventVersion -json -yaml -trimprefix=CloudEventVersion"; DO NOT EDIT.

package interfaces

import (
	"encoding/json"
	"fmt"
)

const _CloudEventVersionName = "v1v2"

var _CloudEventVersionIndex = [...]uint8{0, 2, 4}

func (i CloudEventVersion) String() string {
	if i >= CloudEventVersion(len(_CloudEventVersionIndex)-1) {
		return fmt.Sprintf("CloudEventVersion(%d)", i)
	}
	return _CloudEventVersionName[_CloudEventVersionIndex[i]:_CloudEventVersionIndex[i+1]]
}

var _CloudEventVersionValues = []CloudEventVersion{0, 1}

var _CloudEventVersionNameToValueMap = map[string]CloudEventVersion{
	_CloudEventVersionName[0:2]: 0,
	_CloudEventVersionName[2:4]: 1,
}

// CloudEventVersionString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CloudEventVersionString(s string) (CloudEventVersion, error) {
	if val, ok := _CloudEventVersionNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to CloudEventVersion values", s)
}

// CloudEventVersionValues returns all values of the enum
func CloudEventVersionValues() []CloudEventVersion {
	return _CloudEventVersionValues
}

// IsACloudEventVersion returns "true" if the value is listed in the enum definition. "false" otherwise
func (i CloudEventVersion) IsACloudEventVersion() bool {
	for _, v := range _CloudEventVersionValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for CloudEventVersion
func (i CloudEventVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for CloudEventVersion
func (i *CloudEventVersion) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("CloudEventVersion should be a string, got %s", data)
	}

	var err error
	*i, err = CloudEventVersionString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for CloudEventVersion
func (i CloudEventVersion) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for CloudEventVersion
func (i *CloudEventVersion) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = CloudEventVersionString(s)
	return err
}