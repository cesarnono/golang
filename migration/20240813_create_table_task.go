package migration

import "gofr.dev/pkg/gofr/migration"

const createTable = `CREATE TABLE IF NOT EXISTS task
(
  id varchar(100)  NOT NULL,
  description text  NOT NULL,
  status varchar(20)  NOT NULL,
  date_created timestamp NOT NULL,
  date_updated timestamp NOT NULL,
  PRIMARY KEY (id)
);`

func createTableTask() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(createTable)
			if err != nil {
				return err
			}
			return nil
		},
	}
}