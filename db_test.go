package ydb
   
import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestName(t *testing.T) {

	Convey("NewDb should default the driver to 'sqlite3'", t, func() {
		db := NewDb("my.db")
		So(db.Driver, ShouldEqual, DefaultSqlDriver)
	})

	Convey("NewDb should not have instantiated a database", t, func() {
		db := NewDb("my.db")
		So(db.Db, ShouldBeNil)
	})

	Convey("NewDb should not have instantiated a database", t, func() {
		db := NewDb("my.db")
		So(db.Db, ShouldBeNil)
	})

	Convey("NewDb should not have any errors", t, func() {
		db := NewDb("my.db")
		So(db.Err, ShouldBeNil)
	})

	Convey("NewDb should hold filename", t, func() {
		db := NewDb("my.db")
		So(db.Filename, ShouldEqual, "my.db")
	})
}


