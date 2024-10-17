package postgres_test

import (
	"testing"

	"github.com/tarunngusain08/culturehub/config"
	"github.com/tarunngusain08/culturehub/pkg/db/postgres"
)

func TestPostgres_ResetDB(t *testing.T) {
	e := config.Startup()
	if e != nil {
		t.Error(e)
	}
	if config.GetBool("test.skipIntegration") {
		t.Skip("skipping test")
	} else {
		// test reset db
		// todo: add statemets
		statements := []string{
			"Create Type round_type AS ENUM ( 'PREFLOP', 'FLOP', 'TURN', 'RIVER', 'SHOWDOWN')",
			"Create Type action AS ENUM ( 'FOLD', 'CHECK', 'BET', 'CALL', 'RAISE')",
		}
		err := postgres.ResetDB(statements)
		if err != nil {
			t.Error(err)
		}
	}
}
