package common

import (
	"bytes"
	"database_anonymizer/app/src/structs"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// DumpAndRestoreDatabase coordena o processo de copiar o schema de um banco para outro
func DumpAndRestoreDatabase(input structs.DatabaseConnectionInfo, output structs.DatabaseConnectionInfo) error {
	// Cria um diretório temporário para os dumps
	tmpDir := "/tmp/database_dumps"
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		return fmt.Errorf("erro ao criar diretório temporário: %v", err)
	}

	dumpFileName := filepath.Join(tmpDir, "dump.sql")

	fmt.Println("Dumping schema from input database")
	// Extrai o schema do banco de origem
	err := DumpSchema(input, dumpFileName)
	if err != nil {
		return err
	}

	fmt.Println("Restoring schema to output database")
	// Restaura no banco de destino
	err = RestoreSchema(output, dumpFileName)
	if err != nil {
		return err
	}

	fmt.Println("Removing temporary dump file")
	// Remove o arquivo temporário após a restauração
	err = os.Remove(dumpFileName)
	if err != nil {
		return err
	}

	return nil
}

// DumpSchema extrai apenas a estrutura (sem dados) do banco de dados de origem
func DumpSchema(input structs.DatabaseConnectionInfo, dumpFileName string) error {
	inputDumpString := input.DumpString()

	// Cria um arquivo temporário para armazenar o dump
	outputFile, err := os.Create(dumpFileName)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo dump.sql: %v", err)
	}
	defer outputFile.Close()

	// Prepara o comando pg_dump
	cmd := exec.Command(
		"pg_dump",
		inputDumpString,
	)

	// Configura os streams de saída para capturar erros e redirecionar o output para o arquivo
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = outputFile

	fmt.Printf("Executando comando: pg_dump %s\n", inputDumpString)

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("erro ao executar pg_dump: %v, stderr: %s", err, stderr.String())
	}

	return nil
}

// RestoreSchema aplica o schema do arquivo de dump no banco de dados de destino
// usando o comando psql com o parâmetro -f para executar os comandos do arquivo
func RestoreSchema(output structs.DatabaseConnectionInfo, dumpFileName string) error {
	outputDumpString := output.DumpString()

	cmd := exec.Command(
		"psql",
		outputDumpString,
		"-f",
		dumpFileName,
	)

	return cmd.Run()
}
