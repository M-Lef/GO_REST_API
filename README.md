I developed a rest api in Go which include CRUD operations.

This project was created with three layers :
- The Controller : dealing with the connecting and sending response
- The Service : it's where all the business logic take place
- The Repository : it's where we interract with the database

In the main.go file you will need :
- Change the URI to connect to your MongoDB server
- Change the name of the database and the collection
- You can also change the port number

I used PostMan to test my requests.

The Crud function :
- Create : The body of this request is the DataSet.json file but you must put "file" for the key value.
- Login : To simulate a user trying to access his informations, i put in the body of the request the id on  the first line and his password on the second (the format is a text file). And then after checking that he is using the correct password thanks to bcrypt library, we send all his information via a Json.
- Delete : Nothing to specified
- Read : Nothing to specified
- Update : Because the Update request doesn't exist in the PostMan app, i used the Put request. So you will need to put in the request body a json representing the user you want to update.

All the users files are saved in the User_data folder.

Thank you