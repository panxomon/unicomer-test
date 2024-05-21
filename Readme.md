# Holiday Service

This service provides information about holidays within a specific date range and type.

## Service Execution

### Prerequisites

- Go (version 1.22) installed on your system. You can download it [here](https://golang.org/dl/).

## Installation

1. Clone this repository:

```bash
 git clone https://gitlab.com/united-marmot-association/wizard-generator.git
```

2. Navigate to the project directory:

```bash
 cd unicomer-test
```

3. Install dependencies:

```bash
 make dep
```


## Usage 

1. Manage your own environment variables in the file .env (you can use the file .env.example as a template)

```bash
 cp .env.example .env
```

2. Run the project:

```bash
 make run
```

## Endpoints
### Generate project

1. Call the API endpoint to generate code:

```bash
curl --location 'localhost:8080/holiday/v1?start=2024-01-01&end=2024-11-30&type=Religioso' \
--header 'Content-Type: application/json'
```

2. ready!

##  Utilities

- postman collection is available to work with the endpoints: `./postman`


