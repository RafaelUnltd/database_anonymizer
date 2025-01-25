# Database Anonymizer

Este repositório apresenta uma ferramenta com a finalidade de anonimizar e replicar dados de um banco de dados de origem para um banco de dados de destino. Fique à vontade para explorar o código e contribuir com seu desenvolvimento. Além disso, sua utilização é livre e sem qualquer cobrança.

## Requisitos

- Docker instalado, em sua última versão estável.
- Linguagem Go instalada, em sua última versão estável.
- Postgres instalado localmente, com acesso ao **psql**.

## Como utilizar a aplicação?

Adicione um arquivo **dump.sql** à pasta **scripts/init_db**, ele será restaurado no banco de dados de origem, para que a aplicação consiga funcionar localmente.

Em seguida, inicie a aplicação com o seguinte comando:

```bash
make run
```

Verifique o estado dos conteineres e então faça a restauração do **dump.sql** para o banco de dados de origem, utilizando o comando:

```bash
make seed-input-db
```

Agora, basta fazer uma requisição com os dados a serem anonimizados, seguindo o template:

```json
{
    "input_connection_info": {
        "host": "database-input",
        "user": "root",
        "password": "root",
        "database": "anonymize-input",
        "port": "5432"
    },
    "output_connection_info": {
        "host": "database-output",
        "user": "root",
        "password": "root",
        "database": "anonymize-output",
        "port": "5432"
    },
    "anonymization_rules": {
        "table": {
            "identifier": "id",
            "columns": {
                "col1": {
                    "type": "mask",
                    "value": "Aluno do estratégia #####",
                    "unique": false
                },
                "col2": {
                    "type": "hide",
                    "value": "",
                    "unique": false
                },
                "col3": {
                    "type": "hide",
                    "value": "",
                    "unique": false
                },
                "col4": {
                    "type": "hide",
                    "value": "",
                    "unique": false
                },
                "col5": {
                    "type": "mask",
                    "value": "########",
                    "unique": false
                },
                "col6": {
                    "type": "mask",
                    "value": "###########",
                    "unique": true
                }
            }
        }
    }
}
```