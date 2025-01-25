package services

import (
	"database_anonymizer/app/src/cache"
	"database_anonymizer/app/src/common"
	"database_anonymizer/app/src/interfaces"
	"database_anonymizer/app/src/repositories"
	"database_anonymizer/app/src/structs"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service struct {
	cache            cache.CacheManager
	inputRepository  interfaces.RepositoriesInterface
	outputRepository interfaces.RepositoriesInterface
}

func NewService(
	input structs.DatabaseConnectionInfo,
	output structs.DatabaseConnectionInfo,
	truncate []string,
	redis cache.CacheManager,
) (interfaces.ServicesInterface, error) {
	err := common.DumpAndRestoreDatabase(input, output)
	if err != nil {
		fmt.Println("Error dumping and restoring database:", err)
		return nil, err
	}

	fmt.Println("Connecting to input database")
	inputDB, err := gorm.Open(postgres.Open(input.ConnectionString()), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to input database:", err)
		return nil, err
	}

	fmt.Println("Connecting to output database")
	outputDB, err := gorm.Open(postgres.Open(output.ConnectionString()), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to output database:", err)
		return nil, err
	}

	fmt.Println("Creating input repository")
	inputRepository := repositories.NewRepository(inputDB)
	fmt.Println("Creating output repository")
	outputRepository := repositories.NewRepository(outputDB)

	fmt.Println("Setting replication role")
	err = outputRepository.SetReplicationRole(structs.REPLICA)
	if err != nil {
		fmt.Println("Error setting replication role:", err)
		return nil, err
	}

	if len(truncate) > 0 {
		fmt.Println("Truncating tables")
		for _, table := range truncate {
			err = outputRepository.TruncateTable(table)
			if err != nil {
				fmt.Println("Error truncating table:", table, err)
				return nil, err
			}
		}
	}

	return Service{
		cache:            redis,
		inputRepository:  inputRepository,
		outputRepository: outputRepository,
	}, nil
}
