# Velo Payor Example - Go

This is a simplified example of an API for a payor that integrates with Velo Payments. This example uses the [gin](https://github.com/gin-gonic/gin) framework. It also uses the Velo golang client: [velo-go](https://github.com/velopaymentsapi/velo-go)

### Usage

You will need add your api key, api secret, & payor id that your recieved from Velo to environment variables. You can see a sample of these in the `.env.example` file.

```
VELO_API_APIKEY=contact_velo_for_info
VELO_API_APISECRET=contact_velo_for_info
VELO_API_PAYORID=contact_velo_for_info
...
```

We have provided a `make` command to apply your info and copy it into the correct location (`.env`) within the codebase. Replace your info in the following to setup your `.env`

``` shell
make key=YOUR_VELO_APIKEY_HERE secret=YOUR_VELO_APISECRET_HERE payor=YOUR_VELO_PAYORID_HERE env
```

Once your info is in place you can spin up the api & database. Under the hood `docker-compose` will create a network and 3 services.

``` shell
make up
```

You should now be able to access the api at [localhost:4567](http://localhost:4567)

#### Testing

To test the example api import the following Postman collection & environment files via

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/86c785f2ee6edbfc1751#?env%5BVelo%20Payor%20Example%20Dev%5D=W3sia2V5IjoiYXBpX3VybCIsInZhbHVlIjoiaHR0cDovL2xvY2FsaG9zdDo0NTY3IiwiZGVzY3JpcHRpb24iOiIiLCJlbmFibGVkIjp0cnVlfSx7ImtleSI6Imp3dF90b2tlbiIsInZhbHVlIjoiIiwiZGVzY3JpcHRpb24iOiIiLCJlbmFibGVkIjp0cnVlfSx7ImtleSI6InBheWVlX2lkIiwidmFsdWUiOiIiLCJkZXNjcmlwdGlvbiI6IiIsImVuYWJsZWQiOnRydWV9LHsia2V5IjoicGF5bWVudF9pZCIsInZhbHVlIjoiIiwiZGVzY3JpcHRpb24iOiIiLCJlbmFibGVkIjp0cnVlfSx7ImtleSI6InNvdXJjZV9hY2NvdW50X2lkIiwidmFsdWUiOiIiLCJkZXNjcmlwdGlvbiI6IiIsImVuYWJsZWQiOnRydWV9LHsia2V5Ijoic291cmNlX2FjY291bnRfbmFtZSIsInZhbHVlIjoiIiwiZGVzY3JpcHRpb24iOiIiLCJ0eXBlIjoidGV4dCIsImVuYWJsZWQiOnRydWV9XQ==)

Make sure to edit the postman environment file variables with any need values.

All calls are dependant on calling Public > Authenticate ... in order to get a JWT. This needs to occur once for an auth token to the flask api.

#### Cleanup

To clean up the containers first stop them ...

``` shell
make down
```

next delete images with ...

``` shell
make destroy
```

### Work In Progress

This codebase is meant to serve as an example of how to integrate with parts the Velo Platform API. The code is not production ready and is a work in progress!