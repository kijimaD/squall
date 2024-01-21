package factories

import (
	"fmt"
	"squall/helper"
	"squall/models"

	"github.com/bluele/factory-go/factory"
	"github.com/rs/xid"
)

var EntryFactory = factory.NewFactory(
	&models.Entry{},
).SeqInt("ID", func(n int) (interface{}, error) {
	return helper.GetPtr(uint(randomID())), nil
}).Attr("URL", func(args factory.Args) (interface{}, error) {
	return fmt.Sprintf("URL-%s", xid.New().String()), nil
}).Attr("Title", func(args factory.Args) (interface{}, error) {
	return fmt.Sprintf("Title-%s", xid.New().String()), nil
}).Attr("IsDone", func(args factory.Args) (interface{}, error) {
	return false, nil
})

func MakeEntry(fields Fields, deps []Dependency) (*models.Entry, []Dependency) {
	m, _ := EntryFactory.MustCreateWithOption(fields).(*models.Entry)
	deps = append(deps, m)

	return m, deps
}
