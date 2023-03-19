# Greenlight API

## How to Use 
This project is a Golang API that provides movie and user resources. Here is how to use it:

### Prerequisites
To use this API, you will need:


- Go version 1.13 or higher
- A PostgreSQL database
- A .envrc file with the following environment variables:
    - GREENLIGHT_DB_DSN: the connection string for your PostgreSQL database


## Development
 
### Running the API
 
To run the API, use the following command:

```
make run/api
```

This will start the API server and connect to your PostgreSQL database using the GREENLIGHT_DB_DSN environment variable.

### Database Migrations

To apply database migrations, use the following command:

```
make db/migrations/up
```

This will apply all up database migrations to your PostgreSQL database.

You can also create a new database migration using the following command:

```
make db/migrations/new name=<migration name>
```

### Connecting to the Database
To connect to your PostgreSQL database using psql, use the following command:
```
make db/psql
```

### Building
To build the API, use the following command:

```
make build/api
```
This will build the cmd/api application and output the binary file to ./bin/api.

### Quality Control
To audit the code, run the following command:

```
make audit
```
This will tidy and vendor dependencies, format, vet, and test all the code.

Conclusion
This API is a powerful tool for managing movie and user resources. By following these instructions, you will be able to easily develop, build, and audit the code to ensure high-quality performance.

## Endpoints 

| Methods |  URL Pattern |   Permission |
|:-----|:--------:| ------:
| GET   | /v1/healthcheck| - |
| GET   |  /v1/movies  |   movies:read |
| POST  | /v1/movies |    movies:write |
| GET | /v1/movies/:id |    movies:read |
| PATCH  | /v1/movies/:id |    movies:write |
| DELETE  | /v1/movies/:id |    movies:write |
| POST  | /v1/users |    - |
| PUT  | /v1/users/activated |    - |
| POST  | /v1/tokens/authentication |    - |
| GET | /debug/vars | -

## More features ideas 

 One feature you could add to make the API more complex is a recommendation engine. This could analyze user behavior and preferences, as well as movie metadata, to generate personalized recommendations for each user. The engine could take into account factors such as genre, director, actors, release date, and user ratings to provide the most relevant recommendations.

 Another feature you could add is social interactions, such as the ability for users to rate movies, write reviews, and share their favorite films with friends. This could also include a feature for users to create and join groups based on shared interests, allowing them to discover new movies and engage in discussions with like-minded individuals.

 You could also add a feature for user-generated content, such as the ability for users to create their own lists of favorite movies, or to upload their own movie reviews or film-related content. This could provide a new level of engagement for users and foster a sense of community around the API.

Finally, you could consider adding a feature for content moderation, to ensure that the API remains a safe and respectful space for all users. This could include measures such as user flagging, automatic content filtering, and human moderation to review and remove inappropriate content.