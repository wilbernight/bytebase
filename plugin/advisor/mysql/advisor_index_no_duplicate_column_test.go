package mysql

// Framework code is generated by the generator.

import (
	"testing"

	"github.com/bytebase/bytebase/plugin/advisor"
)

func TestIndexNoDuplicateColumn(t *testing.T) {
	tests := []advisor.TestCase{
		{
			Statement: `CREATE TABLE t (a int, PRIMARY KEY (a))`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Success,
					Code:    advisor.Ok,
					Title:   "OK",
					Content: "",
				},
			},
		},
		{
			Statement: `CREATE TABLE t (
				a int,
				INDEX idx_a (a, a))`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.DuplicateColumnInIndex,
					Title:   "index.no-duplicate-column",
					Content: "INDEX `idx_a` has duplicate column `t`.`a`",
					Line:    3,
				},
			},
		},
		{
			Statement: `CREATE INDEX idx_a on t(a, a)`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.DuplicateColumnInIndex,
					Title:   "index.no-duplicate-column",
					Content: "INDEX `idx_a` has duplicate column `t`.`a`",
					Line:    1,
				},
			},
		},
		{
			Statement: `ALTER TABLE t ADD INDEX idx_a (a, a)`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.DuplicateColumnInIndex,
					Title:   "index.no-duplicate-column",
					Content: "INDEX `idx_a` has duplicate column `t`.`a`",
					Line:    1,
				},
			},
		},
		{
			Statement: `ALTER TABLE t ADD PRIMARY KEY pk_a (a, a)`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.DuplicateColumnInIndex,
					Title:   "index.no-duplicate-column",
					Content: "PRIMARY KEY `pk_a` has duplicate column `t`.`a`",
					Line:    1,
				},
			},
		},
		{
			Statement: `ALTER TABLE t ADD UNIQUE KEY uk_a (a, a)`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.DuplicateColumnInIndex,
					Title:   "index.no-duplicate-column",
					Content: "UNIQUE KEY `uk_a` has duplicate column `t`.`a`",
					Line:    1,
				},
			},
		},
		{
			Statement: `ALTER TABLE t ADD FOREIGN KEY fk_a (a, a) REFERENCES t1(a, b)`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.DuplicateColumnInIndex,
					Title:   "index.no-duplicate-column",
					Content: "FOREIGN KEY `fk_a` has duplicate column `t`.`a`",
					Line:    1,
				},
			},
		},
	}

	advisor.RunSQLReviewRuleTests(t, tests, &IndexNoDuplicateColumnAdvisor{}, &advisor.SQLReviewRule{
		Type:    advisor.SchemaRuleIndexNoDuplicateColumn,
		Level:   advisor.SchemaRuleLevelWarning,
		Payload: "",
	}, advisor.MockMySQLDatabase)
}