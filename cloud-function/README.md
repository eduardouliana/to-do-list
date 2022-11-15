# to-do-list

# Exemplos:

### Salvar
```
curl -X POST https://us-east1-{PROJECT_ID}.cloudfunctions.net/task-go?operation=save \
-H "Content-Type: application/json" \
-d '{"id": "1", "description": "teste", "done": false}'
```

### Buscar um registro
```
curl -X POST https://us-east1-{PROJECT_ID}.cloudfunctions.net/task-go?operation=read \
-H "Content-Type: application/json" \
-d '{"id": "1"}'
```

### Buscar todos os registros
```
curl -X POST https://us-east1-{PROJECT_ID}.cloudfunctions.net/task-go?operation=readAll
```

### Excluir um registro
```
curl -X POST https://us-east1-{PROJECT_ID}.cloudfunctions.net/task-go?operation=delete \
-H "Content-Type: application/json" \
-d '{"id": "1"}'
```

### Excluir todos os registros
```
curl -X POST https://us-east1-{PROJECT_ID}.cloudfunctions.net/task-go?operation=deleteAll
```