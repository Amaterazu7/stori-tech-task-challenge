# Stori Tech Task Challenge

## Description
The goal of this project is to develop a system that processor a .csv file with a list of transaction and send summary information by an email, this challenge is for Stori.

The project is implemented in Golang, using Serverless Framework and Make as the build tool. The architecture follows the principles of either Hexagonal Architecture or
Domain-Driven Design (DDD), which organizes the code into layers: domain, application, and infrastructure.

1. Domain Layer:
    - This layer contains all entities and interfaces related to the domain of the application.
    - It encapsulates the core business logic that remains stable and unchanged over time.
    - It use the Repository pattern to improve decoupling to the service.

2. Application Layer:
    - The application layer hosts services that orchestrate the use cases of the application.
    - It defines the use cases related to the csv processor feature, such as handling the creation of transaction, and the gomail email sender implementation.

3. Infrastructure Layer:
    - In the infrastructure layer, implementations associated with the specific requirements of the project are housed.
    - This includes components such as repositories for storing data in MySql or retrieve info from S3, also handle the MySql database connection and any other infrastructure-related logic.


Overall, this project aims to provide a robust and user-friendly solution for managing bulk transactions and send emails with the results. 
By adhering to architectural best practices and leveraging the Golang language, we ensure the project's maintainability, scalability, and extensibility.


### Prerequisites

Before run this project, make sure you have the following tools installed:

- `Go v1.22`
- `Make 3.81`
- `Docker v26.0.0`
- `docker-compose v2.26.1-desktop.1`
- `Serverless Framework v4.1.4`


### Installation and run the app

To run the application, I would recommend have installed all the tools, and now we can continue w/ the project customization:

1. Rename the `.env.example` file to `.env` and change the credentials with your access key.
   - For "SMTP_PASSWORD" I leave the url ti get the access with your gmail in the "SMTP_URL" env_var.
2. Also, in order to receive the email, change the MySql entry point file "03":
   - In the root go to `volumes` > `03-insert-values.sql` and change the email value in the insert for an email own by you. 
   - As a tip, is important, that the accountId matches with the "Id" in the table "account" from the database, for safety, I assumed that if there is no Id in the table account, that is mean that the process not belong to the system
3. After all is sett, in the root directory path, go to `transaction-processor` folder and run the `Makefile` command in terminal.
   - ` make start-app-local`
4. Now, you should be able to:
   - Get - ` http://localhost:4007/processor/:accountId `
   - accountId >> "17d340fa-5bf5-4429-8167-bafe4c0af0a7"


### Final Email Screenshot
![Alt text](/source/screenshot.jpeg?raw=true "Stori Email")


## Contributor
[![Linkedin](https://i.stack.imgur.com/gVE0j.png) LinkedIn - Diego Leonel Cañete Crescini](https://www.linkedin.com/in/diegocanete/)
