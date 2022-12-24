// In Go lang, a map is a data structure that stores key-value pairs where keys must be unique but values can be duplicate.
// Maps provide easy retrieval of values.
/*
Declare Maps
Maps can be declared using var or shorthand syntax.

syntax:
// Using var keyword
var mymap = map[TypeOfKey] TypeOfValue { key1:value1, key2:value2,...}

// shorthand syntax
mymap :=  map[TypeOfKey] TypeOfValue { key1:value1, key2:value2,...}

Valid Types for Map
The data type of a key element should be a type that can be compared using the == operator.

Valid Key Types: Numbers, string, Boolean, arrays, pointers, struct

Invalid Key Types: functions, slices, and maps

The type of value can be of any type.

Access Maps
The values in a map can be retrieved by passing the key in the square bracket, e.g. map[key].
*/