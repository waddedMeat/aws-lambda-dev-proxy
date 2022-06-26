# AWS Lambda Proxy

This is a proxy to mimic the API Gateway request and response for local
development.

## Lambda Event
```
{
  body: string
}
```


## Lambda Response
```
{
  statusCode: number,
  body: string
  headers: {[p: string]: string}
}
```

