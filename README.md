# Service for searching manga 

## Used
- gRPC
- Mongo
- Docker
- Go 1.18.3

## Todo
- Write tests
- SSL/TLS

## Example 
``` 
    REQUEST 
    {
        "mangaName": "волейбол"
    }

    RESPONSE 
      {
      "name": "Волейбол!!",
      "link": "https://mangapoisk.ru/manga/voleybol-abs3nYF",
      "genre": "комедия, драма",
      "chapters": "Глава: 401",
      "description": "Хината Шоё – паренек маленького роста, но с большой мечтой! Как-то раз он посмотрел национальный чемпионат по волейболу, который проходил ср...",
      "year": "Год: 2012",
      "rate": "4.12"
    }
    etc...
```

