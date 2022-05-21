# Resource-Management-App

![h25](https://user-images.githubusercontent.com/11625672/168671979-39c58c99-3cd2-45b5-bda0-bea722563289.png)


# Development

In order to execute the program, you need to have `psql` installed in your system:

Using Linux env:

`sudo apt install postgresql` -> installs the `psql`

`sudo -u postgres psql` -> sudo into user name as postgress

`psql --version` -> checks version

`export RESC_DB_DSN='postgres://USER_NAME:YOUR_PWD@localhost/DB_NAME'`
`psql $RESC_DB_DSN `

`psql --host=localhost --dbname=DB_NAME --username=USER_NAME`

Then exec into that and create the table as below

```
CREATE TABLE IF NOT EXISTS resources (
  id SERIAL PRIMARY KEY NOT NULL,
  title varchar(255) NOT NULL,
  category varchar(32) NOT NULL,
  status varchar(15) NOT NULL,
  types varchar(10) NOT NULL,
  content text,
  file_link varchar(255) DEFAULT NULL,
  created_by int NOT NULL,
  created_at bigint NOT NULL,
  updated_by int NOT NULL,
  updated_at bigint NOT NULL
);
```

Steps to run the source code:

`go mod tidy` -> Download the required dependencies

`go run main.go` -> Runs the App

# 🧠 Inspiration
Inspired by CERN Collaboration tracks and generic functions

# 💻 What it does
Our API enables you to make use of managing and accessing your resources for storing PDF, TXT files and images [Blogs]

# 🔨 How we built it
Built around Gin-Gonic framework

# 🏃🏻 Challenges we ran into
How to configure our service api with integrating generic function

# 📌 Accomplishments that we're proud of
API configuration and Generic function and types

# 📖 What we learned
Go Generic function and its use of `any` type

# What's next for Resource Management APP
To integrate with UI and setup kubernetes cluster



## ScreenShots

![h22](https://user-images.githubusercontent.com/11625672/168672115-6b7dab29-4dc0-4a62-a07c-f118fdcc15df.png)

![h23](https://user-images.githubusercontent.com/11625672/168672171-e9538431-069f-4e54-bae0-2df5b81b7d73.png)

![h24](https://user-images.githubusercontent.com/11625672/168672228-154baac7-9ff2-4a7d-8cd0-94de9fb97d66.png)
