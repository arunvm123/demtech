# demtech

This mock server emulates Amazon SES's SendEmail V2 [API](https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_SendEmail.html), allowing developers to test various email sending scenarios without using the actual SES service. Different test scenarios can be triggered by setting the ```Scenario``` header when calling the ```/v2/email/outbound-emails``` endpoint. For simplicity, authentication only requires setting a ```UserName``` header - any requests with this header will be processed and their logs will be stored under the specified username. Following are the currently supported scenarios 

- success
- unverified_email
- account_suspended
- rate_exceeded
- missing_from
- domain_not_verified
- daily_quota_exceeded

The service provides an analytics API(```/logs```) that aggregates API call logs at two levels:
- User-specific analytics when a username is provided as a query parameter
- Scenario-level overview when no specific user is specified

## How to run on local?

Requirements: 
- go 1.23
- psql 16

The service can read configuration from either a yaml file or environment variables. An example yaml file is given along with this repository. By default the service looks for a config.yaml file at the root.

    go run *.go

The path to the config file can be specified by a flag

    go run *.go -config-path /path/to/yaml

To read from environment variables

    go run *.go -config-env true

## Running with docker

The project can be run with Docker. The environment variables for the docker image can be passed in with the `--env-file` which reads in a file of environment variables. Create a config-env file based on the example given in the repo. To connect to a database outside the docker container on a mac, set DB_HOST as `host.docker.internal` and on linux it should be `172.17.0.1`
  
To build the docker image,

    docker build -t demtech .

and to run it

    docker run -it --env-file config-env -p 9090:9090 demtech


## Running with docker-compose

To quickly set up and run the service and database without manual configuration, you can use Docker Compose:

    docker compose --env-file config-env up