package main
type Gender string

type Sample struct {
	Name   string `default:"John Smith"`
	Age    int    `default:"27"`
	Gender Gender `default:"m"`

	Slice       []string       `default:"[]"`
	SliceByJSON []int          `default:"[1, 2, 3]"` // Supports JSON
	Map         map[string]int `default:"{}"`
	MapByJSON   map[string]int `default:"{\"foo\": 123}"`

	Struct    OtherStruct  `default:"{}"`
	StructPtr *OtherStruct `default:"{\"Foo\": 123}"`

	NoTag  OtherStruct               // Recurses into a nested struct by default
	OptOut OtherStruct `default:"-"` // Opt-out
}

type OtherStruct struct {
	Hello  string `default:"world"` // Tags in a nested struct also work
	Foo    int    `default:"-"`
	Random int    `default:"-"`
}

// SetDefaults implements defaults.Setter interface
func (s *OtherStruct) SetDefaults() {
	if defaults.CanUpdate(s.Random) { // Check if it's a zero value (recommended)
	s.Random = rand.Int() // Set a dynamic value
	}
}
func main(){
	obj := &Sample{}
	if err := defaults.Set(obj); err != nil {
		panic(err)
	}
}
