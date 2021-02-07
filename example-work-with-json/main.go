package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// StudentStruct declares `StudentStruct` structure
type StudentStruct struct {
	FirstName, lastName string
	Email               string
	Age                 int
	HeightInMeters      float64
	IsMale              bool
}

// StudentMap declares `StudentMap` map
type StudentMap map[string]interface{}

// Profile declares `Profile` structure
type Profile struct {
	Username  string
	followers int
	Grades    map[string]string
}

// ProfileI interface defines `ProfileI` method
type ProfileI interface {
	Follow()
}

// ProfileToImplementFollowMethodInProfileInterface declares `ProfileToImplementFollowMethodInProfileInterface` structure
type ProfileToImplementFollowMethodInProfileInterface struct {
	Username  string
	Followers int
}

// StudentAbstractDataType declares `StudentAbstractDataType` structure
type StudentAbstractDataType struct {
	FirstName, lastName string
	Age                 int
	Profile             Profile
	Languages           []string
}

// StudentAbstratDataTypeWithProfileAnonymous declares `StudentAbstratDataTypeWithProfileAnonymous` structure
type StudentAbstratDataTypeWithProfileAnonymous struct {
	FirstName, lastName string
	Age                 int
	Profile
	Languages []string
}

// Follow method implementation
func (p *ProfileToImplementFollowMethodInProfileInterface) Follow() {
	p.Followers++
}

// StudentWithProfileInterface declares `StudentWithProfileInterface` structure
type StudentWithProfileInterface struct {
	FirstName, lastName string
	Age                 int
	Primary             ProfileI
	Secondary           ProfileI
}

// MarshalJSON - implement `Marshaler` interface
func (p ProfileToImplementFollowMethodInProfileInterface) MarshalJSON() ([]byte, error) {
	// return JSON value
	// TODO: handle error gracefully
	return []byte(fmt.Sprintf(`{"f_count": "%d"}`, p.Followers)), nil
}

// MarshalText - implement `TextMarshaler` interface
func (a Age) MarshalText() ([]byte, error) {
	// return string value
	// TODO: handle error gracefully
	return []byte(fmt.Sprintf(`{"age": %d}`, int(a))), nil
}

// Age declares `Age` type
type Age int

// StudentWithFileAge declares `StudentWithFileAge` structure
type StudentWithFileAge struct {
	FirstName, lastName string
	Age                 Age
	Profile             ProfileToImplementFollowMethodInProfileInterface
}

// ProfileWithStructureTags declares `ProfileWithStructureTags` structure
type ProfileWithStructureTags struct {
	Username  string `json:"uname"`
	Followers int    `json:"followers,omitempty,string"`
}

// StudentWithStructureTags declares `StudentWithStructureTags` structure
type StudentWithStructureTags struct {
	FirstName string                   `json:"fname"`           // `fname` as field name
	LastName  string                   `json:"lname,omitempty"` // discard if value is empty
	Email     string                   `json:"-"`               // always discard
	Age       int                      `json:"-,"`              // `-` as field name
	IsMale    bool                     `json:",string"`         // keep original field name, coerce to a string
	Profile   ProfileWithStructureTags `json:""`                // no effect
}

// StudentToDecodeJSON declares `Student` structure
type StudentToDecodeJSON struct {
	FirstName, lastName string
	Email               string
	Age                 int
	HeightInMeters      float64
}

// ProfileComplexData declares `ProfileComplexData` structure
type ProfileComplexData struct {
	Username  string
	Followers int
}

// StudentComplexDataToDecodeJSON declares `StudentComplexDataToDecodeJSON` structure
type StudentComplexDataToDecodeJSON struct {
	FirstName, lastName string
	HeightInMeters      float64
	IsMale              bool
	Languages           [2]string
	Subjects            []string
	Grades              map[string]string
	Profile             ProfileComplexData
}

// StudentComplexDataToDecodeJSONII declares `StudentComplexDataToDecodeJSONII` structure
type StudentComplexDataToDecodeJSONII struct {
	FirstName, lastName string
	HeightInMeters      float64
	IsMale              bool
	Languages           [2]string
	Subjects            []string
	Grades              map[string]string
	Profile             *ProfileComplexData
}

// AccountComplexData declares `AccountComplexData` structure
type AccountComplexData struct {
	IsMale bool
	Email  string
}

// StudentWithProfileAndAccountComplexData declares `StudentWithProfileAndAccountComplexData` structure
type StudentWithProfileAndAccountComplexData struct {
	FirstName, lastName string
	HeightInMeters      float64
	IsMale              bool
	ProfileComplexData
	AccountComplexData
}

// ProfileToDecodeUsingStructureTags declares `ProfileToDecodeUsingStructureTags` structure
type ProfileToDecodeUsingStructureTags struct {
	Username  string `json:"uname"`
	Followers int    `json:"f_count"`
}

// StudentToDecodeUsingStructureTags declares `StudentToDecodeUsingStructureTags` structure
type StudentToDecodeUsingStructureTags struct {
	FirstName      string                            `json:"fname"`
	LastName       string                            `json:"-"` // discard
	HeightInMeters float64                           `json:"height"`
	IsMale         bool                              `json:"male"`
	Languages      []string                          `json:",omitempty"`
	Profile        ProfileToDecodeUsingStructureTags `json:"profile"`
}

type StudentToWorkWithMaps map[string]interface{}

func main() {
	// define `john` struct
	john := StudentStruct{
		FirstName:      "John",
		lastName:       "Doe",
		Age:            21,
		HeightInMeters: 1.75,
		IsMale:         true,
	}

	/********************************************/
	/*             Encode JSON                  */
	/********************************************/

	fmt.Println()
	fmt.Println("Encode JSON")

	/********************************************/
	/*            Encode a struct               */
	/********************************************/

	fmt.Println()
	fmt.Println("Encode a struct")

	// encode `john` as JSON
	johnJSONFromStruct, _ := json.Marshal(john)

	// print JSON string
	fmt.Println(string(johnJSONFromStruct))

	/********************************************/
	/*             Encode a map                 */
	/********************************************/

	fmt.Println()
	fmt.Println("Encode a map")

	// encode `john` as JSON
	johnJSONFromMap, _ := json.Marshal(john)

	// print JSON string
	fmt.Println(string(johnJSONFromMap))

	/********************************************/
	/*          Abstract data type              */
	/********************************************/

	fmt.Println()
	fmt.Println("Abstract data type ")

	var mike StudentAbstractDataType

	// define `mike` struct
	mike = StudentAbstractDataType{
		FirstName: "Mike",
		lastName:  "Doe",
		Age:       21,
		Profile: Profile{
			Username:  "mikedoe91",
			followers: 1975,
			Grades:    map[string]string{"Math": "A", "Science": "A+"},
		},
		Languages: []string{"English", "French"},
	}

	// encode `mike` as JSON
	mikeJSON, err := json.MarshalIndent(mike, "", "  ")

	// print JSON string
	fmt.Println()
	fmt.Println("Nested struct fields")
	fmt.Println(string(mikeJSON), err)

	var peter StudentAbstratDataTypeWithProfileAnonymous

	// define `john` struct
	peter = StudentAbstratDataTypeWithProfileAnonymous{
		FirstName: "John",
		lastName:  "Doe",
		Age:       21,
		Profile: Profile{
			Username:  "johndoe91",
			followers: 1975,
		},
		Languages: []string{"English", "French"},
	}

	// encode `john` as JSON
	peterJSON, _ := json.MarshalIndent(peter, "", "  ")

	// print JSON string
	fmt.Println()
	fmt.Println("Promoted struct nested anonymous files")
	fmt.Println(string(peterJSON))

	// define `john` struct (pointer)
	bryan := &StudentWithProfileInterface{
		FirstName: "Peter",
		lastName:  "Doe",
		Age:       21,
		Primary: &ProfileToImplementFollowMethodInProfileInterface{
			Username:  "peterdoe91",
			Followers: 1975,
		},
	}

	// follow `peter`
	bryan.Primary.Follow()

	// encode `john` as JSON
	bryanJSON, _ := json.MarshalIndent(bryan, "", "  ")

	// print JSON string
	fmt.Println()
	fmt.Println("Nested struct interfaces fields")
	fmt.Println(string(bryanJSON))

	/********************************************/
	/*         Conversion data type             */
	/********************************************/

	fmt.Println()
	fmt.Println("Conversion data type")

	// define `john` struct (pointer)
	steve := &StudentWithFileAge{
		FirstName: "Steve",
		lastName:  "Doe",
		Age:       21,
		Profile: ProfileToImplementFollowMethodInProfileInterface{
			Username:  "stevedoe91",
			Followers: 1975,
		},
	}

	// encode `john` as JSON
	steveJSON, _ := json.MarshalIndent(steve, "", "  ")

	// print JSON string
	fmt.Println(string(steveJSON))

	/********************************************/
	/*         Using Structure Tags             */
	/********************************************/

	fmt.Println()
	fmt.Println("Using Structure Tags")

	// define `john` struct (pointer)
	dick := &StudentWithStructureTags{
		FirstName: "Dick",
		LastName:  "", // empty
		Age:       21,
		Email:     "dick@doe.com",
		Profile: ProfileWithStructureTags{
			Username:  "dickdoe91",
			Followers: 1975,
		},
	}

	// encode `john` as JSON
	dickJSON, _ := json.MarshalIndent(dick, "", "  ")

	// print JSON string
	fmt.Println(string(dickJSON))

	/********************************************/
	/*            Decode JSON                   */
	/********************************************/

	fmt.Println()
	fmt.Println("Decode JSON")

	/********************************************/
	/*            json.isValid                  */
	/********************************************/

	fmt.Println()
	fmt.Println("json.isValid")

	// some JSON data
	JSONData := []byte(`
  {
    "FirstName": "John",
    "Age": 21,
    "Username": "johndoe91",
    "Grades": null,
    "Languages": [
      "English",
      "French"
    ]
  }`)

	// check if `data` is valid JSON
	isValid := json.Valid(JSONData)
	fmt.Println(isValid)

	/********************************************/
	/*            json.Unmarshal                */
	/********************************************/

	fmt.Println()
	fmt.Println("json.Unmarshal ")

	// some JSON data
	billJSON := []byte(`
  {
    "FirstName": "Bill",
    "lastName": "Doe",
    "Age": 21,
    "HeightInMeters": 175,
    "Username": "billdoe91"
  }`)

	// create a data container
	var bill StudentAbstractDataType

	// unmarshal `JSONDataWithBill`
	fmt.Printf("Error: %v\n", json.Unmarshal(billJSON, &bill))

	// print `bill` struct
	fmt.Printf("%#v\n", bill)

	/********************************************/
	/*         Handling complex data            */
	/********************************************/

	fmt.Println()
	fmt.Println("Handling complex data I")

	// some JSON data
	marcJSON := []byte(`
  {
    "FirstName": "Marc",
    "HeightInMeters": 1.75,
    "IsMale": null,
    "Languages": [ "English", "Spanish", "German" ],
    "Subjects": [ "Math", "Science" ],
    "Grades": { "Math": "A" },
    "Profile": {
      "Username": "marcdoe91",
      "Followers": 1975
    }
  }`)

	// create a data container
	marc := StudentComplexDataToDecodeJSON{
		IsMale:   true,
		Subjects: []string{"Art"},
		Grades:   map[string]string{"Science": "A+"},
	}

	// unmarshal `data`
	fmt.Printf("Error: %v\n", json.Unmarshal(marcJSON, &marc))

	// print `john` struct
	fmt.Printf("%#v\n", marc)

	fmt.Println()
	fmt.Println("Handling complex data II")

	// some JSON data
	dereckJSON := []byte(`
  {
    "FirstName": "Dereck",
    "HeightInMeters": 1.75,
    "IsMale": null,
    "Languages": [ "English" ],
    "Subjects": [ "Math", "Science" ],
    "Grades": null,
    "Profile": { "Followers": 1975 }
  }`)

	// create a data container
	var dereck StudentComplexDataToDecodeJSON = StudentComplexDataToDecodeJSON{
		IsMale:    true,
		Languages: [2]string{"Korean", "Chinese"},
		Subjects:  nil,
		Grades:    map[string]string{"Math": "A"},
		Profile:   ProfileComplexData{Username: "dereckdoe91"},
	}

	// unmarshal `data`
	fmt.Printf("Error: %v\n\n", json.Unmarshal(dereckJSON, &dereck))

	// print `john` struct
	fmt.Printf("%#v\n\n", dereck)
	fmt.Printf("%#v\n", dereck.Profile)

	/********************************************/
	/*         Promoted Fields                  */
	/********************************************/

	fmt.Println()
	fmt.Println("Promoted Fields")

	// some JSON data
	sonyJSON := []byte(`
  {
    "FirstName": "Sony",
    "HeightInMeters": 1.75,
    "IsMale": true,
    "Username": "sonydoe91",
    "Followers": 1975,
    "Account": { "IsMale": true, "Email": "sony@doe.com" }
  }`)

	// create a sonyJSON container
	var sony StudentWithProfileAndAccountComplexData

	// unmarshal `sonyJSON`
	fmt.Printf("Error: %v\n", json.Unmarshal(sonyJSON, &sony))

	// print `sony` struct
	fmt.Printf("%#v\n", sony)

	/********************************************/
	/*       Usando Structure Tags              */
	/********************************************/

	fmt.Println()
	fmt.Println("Usando Structure Tags ")

	// some JSON data
	alexJSON := []byte(`
  {
    "fname": "Alex",
    "LastName": "Doe",
    "height": 1.75,
    "IsMale": true,
    "Languages": null,
    "profile": {
      "uname": "alexdoe91",
      "Followers": 1975
    }
  }`)

	// create a data container
	var alex StudentToDecodeUsingStructureTags = StudentToDecodeUsingStructureTags{
		Languages: []string{"English", "French"},
	}

	// unmarshal `alexJSON`
	fmt.Printf("Error: %v\n", json.Unmarshal(alexJSON, &alex))

	// print `alex` struct
	fmt.Printf("%#v\n", alex)

	/********************************************/
	/*          Working with maps              */
	/********************************************/

	fmt.Println()
	fmt.Println("Working with maps I")

	// some JSON data
	phillJSON := []byte(`
  {
    "id": 123,
    "fname": "Phill",
    "height": 1.75,
    "male": true,
    "languages": null,
    "subjects": [ "Math", "Science" ],
    "profile": {
      "uname": "johndoe91",
      "f_count": 1975
    }
  }`)

	// create a data container
	var phill StudentToWorkWithMaps

	// unmarshal `data`
	fmt.Printf("Error: %v\n", json.Unmarshal(phillJSON, &phill))

	// print `john` map
	fmt.Printf("%#v\n\n", phill)

	// iterate through keys and values
	i := 1

	for k, v := range phill {
		fmt.Printf("%d: key (`%T`)`%v`, value (`%T`)`%#v`\n", i, k, k, v, v)
		i++
	}

	fmt.Println()
	fmt.Println("Working with maps II")

	// some JSON data
	meryJSON := []byte(`
  {
    "id": 123,
    "fname": "Mery",
    "height": 1.75,
    "male": true,
    "languages": null,
    "subjects": [ "Math", "Science" ],
    "profile": {
      "uname": "merydoe91",
      "f_count": 1975
    }
  }`)

	// create a data container
	var mery interface{}
	fmt.Printf("Before: `type` of `john` is %T and its `value` is %v\n", john, john)

	// unmarshal `data`
	fmt.Printf("Error: %v\n", json.Unmarshal(meryJSON, &mery))
	fmt.Printf("After: `type` of `john` is %T\n\n", mery)

	// print `john` map
	fmt.Printf("%#v\n", mery)

	/********************************************/
	/*      Codificador y Decodificador         */
	/********************************************/

	fmt.Println()
	fmt.Println("Codificador y Decodificador")

	fmt.Println()
	fmt.Println("Codificador")

	type Person struct {
		Name string
		Age  int
	}

	// create a buffer to hold JSON data
	buf := new(bytes.Buffer)
	// create JSON encoder for `buf`
	bufEncoder := json.NewEncoder(buf)

	// encode JSON from `Person` structs
	bufEncoder.Encode(Person{"Ross Geller", 28})
	bufEncoder.Encode(Person{"Monica Geller", 27})
	bufEncoder.Encode(Person{"Jack Geller", 56})

	// print contents of the `buf`
	fmt.Println(buf) // calls `buf.String()` method

	fmt.Println()
	fmt.Println("Decodificador")

	// create a strings reader
	jsonStream := strings.NewReader(`
    {"Name":"Ross Geller","Age":28}
    {"Name":"Monica Geller","Age":27}
    {"Name":"Jack Geller","Age":56}
  `)

	// create JSON decoder using `jsonStream`
	decoder := json.NewDecoder(jsonStream)

	// create `Person` structs to hold decoded data
	var ross, monica Person

	// decode JSON from `decoder` one line at a time
	decoder.Decode(&ross)
	decoder.Decode(&monica)

	// see value of the `ross` and `monica`
	fmt.Printf("ross: %#v\n", ross)
	fmt.Printf("monica: %#v\n", monica)
}
