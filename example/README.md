# metrics

```json
{
  "id": "3031305742566248080",
  "latency": 24899,
  "started_at": "2024-05-21T22:56:08.522689596-03:00",
  "finished_at": "2024-05-21T22:56:08.522714647-03:00",
  "watcher": {
    "id": "8b1d6a18-5f3d-4482-a574-35d3965c8783",
    "name": "v1/example",
    "started_at": "2024-05-21T22:55:24.50165334-03:00"
  },
  "routine": {
    "id": "0b2d07ca-e3db-478a-9455-d5f476ac8d37",
    "name": "Example routine",
    "path": "outis/example/main.go:22",
    "started_at": "2024-05-21T22:55:24.501752417-03:00"
  },
  "metadata": {
    "client": {
      "address": {
        "id": 2345564,
        "number": "S/N",
        "street": "Av. 01"
      },
      "email": "antonio.francisco.silva@email.com",
      "id": 2134234,
      "name": "Antonio Francisco da Silva"
    },
    "uuid": {
      "failure": [
        23423,
        4546,
        3423
      ],
      "success": [
        23423,
        1423,
        4322
      ]
    }
  },
  "logs": [
    {
      "level": "INFO",
      "message": "script 'Example routine' (rid: 0b2d07ca-e3db-478a-9455-d5f476ac8d37, id: 3031305742566248080) initialized",
      "timestamp": "2024-05-21T22:56:08.52270704-03:00"
    },
    {
      "level": "DEBUG",
      "message": "Hello 02",
      "timestamp": "2024-05-21T22:56:08.522711487-03:00"
    },
    {
      "level": "INFO",
      "message": "script 'Example routine' (rid: 0b2d07ca-e3db-478a-9455-d5f476ac8d37, id: 3031305742566248080) finished",
      "timestamp": "2024-05-21T22:56:08.522714171-03:00"
    }
  ]
}
```
