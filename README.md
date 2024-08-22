# tennis buddy

golang back-end; vue.js front-end
- want a challenge?

## TODO

hacks s03e06 kiss music
- [ ] start with ladders ... build from there
- [ ] create user
- [ ] create user dashboard
- [ ] create users admin dashboard
- [ ] which info for each user?
- [ ] api for admin?
- [ ] api for users
- [ ] set up tables (postgres)
- [ ] matching algorithm using ML/AI
- [ ] scrape data from ladders, present info with better interface, and allow users to do something on the site

### changelog

- postgres db name is tennisbuddy
- pro shop needs an inventory management system. 21aug2024
    - dbname is inventorydb
- 22aug2024: changes to main.go
    Configuration and Error Handling:
        - Environment variables for database credentials improve security and flexibility.
        - Centralized configuration simplifies management.
        - Clearer error messages help with debugging.

    API Routing:
        - Modularized API route handling enhances code organization.
        - Explicit error handling within the API route improves robustness.

    Static File Serving:
        - Error handling during file server setup ensures graceful failures.

    Server Startup:
        - Dynamic port number from environment variable allows for easier deployment.
        - Informative log message indicates the actual port the server is running on.
        - Remember to add error handling for parsing the integer in getIntEnv and replace the placeholder comment with the actual parsing logic.

#### POSTGRESQL quickies
- Open pgsql $psql
- List db $\l
- Use db $psql -d dbname, or $\c dbname
- List tables $\dt
- Describe tables $\d table_name
- To see what's inside of table: $> SELECT * FROM table_name;

#### vue quickies
- npm init vue@latest > npm i > run dev
