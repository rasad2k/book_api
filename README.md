# Book API to practice Go Fiber

## Usage

Available endpoints

| Endpoint                      | Description                          |
|-------------------------------|--------------------------------------|
| /api/v1/books     (GET)       | Get all the books                    |
| /api/v1/book/:id  (GET)       | Get the book with the given id           |
| /api/v1/book      (POST)      | Add a new book                       |
| /api/v1/book/:id  (DELETE)    | Delete the book with the given id    |


## Examples

Sample POST request

```
{
    "author": "J.K.Rowling",
    "name":   "Harry Potter and the Prisoner of Azkaban",
    "year":   1999,
    "rating": 10
}
```

Sample response

```
{
    "ID": 1,
    "CreatedAt": "2021-01-08T21:28:04.444061446+04:00",
    "UpdatedAt": "2021-01-08T21:28:04.444061446+04:00",
    "DeletedAt": null,
    "name": "Harry Potter and the Prisoner of Azkaban",
    "year": 1999,
    "author": "J.K.Rowling",
    "rating": 10
}
```
