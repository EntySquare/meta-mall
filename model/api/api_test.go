package api

import (
	"meta-mall/database"
	"testing"
)

func TestRunP(t *testing.T) {
	database.ConnectDB()

	//IncomeRunP(database.DB)
	CovenantCycle(database.DB)

}
