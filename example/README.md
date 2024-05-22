# metrics

```json
{
  "id": "3463095254098643911",
  "latency": 48514700,
  "started_at": "2024-05-22T19:34:01.072143938-03:00",
  "finished_at": "2024-05-22T19:34:01.120658802-03:00",
  "watcher": {
    "id": "8b1d6a18-5f3d-4482-a574-35d3965c8783",
    "name": "v1/example",
    "started_at": "2024-05-22T19:33:58.178673456-03:00"
  },
  "routine": {
    "id": "0b2d07ca-e3db-478a-9455-d5f476ac8d37",
    "name": "Example routine",
    "path": "/outis/example/main.go:23",
    "started_at": "2024-05-22T19:33:58.17873695-03:00"
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
  "indicators": [
    {
      "key": "customers_notified",
      "value": 13.21,
      "created_at": "2024-05-22T19:34:01.108226823-03:00"
    }
  ],
  "histograms": [
    {
      "key": "problems_resolved",
      "values": [
        {
          "value": 10,
          "created_at": "2024-05-22T19:34:01.082507408-03:00"
        },
        {
          "value": 3,
          "created_at": "2024-05-22T19:34:01.085635542-03:00"
        },
        {
          "value": 10,
          "created_at": "2024-05-22T19:34:01.095840753-03:00"
        },
        {
          "value": 2,
          "created_at": "2024-05-22T19:34:01.098013473-03:00"
        },
        {
          "value": 10,
          "created_at": "2024-05-22T19:34:01.108188919-03:00"
        }
      ]
    },
    {
      "key": "canceled_processes",
      "values": [
        {
          "value": 8,
          "created_at": "2024-05-22T19:34:01.116398926-03:00"
        },
        {
          "value": 4,
          "created_at": "2024-05-22T19:34:01.120625851-03:00"
        },
        {
          "value": 0,
          "created_at": "2024-05-22T19:34:01.120638405-03:00"
        }
      ]
    }
  ],
  "logs": [
    {
      "level": "INFO",
      "message": "script 'Example routine' (rid: 0b2d07ca-e3db-478a-9455-d5f476ac8d37, id: 3463095254098643911) initialized",
      "timestamp": "2024-05-22T19:34:01.072212527-03:00"
    },
    {
      "level": "DEBUG",
      "message": "Hello 02",
      "timestamp": "2024-05-22T19:34:01.108226308-03:00"
    },
    {
      "level": "INFO",
      "message": "script 'Example routine' (rid: 0b2d07ca-e3db-478a-9455-d5f476ac8d37, id: 3463095254098643911) finished",
      "timestamp": "2024-05-22T19:34:01.120658155-03:00"
    }
  ]
}
```
