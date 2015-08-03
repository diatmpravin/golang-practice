package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	AF1 string
}

type B struct {
	BF1 string
}

type X struct {
	Things []A
	Thangs []B

	// Or perhaps json.RawMessage if you just
	// want to pass them through.
	// Or map of string/int/etc if the value type is fixed.
	Extra map[string]interface{}
}

// Marshal Way 1: call unmarshal twice on whole input

type xsub struct {
	Things []A `json:"things"`
	Thangs []B `json:"thangs"`
}

func (x *X) _UnmarshalJSON(b []byte) error {
	// First unmarshall the known keys part:
	var tmp xsub
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	// Then unmarshall the whole thing again:
	var vals map[string]interface{}
	if err := json.Unmarshal(b, &vals); err != nil {
		return err
	}

	// Everything worked, chuck the map entries for
	// "known" fields and store results.
	delete(vals, "things")
	delete(vals, "thangs")
	x.Things = tmp.Things
	x.Thangs = tmp.Thangs
	x.Extra = vals
	return nil
}

// Way 2:

// func (x *X) UnmarshalJSON(b []byte) error {
//     // Only partially decode:
//     var tmp map[string]json.RawMessage
//     if err := json.Unmarshal(b, &tmp); err != nil {
//         return err
//     }

//     // Now handle the known fields:
//     var things []A
//     if err := json.Unmarshal(tmp["things"], &things); err != nil {
//         return err
//     }
//     var thangs []B
//     if err := json.Unmarshal(tmp["thangs"], &thangs); err != nil {
//         return err
//     }

//     // And the unknown fields.
//     var extra map[string]interface{}

//     // Either:
//     if true {
//         // this has more calls to Unmarshal, but may be more desirable
//         // as it completely skips over the already handled things/thangs.
//         delete(tmp, "things")
//         delete(tmp, "thangs")
//         // If you only needed to store the json.RawMessage for use
//         // in MarshalJSON then you'd just store "tmp" and stop here.

//         extra = make(map[string]interface{}, len(tmp))
//         for k, raw := range tmp {
//             var v interface{}
//             if err := json.Unmarshal(raw, &v); err != nil {
//                 return err
//             }
//             extra[k] = v
//         }
//     } else { // Or:
//         // just one more call to Unmarshal, but it will waste
//         // time with things/thangs again.
//         if err := json.Unmarshal(b, &extra); err != nil {
//             return err
//         }
//         delete(extra, "things")
//         delete(extra, "thangs")
//     }

//     // no error, we can store the results
//     x.Things = things
//     x.Thangs = thangs
//     x.Extra = extra
//     return nil
// }

// func (x X) MarshalJSON() ([]byte, error) {
//     // abusing/reusing x.Extra, could copy map instead
//     x.Extra["things"] = x.Things
//     x.Extra["thangs"] = x.Thangs
//     result, err := json.Marshal(x.Extra)
//     delete(x.Extra, "things")
//     delete(x.Extra, "thangs")
//     return result, err
// }

func main() {
	inputs := []string{
		`{"things": [], "thangs": []}`,

		`
{
    "things": [
    {
        "AF1": "foo"
    },
    {
        "AF1": "bar"
    }
    ],
    "thangs": [
        {
            "BF1": "string value"
        }
    ],
    "xRandomKey":       "not known ahead of time",
    "xAreValueTypesKnown": 172
}`,
	}

	for _, in := range inputs {
		fmt.Printf("\nUnmarshal(%q):\n", in)
		var x X
		err := json.Unmarshal([]byte(in), &x)
		if err != nil {
			fmt.Println("unmarshal:", err)
		} else {
			fmt.Printf("\tas X: %+v\n", x)
			fmt.Printf("\twith map: %v\n", x.Extra)
			out, err := json.Marshal(x)
			if err != nil {
				fmt.Println("marshal:", err)
				continue
			}
			fmt.Printf("\tRemarshals to: %s\n", out)
		}
	}
}
