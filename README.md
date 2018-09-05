# Stub Party Service
This repository contains a stub implementation of the [Party service](https://github.com/ONSdigital/ras-party). Every request returns an `HTTP 200 OK` status code.

## Service Information API
* `GET /info` will return information about this service, collated from when it was last built.

### Example JSON Response
```json
{
  "name": "partysvc",
  "version": "1.0.0",
  "origin": "git@github.com:ONSdigital/go-party.git",
  "commit": "b7ae66fbc54ca0abafe30950f69e48de72f66699",
  "branch": "master",
  "built": "2018-09-05T13:00:00Z"
}
```

## Environment
The following environment variables may be overridden:

| Environment Variable | Purpose            | Default Value |
| :------------------- | :------------------| :-------------|
| PORT                 | HTTP listener port | :8081         |

## Copyright
Copyright (C) 2018 Crown Copyright (Office for National Statistics)