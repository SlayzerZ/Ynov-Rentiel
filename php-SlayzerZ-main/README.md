[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/YbKxHPdJ)

## All your commands needs to be type in the projct's src folder !

for run the project the depedencies twig and faker are necessary.

These are the required commands :

    $ docker run --rm --interactive --tty   --volume $PWD:/app   composer require fakerphp/faker

    $ docker run --rm --interactive --tty   --volume $PWD:/app   composer require twig/twig:^3.0

## You'll need a PHP version higher than 8.1.0

For that create a Dockerfile file with this lines:


    FROM php:8.2-apache

    RUN a2enmod rewrite

    RUN docker-php-ext-install pdo pdo_mysql

    RUN service apache2 restart




And another one named docker-compose.yml with : 



    version: '3'

    services:

      web:
  
      build: .
    
      ports:
    
        - "8080:80" # Expose port 8080 on WSL to port 80 in the container
      
      volumes:
    
        - ./src:/var/www/html
      

    mysql:
  
      image: mysql:5.7
    
      environment:
    
        MYSQL_ROOT_PASSWORD: my-secret-pw
      
        MYSQL_DATABASE: my_database
      
        MYSQL_USER: my_user
      
        MYSQL_PASSWORD: my_password
      
      volumes:
    
        - db_data:/var/lib/mysql
      
      ports:
    
        - "3306:3306" # Expose port 3306 on the host to port 3306 in the container


    volumes:

      db_data:
  

Both have to be in the project's root (with the README).

After that delete all yours containers, images, volumes and rebuild them with the command below :

        docker-compose up -d --force-recreate

Also you have modify the host setup by yours, else the database not gonna work, for do this :
  
      - In your terminal, run command : "ip addr show eth0"
      
      - Note your local ip near of inet
      
      - In all manager's files replace host=172.20.27.212 BY host="your ip"

## There a Admin Account setup by default
        Email : ibrahimallari84@gmail.com
        Password : ratchet
