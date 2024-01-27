package fetcher

import (
	"io"
	"os"
	"squall/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/prashantv/gostub"
)

func TestMain(m *testing.M) {
	gin.DefaultWriter = io.Discard
	stubs := gostub.StubFunc(&getDB, models.NewTestDB())

	exitVal := m.Run()

	stubs.Reset()

	os.Exit(exitVal)
}
