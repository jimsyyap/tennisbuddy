# tennis buddy

- Platform for rec tennis players in local to connect, play, and compete
- golang back-end; vue.js front-end
- the utr app seems similar, but it doesnt have reliable doubles stats -> opportunity?
    - how to verify score?
    - start with blackburn ladders

## TODO

- [ ] main.go is too big. Break it up into smaller files and folders
- [ ] start with ladders ... build from there
- [ ] create user
- [ ] create user dashboard
- [ ] create users admin dashboard
- [ ] which info for each user?
- [ ] api for admin?
- [ ] api for users
- [ ] users can set up socials
- [ ] users can join socials
- [ ] users can setup matches/comp
- [ ] users can join matches/comp
- [ ] users can view matches/comp results
- [ ] set up tables (postgres)
- [ ] matching algorithm using ML/AI
- [ ] social-media/forum/chat
- [ ] video streaming
- [ ] marketplace to sell, auction with bidding
- [ ] membership dashboard
- [ ] where to get stats, how to compute player ratings, rank
- [ ] also do for paa
- [ ] see how you can use this https://github.com/PacktPublishing/Full-Stack-React-Projects-Second-Edition/tree/master
- [ ] scrape data from ladders, present info with better interface, and allow users to do something on the site

### changelog

- 26aug2024 - frontend to react.js, but keeping vue_this folder just in case.
- created helloworld app.js...connected to postgres == joy
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
- why not use vue on parts you need to learn?

