package refgen

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestKeys(t *testing.T) {
	Convey("Keys", t, func(){
		Convey("Should return strings", func(){
			keys, err := Keys("github.com/robbert229/refgen/resources", "Person")
			So(err, ShouldBeNil)
			So(keys, ShouldResemble, []Parameter{
				{
					"FirstName",
					"string",
				}, {
					"LastName",
					"string",
				},
			})
		})

		Convey("Should return floats", func(){
			keys, err := Keys("github.com/robbert229/refgen/resources", "Movie")
			So(err, ShouldBeNil)
			So(keys, ShouldResemble, []Parameter{
				{
					"Title",
					"string",
				},
				{
				"Rating",
					"float32",
				},
				{
					"Year",
					"int",
				},
			})
		})
	})
}