# metrics

```json
{
  "id": "5388751619798277776",
  "latency": 574336154,
  "started_at": "2024-05-24T22:19:40.390569286-03:00",
  "finished_at": "2024-05-24T22:19:40.96490561-03:00",
  "watcher": {
    "id": "79e4b3cb-f280-4511-bf95-d4157584baba",
    "name": "/v1/example/outis",
    "started_at": "2024-05-24T22:19:39.389977699-03:00"
  },
  "routine": {
    "routine_id": "f92f61a9-2ddc-4025-af64-212bfda3e151",
    "name": "Teste da minha rotina",
    "path": "/outis/example/main.go:24",
    "started_at": "2024-05-24T22:19:39.390018896-03:00"
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
      "created_at": "2024-05-24T22:19:40.782066655-03:00"
    }
  ],
  "histograms": [
    {
      "key": "problems_resolved",
      "values": [
        {
          "value": 48,
          "created_at": "2024-05-24T22:19:40.438977176-03:00"
        },
        {
          "value": 107,
          "created_at": "2024-05-24T22:19:40.546351975-03:00"
        },
        {
          "value": 101,
          "created_at": "2024-05-24T22:19:40.647606016-03:00"
        },
        {
          "value": 100,
          "created_at": "2024-05-24T22:19:40.747834238-03:00"
        },
        {
          "value": 34,
          "created_at": "2024-05-24T22:19:40.782038289-03:00"
        }
      ]
    },
    {
      "key": "canceled_processes",
      "values": [
        {
          "value": 16,
          "created_at": "2024-05-24T22:19:40.798237357-03:00"
        },
        {
          "value": 61,
          "created_at": "2024-05-24T22:19:40.859567254-03:00"
        },
        {
          "value": 105,
          "created_at": "2024-05-24T22:19:40.964847216-03:00"
        }
      ]
    }
  ],
  "logs": [
    {
      "level": "INFO",
      "message": "script 'Teste da minha rotina' (rid: f92f61a9-2ddc-4025-af64-212bfda3e151) initialized",
      "timestamp": "2024-05-24T22:19:39.390109993-03:00"
    },
    {
      "level": "INFO",
      "message": "script 'Teste da minha rotina' (rid: f92f61a9-2ddc-4025-af64-212bfda3e151, id: 5388751619798277776) initialized",
      "timestamp": "2024-05-24T22:19:40.390630463-03:00"
    },
    {
      "level": "DEBUG",
      "message": "Hello 02",
      "timestamp": "2024-05-24T22:19:40.782066143-03:00"
    },
    {
      "level": "INFO",
      "message": "script 'Teste da minha rotina' (rid: f92f61a9-2ddc-4025-af64-212bfda3e151, id: 5388751619798277776) finished",
      "timestamp": "2024-05-24T22:19:40.96490507-03:00"
    }
  ]
}
```
