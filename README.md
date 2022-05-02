# Assignment-3

This is the work of Lars Blütecher Holter

## How to run the cloud technologies assignment 2
To run the assignment, you got to have some things installed on your computer. 
What you need is:
* Golang
* Postman
* IDE Environment, we are using Visual Studio Code
* Cloned repo to your computer

After you have made sure you have all those things downloaded and setup you can start to run the program.
Start by opening a terminal in the project folder, then type : "go run .\cmd\server.go" and press enter
When all this is done, open up _Postman_ and enter the GET, POST, PUT or DELETE endpoints of your choosing. **NB!** Remember to start the address in Postman with the default url followed by the endpoint you want.

In this README you'll see examples provided of how to interract with the database via postman. The URL/ IP address for the hosted server is : 10.212.139.118
Example for url:
**http://10.212.139.118:8080/corona/v1/notifications/**

## Endpoint overview

### Default url: http://localhost:8080/
### Default server url: http://10.212.139.118:8080/

### Endpoints

    /corona/v1/cases/
    /corona/v1/policy/
    /corona/v1/status/
    /corona/v1/notifications/


#### Cases Endpoint (/corona/v1/cases/):
    Method: GET
    Path: /corona/v1/cases/{:country_name}



##### Example Endpoint commands in Postman:
Method: **Get**

URL: http://localhost:8080/corona/v1/cases/Norway

Body:

    {
    "name": "Norway",
    "date": "2022-04-06",
    "confirmed": 1412969,
    "recovered": 2667,
    "deaths": 0,
    "growthRate": 0.001005277886011831
    }


#### Policy Endpoint (/corona/v1/policy/):
    Method: GET
    Path: /corona/v1/policy/{:country_name}{?scope=YYYY-MM-DD}

##### Example Endpoint commands in Postman:
Method: GET

URL: http://localhost:8080/corona/v1/policy/Nor/

Body:

    {
    "date_value": "",
    "country_code": "",
    "confirmed": -1,
    "deaths": -1,
    "stringency_actual": -1,
    "stringency": -1
    }

URL: http://localhost:8080/corona/v1/policy/Nor/2021-09-17

    {
    "date_value": "2021-09-17",
    "country_code": "NOR",
    "confirmed": 181195,
    "deaths": 841,
    "stringency_actual": 38.89,
    "stringency": 38.89
    }

#### Status Endpoint (/corona/v1/status/):
    Method: GET
    Path: /corona/v1/status/

##### Example Endpoint commands in Postman:
Method: GET

URL: http://localhost:8080/corona/v1/status/


Body:

    {
    "cases_api": 200,
    "policy_api": 200,
    "uptime": 18.6753544,
    "version": "v1",
    "webhooks": 3
    }

#### Notification Endpoint (/corona/v1/notifications/):
    Method: GET, POST, DELETE
    Path: /corona/v1/notifications/

##### Example Endpoint commands in Postman:
Method: GET

URL: http://localhost:8080/corona/v1/notifications/


Body:

    [
        {"url":"https://localhost:8080/client/","country":"France","calls":5,"webhook_id":"32ce679697f07268b0be398ef2ed336b6d17ea24cfd855f5da603aea80c175dd29bf169f8180f3661a758efb291ade759723f8f83d4bd49171ab4f518df8de95"},{"url":"https://localhost:8080/client/","country":"France","calls":5,"webhook_id":"0b4133b964c7d4e8c8cd30ae85e36ddb152cae88171b4aa3b17d9d53675ce4cf0332316f3fde6909b60a475d8bbb25725a92126dbc722287f3c2c6bfbe097be3"},{"url":"https://localhost:8080/client/","country":"France","calls":5,"webhook_id":"57ccb7b5912f02984533e5f1479af090c640fe9b67de485d89eb6b86e6bdb03c16e949f04bbdd582aa6ec00ead5e4193f304edac0f644ddf41c7d5a5023160b0"}
    ]


Method: POST

URL: http://localhost:8080/corona/v1/notifications/

Body:

    {
    "url": "https://localhost:8080/client/",
    "country": "France",
    "calls": 5
    }


#### Status Endpoint (/corona/v1/status/):
    Method: GET
    Path: /corona/v1/status/


##### Example Endpoint commands in Postman:
Method: **Get**

URL: http://10.212.139.118:8080/corona/v1/status/

Body:

    {
    "cases_api": 200,
    "policy_api": 200,
    "uptime": 318.378893772,
    "version": "v1",
    "webhooks": 2
    }


## Assignment 2 text https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022/-/wikis/Assignment-2
