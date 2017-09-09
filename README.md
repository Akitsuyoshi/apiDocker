# apiDocker_Golang with DB

This api is written in golang, using with mongoDB.


### prerequisite
1. installed Docker in local machine or VM
2. somewhat familiarity with Docker

### Here are tha instruction to use this api

1 Move your project folder

`$ cd project_directory/`

2 Install this repo, and to apiDocker

```
$ git install https://github.com/Akitsuyoshi/apiDocker
$ cd apiDocker
```

3 Compile this project images, and run the application

`$ docker-compose build && docker-compose up`

If you hit http://localhost:3000/products in the browser, you'll see blank JSON data.


this api is instructed in [my blog](https://medium.com/p/e63fb2631bfc/edit)
