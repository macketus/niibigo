package main

func doMigrate(arg2, arg3 string) error {
	dsn := getDSN()

	// Run the migration command
	switch arg2 {
	case "up":
		err := nii.MigrateUp(dsn)
		if err != nil {
			return err
		}

	case "down":
		if arg3 == "all" {
			err := nii.MigrationDownAll(dsn)
			if err != nil {
				return err
			}
		} else {
			err := nii.Steps(-1, dsn)
			if err != nil {
				return err
			}
		}
	case "reset":
		err := nii.MigrationDownAll(dsn)
		if err != nil {
			return err
		}
		err = nii.MigrateUp(dsn)
		if err != nil {
			return err
		}
	default:
		showHelp()
	}

	return nil
}
