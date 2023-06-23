
## Installation
> ### Go-migrate
> #### MacOS
> ```bash
>$ brew install golang-migrate
>```
> #### Linux
> ```bash
>$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
>$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
>$ apt-get update
>$ apt-get install -y migrate
>```
> #### Windows: Using [scoop](https://scoop.sh/)
> ```bash
>$ scoop install migrate
>```
> ### Download pre-built binary (Windows, MacOS, or Linux)
>[Release Downloads](https://github.com/golang-migrate/migrate/releases)
>```bash
>$ curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$os-$arch.tar.gz | tar xvz
>```

---

### CLI
> Create migration
> ```shell
> make NAME="init_db" migrate-create 
> ``` 


> Up migration
> ```shell
> make migrate-up
>``` 

> Rollback migration
> ```shell
> migrate migrate-down
>```

> Drop DB
> ```shell
> migrate drop-db
>```