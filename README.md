# Assignment-2

This is the work of Lars Blütecher Holter

## How to run the cloud technologies assignment 2
To run the assignment, you got to have some things installed on your computer. 
What you need is:
* Golang
* Postman
* IDE Environment, we are using Visual Studio Code
* Cloned repo to your computer

After you have made sure you have all those things downloaded and setup you can start to run the program.
Start by opening a terminal in the project folder, then type : "go run .\cmd\firebase.go" and press enter
When all this is done, open up _Postman_ and enter the GET, POST, PUT or DELETE endpoints of your choosing. **NB!** Remember to start the address in Postman with the default url followed by the endpoint you want.

In this README you'll see examples provided of how to interract with the database via postman. We have not implemented arguments via the URL so everything is strictly through postman.

## Endpoint overview

### Default url: http://localhost:8080/

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



## Assignment 2 text https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022/-/wikis/Assignment-2

[TOC]

# Overview

In this assignment, you are going to develop a REST web application in Golang that provides the client with the ability to retrieve information about Corona cases occurring in different countries, as well as the number and stringency of current policies in place. For this purpose, you will interrogate existing web services and return the result in a given output format. 

The REST web services you will be using for this purpose are:

* *Covid 19 Cases API*: https://github.com/rlindskog/covid19-graphql
<!-- uses Country names as input -->

* *Corona Policy Stringency API*: https://covidtracker.bsg.ox.ac.uk/about-api
<!-- uses ISO Codes as input -->

The first API focuses on the provision of information about Corona cases per country as reported by the John Hopkins Institute. The second API provides you with an assessment of policy responses addressing the corona situation.

The API documentation is provided under the corresponding links, and both services vary vastly with respect to feature set and quality of documentation. Use [Postman](https://www.postman.com/) to explore the APIs, but **be mindful of rate-limiting**.

*A general note: When you develop your services that interrogate existing services, **try to find the most efficient way of retrieving the necessary information**. This generally means reducing the number of requests to these services to a minimum by using the most suitable endpoint that those APIs provide. As part of the development, and **for the purpose of testing, we expect you to [stub](https://softwareengineering.stackexchange.com/questions/271720/what-does-stubbing-mean-in-programming) the services**.* e.g. make sure NOT to use the API services in your tests.

The final web service should be deployed on our local OpenStack instance SkyHigh. The initial development should occur on your local machine. For the submission, you will need to provide both a URL to the deployed service as well as your code repository.

In the following, you will find the specification for the REST API exposed to the user for interrogation/testing.

# Specification

Note: Please post an issue if the specification is unclear - so we can clarify and refine it if needed.

Some of the tasks mentioned below are marked as **Advanced Tasks**. These are **optional**, and not a requirement to pass the assignment, but will of course contribute positively to your grade.

## Endpoints

Your web service will have four resource root paths (but potentially further optional ones - see at the bottom): 

```
/corona/v1/cases/
/corona/v1/policy/
/corona/v1/status/
/corona/v1/notifications/
```


## Covid-19 Cases per Country

The initial endpoint focuses on return the latest number of confirmed cases and deaths for a given country, alongside growth rate of cases. 

<!--Optionally, the user can specify a date range. 
 * Where such range is specified (in YYYY-MM-DD format), the endpoint provides the newly reported confirmed and recovered cases within this time frame (i.e., not including previous ones). 
 * Where no further constraints/parameters are specified, the endpoint reports the *total numbers* for both confirmed and recovered cases.-->
### - Request

```
Method: GET
Path: /corona/v1/cases/{:country_name}
```

```{:country_name}``` refers to the name for the country as supported by the *Covid 19 cases API*.

Example request: ```/corona/v1/cases/Norway```

### - Response

* Content type: `application/json`
* Status code: 200 if everything is OK, appropriate error code otherwise. Ensure to deal with errors gracefully.

Body (Example):
```
{
    "country": "Norway",
    "date": "2022-03-05",
    "confirmed": 1305006,
    "recovered": 0,
    "deaths": 1664,
    "growth_rate": 0.004199149089414866
}
```

* **Advanced Task**: Extend ```{:country_name}``` to allow for ISO 3166-1 alpha-3 country code as input.

## Covid Policy Stringency per Country

The second endpoint provides an overview of the *current stringency level* of policies regarding Covid-19 for a given country, in addition to the number of currently active policies. 

Note:
* The stringency information should be drawn from the `stringency_actual` field in the *Corona Policy Stringency API*. Where the `stringency_actual` field is not filled, fall back to the `stringency` field.

### - Request

```
Method: GET
Path: /corona/v1/policy/{:country_name}{?scope=YYYY-MM-DD}
```

```{:country_name}``` refers to the ISO 3166-1 alpha-3 country code.

```{?scope=YYYY-MM-DD}``` indicates the date for which the policy stringency information should be returned. Note that this field is optional (see below for more information). 

Example request: ```/corona/v1/policy/FRA?scope=2022-01-01```

### - Response

* Content type: `application/json`
* Status code: 200 if everything is OK, appropriate error code otherwise. Ensure to deal with errors gracefully.

Body (Example):
```
{
    "country_code": "FRA",
    "scope": "2022-03-01",
    "stringency": 63.89,
    "policies": 0
}
```

Notes:
* If no date range is provided, the values for the current date should be returned. 
* More generally, where information is missing (e.g., stringency information), report the value -1.
* Where no policies are in place, a value of 0 should be returned for ```policies```.

## Status Interface

The status interface indicates the availability of all individual services this service depends on. These can be more services than the ones specified above (if you considered the advanced tasks). If you include more, you can specify additional keys with the suffix `api`. The reporting occurs based on status codes returned by the dependent services. The status interface further provides information about the number of registered webhooks (more details is provided in the next section), and the uptime of the service.

### - Request

```
Method: GET
Path: corona/v1/status/
```

### - Response

* Content type: `application/json`
* Status code: 200 if everything is OK, appropriate error code otherwise. 

Body:
```
{
   "cases_api": "<http status code for *Covid 19 Cases API*>",
   "policy_api": "<http status code for *Corona Policy Stringency API*>",
   ...
   "webhooks": <number of registered webhooks>,
   "version": "v1",
   "uptime": <time in seconds from the last service restart>
}
```

Note: ```<some value>``` indicates placeholders for values to be populated by the service as described for the corresponding values.

## Notification Endpoint

As an additional feature, users can register webhooks that are triggered by the service based on specified events, specifically if information about given countries is invoked, where the minimum frequency can be specified. Users can register multiple webhooks. The registrations should survive a service restart (i.e., be persistent).

### Registration of Webhook

### - Request

```
Method: POST
Path: /corona/v1/notifications/
```

* Content type: `application/json`

The body contains 
 * the URL to be triggered upon event (the service that should be invoked)
 * the country for which the trigger applies
 * the minimum number of repeated invocations before notification should occur (i.e., "greater equals")

Body (Example):
```
{
   "url": "https://localhost:8080/client/",
   "country": "France",
   "calls": 5
}
```

<!-- <number of minimum invocations after notification should occur> -->

**Advanced Task**: Consider unified handling of country name, independent of whether it is specified as common name or 3-letter ISO code.

### - Response

The response contains the ID for the registration that can be used to see detail information or to delete the webhook registration. The format of the ID is not prescribed, as long it is unique. Consider best practices for determining IDs.

* Content type: `application/json`
* Status code: Choose an appropriate status code

Body (Example):
```
{
    "webhook_id": "OIdksUDwveiwe"
}
```

### Deletion of Webhook

### - Request

```
Method: DELETE
Path: /corona/v1/notifications/{id}
```

* {id} is the ID returned during the webhook registration

### - Response

Implement the response according to best practices.

### View registered webhook

### - Request

```
Method: GET
Path: /corona/v1/notifications/{id}
```

* {id} is the ID for the webhook registration

### - Response

The response is similar to the POST request body, but further includes the ID assigned by the server upon adding the webhook.

* Content type: `application/json`

Body (Example):
```
{
   "webhook_id": "OIdksUDwveiwe",
   "url": "https://localhost:8080/client/",
   "country": "France",
   "calls": 5
}

```

### View all registered webhooks

### - Request

```
Method: GET
Path: /corona/v1/notifications/
```

### - Response

The response is a collection of all registered webhooks.

* Content type: `application/json`

Body (Example):
```
[{
    "webhook_id": "OIdksUDwveiwe",
    "url": "https://localhost:8080/client/",
    "country": "France",
    "calls": 5
 },
 {
    "webhook_id": "DiSoisivucios",
    "url": "https://localhost:8080/client/",
    "country": "Norway",
    "calls": 2
 },
...
]
```

* Note: If you did not unify the country based on common name and ISO code, you can return those separately.

### Webhook Invocation (upon trigger)

When a webhook is triggered, it should send information as follows. Where multiple webhooks are triggered, the information should be sent separately (i.e., not be combined).

```
Method: POST
Path: <url specified in the corresponding webhook registration>
```

* Content type: `application/json`

Body (Example):
```
{
   "webhook_id": "OIdksUDwveiwe",
   "country": "Norway",
   "calls": 3
}
```

# Additional requirements

* All endpoints should be *tested using automated testing facilities provided by Golang*. This includes the stubbing of the third-party endpoints to ensure test reliability. Try to maximize test coverage as reported by Golang.
* Repeated invocations for a given country and date should be cached to minimise invocation on the third-party libraries.
  * **Advanced Task**: Implement purging of cached information for requests older than a given number of days.

# Deployment

The service is to be deployed on an IaaS solution OpenStack using Docker (to be discussed in class). You will need to provide the URL to the deployed service as part of the submission, in addition the source repository.

# Notes

* Feel free to introduce additional endpoints to support the development and debugging.
* Where specification details are missing (but you can infer those), operate based on best practices and document it accordingly.
* Where information is unclear, get in touch with teaching staff for clarification. Where needed, the assignment information will be updated accordingly (and the updated information will be highlighted as **UPDATE**).

# General Aspects

## Professionalism

As indicated during the initial sessions, ensure you work with professionalism in mind (see Course Rules). In addition to professionalism, you are at liberty to introduce further features into your service, as long it does not break the specification given above. 

## Workspace environment

Please work in the provided workspace environment (see [here](https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022/-/wikis/Workspace-Conventions) - lodge an issue if you have trouble accessing it) for your user and create a project `assignment-2` in this workspace.

## Rate limits reminder

As mentioned above, be sensitive to rate limits of external services. This has proven very important, given the large number of projects (and hence invocations) on the third-party services.

## Resources

The course repository provides a range of example projects for various features discussed throughout the lecture sessions. Feel free to borrow from those projects, or use them to understand a concept you are struggling with (e.g., learning the use of Firestore). 

## Third-party libraries

Be deliberative about using third-party libraries (Don't just do it because someone did it on StackOverflow). While those libraries often allow for convenience and functionality you would otherwise need to reimplement, they can also mean the "import" of technological debt, especially if you were to think about maintainability. So, be very clear *why* you want to use the library (lack of functionality in standard packages, convenience, etc.). Note that it may be challenging for us to provide the necessary support, especially if the library is rather specialised (we will rely on the same resources available to you). As the assignment is designed, you will only need Golang standard API functionality, alongside Firebase/store functionality as a third-party dependency.

# Submission

The assignment is an individual assignment. The submission deadline is provided on the course main wiki page. No extensions will be given for late submissions (unless the deadline is collectively extended, i.e., if we agree in class). 

As part of the submission you will need to provide:
* a link to your code repository (ensure it is `internal` at that stage)
* a link to the deployed service

In addition, we will provide you with an option to clarify aspects of your submission (e.g., aspects that don't quite work, or additional features).

The submission occurs via our [submission system](Submission System). Early submission is explicitly encouraged - you can change it (or even withdraw) any time before the deadline.

**Deadline: 8th April 2022, noon, 12:00 CET (Note: if in conflict, the deadline on the main page applies)**

# Peer Review

After the submission deadline, there will be a separate deadline during which you will review other students' submissions. To do this the system provides you with a checklist of aspects to assess. You will need to review *at least two submissions* to meet the mandatory requirements of peer review, but you can review as many submissions as you like, which counts towards your participation mark for the course. The peer-review deadline is indicated on the main course wiki page.
