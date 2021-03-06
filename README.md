# Data integration challenge


Welcome to Data Integration challenge.

Yawoen company has hired you to implement a Data API for Data Integration team.

Data Integration team is focused on combining data from different heterogeneous sources and providing it to an unified view into entities.

## The challenge

It would be really good if you try to make the code using Go language :)
The other technologies you can feel free to choose.

### 1 - Load treated company data in a database

Read data from CSV file and load into the database to create an entity named **companies**.

This entity should contain the following fields: id, company name and zip code. 

- The loaded data should have the following treatment:
    - **Name:** upper case text
    - **zip:** a five digit text

support file: q1_catalog.csv


### 2 - An API to integrate data using a database

Yawoen now wants to get website data from another source and integrate it with the entity you've just created on the database. When the requirements are met, it's **mandatory** that the **data are merged**.

This new source data must meet the following requirements:

- Input file format: CSV
- Data treatment
    - **Name:** upper case text
    - **zip:** a five digit text
    - **website:** lower case text
- Parameters
    - Name: string
    - Zip: string 
    - Website: string

Build an API to **integrate** `website` data field into the entity records you've just created using **HTTP protocol**.

The `id` field is non existent on the data source, so you'll have to use the available fields to aggregate the new attribute **website** and store it. If the record doesn't exist, discard it.

support file: q2_clientData.csv


### Extra - Matching API to retrieve data

Now Yawoen wants to create an API to provide information getting companies information from the entity to a client. 
The parameters would be `name` and `zip` fields. To query on the database an **AND** logic operator must be used between the fields.

You will need to have a matching strategy because the client might only have a part of the company name. 
Example: "Yawoen" string from "Yawoen Business Solutions".

Output example: 
 ```
 {
 	"id": "abc-1de-123fg",
 	"name": "Yawoen Business Solutions",
 	"zip":"10023",
 	"website": "www.yawoen.com"
 }
 ```

## Notes


- Make sure other developers can easily run the application locally.
- Yawoen isn't picky about the programming language, the database and other tools that you might choose. Just take notice of the market before making your decision.
- Automated tests are mandatory.
- Document your API: fill out a **README.md** file with instructions on how to install and use it.


## Deliverable


- :heavy_check_mark: It would be REALLY nice if it was hosted in a git repo of your **own**. You can create a new empty project, create a branch and Pull Request it to the new master branch you have just created. Provide the PR URL for us so we can discuss the code :grin:. BUT if you'd rather, just compress this directory and send it back to us.
- :heavy_check_mark: Make sure Yawoen folks will have access to the source code.
- :heavy_check_mark: Fill the **Makefile** targets with the apropriated commands (**TODO** tags). That is for easy executing the deliverables (tests and execution). If you have other ideas besides a Makefile feel free to use and reference it on your documentation.
- :x: **Do not** start a Pull Request to this project.

Have fun!



## Setup local development

### Install tools

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [Golang](https://golang.org/)
- [Homebrew](https://brew.sh/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

    ```bash
    brew install golang-migrate
    ```

### Setup infrastructure

- Start postgres container:

    ```bash
    make postgres
    ```

- Create companies database:

    ```bash
    make createdb
    ```

- Run database migration:

    ```bash
    make migrateup
    ```

### How to run

- Run 
    ```bash
    make start
    ```
`if API running correctly, you can see in http://localhost:3333/`

### API Rotes

- GET: http://localhost:3333/companies/all 
- GET (Id): http://localhost:3333/companies/:id   
- POST: http://localhost:3333/companies/upload
- GET with parameters:http://localhost:3333/companies?name={NAME}&zipcode={ZIPCODE}  `exemple: http://localhost:3333/api/v1/companies?name=CORRECTIONS&zipcode=94002`

### Notes

- Postgress database
- [Docker desktop](http://gorm.io/)
- [Fiber](https://github.com/gofiber/fiber)