# date app

golang back-end; react front-end
- see jimnotes/dateapp notes for URS
- want a challenge?

## TODO

- [ ] api recipes available for users
- [ ] api for admin?
- [ ] set up db (postgres)
- [ ] matching algorithm using ML/AI
- [ ] redifine Recipe struct (backend/cmd/api) to include details for user admin & matching
- [ ] scrape data from ladders, present info with better interface, and allow users to do something on the site

### changelog

- looking into mongodb over postgres. See terms&conditions for details. 21aug2024 ... decided on postgres
- friend needs an inventory management system. 21aug2024
    - dbname is inventorydb

#### POSTGRESQL quickies
- Open pgsql $psql
- List db $\l
- Use db $psql -d dbname, or $\c dbname
- List tables $\dt
- Describe tables $\d table_name
- To see what's inside of table: $> SELECT * FROM table_name;

#### vue quickies
- npm init vue@latest > npm i > run dev
