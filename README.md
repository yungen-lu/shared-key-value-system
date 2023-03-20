🔑📝 共用 Key-Value 列表系统

# API

## Create a list

### Request

`POST /v1/head`

```json
{
    "key": "test-head",
    "next_page_key": null,
    "latest_page_key": null
}
```

## Get all lists

### Request

`GET /v1/head`

### Response

```json
[
    {
        "ID": 1,
        "Key": "test-head",
        "NextPageKey": "test-page1",
        "LatestPageKey": "test-page3"
    }
]
```

## Get one list

### Request

`GET /v1/head/:key`

### Response

```json
{
    "ID": 4,
    "Key": "test-head",
    "NextPageKey": "test-page1",
    "LatestPageKey": "test-page3"
}
```

## Delete a list

### Request

`DELETE /v1/head/:key`

## Create a page

### Request

`POST /v1/page`

```json
{
    "key": "test-page1",
    "next_page_key": null,
    "list_key": "test-head"
}
```

## Get all pages

### Request

`GET /v1/page`

### Response

```json
[
    {
        "ID": 13,
        "Key": "test-page1",
        "Articles": null,
        "NextPageKey": "test-page2",
        "ListKey": "test-head"
    },
    {
        "ID": 14,
        "Key": "test-page2",
        "Articles": null,
        "NextPageKey": "test-page3",
        "ListKey": "test-head"
    },
    {
        "ID": 15,
        "Key": "test-page3",
        "Articles": null,
        "NextPageKey": null,
        "ListKey": "test-head"
    }
]
```

## Get one page

### Request

`GET /v1/page/:key`

### Response

```json
{
    "ID": 14,
    "Key": "test-page2",
    "Articles": null,
    "NextPageKey": "test-page3",
    "ListKey": "test-head"
}
```

## Delete a page

### Request

`DELETE /v1/page/:key`

## Delete outdated list and pages

### Request

`DELETE /v1/head`

### Response

```json
{
    "count": 1
}
```



# How it works

There are two major tables, list and page; the list table stores both the first and the last key of the pages, and every page holds a key that points to the next page.

Pages are inserted into the database as usual but automatically update the corresponding list table and pages when inserted or deleted by PostgreSQL triggers.

Every page has a column with a foreign key that references the list table; we use DELETE CASCADE so that when the list row is deleted, the whole list of pages is also deleted.

# Why choose PostgreSQL

Considering that the article noted the high traffic accessing the database and the consistent updates by users, I opted for PostgreSQL due to its great ability to handle concurrent access and maintain consistency while allowing concurrent operations.
