package Convertions

import "encoding/json"

func MapMorse() map[string]string {

	var rep map[string]string
	m := `{"A": "°- ","B": "-°°° ","C": "-°-° ","D": "-°° ","E": "° ","F": "°°-° ","G": "--° ","H": "°°°° ","J": "°--- ","K": "-°- ","L": "°-°° ", "M": "-- ","N": "-° ","O": "--- ","P": "°--° ","Q": "--°- ","R": "°-° ","S": "°°° ","T": "- ","U ": "°°- ","V": "°°°- ","W": "°-- ","X": "-°°- ","Y": "-°-- ","Z": "--°° ", " ":"   " }`
	json.Unmarshal([]byte(m), &rep)
	return rep
}
