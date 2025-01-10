--> Run the application:
    . "go run cmd/Rest_APIs_In_Golang/main.go -config config/local.yaml".
    . go to browser "http://localhost:8082/api/students". Provides the list of all the students.

--> Use of Postman:
    1. Type: POST
    2. Inside body write
    {
        "name": "Kumar",
        "email": "kumar@gmail.com",
        "age": 25
    }
    3. 2nd Type: GET, "http://localhost:8082/api/students/id" where id = 1,2,3... till no. of students in the database.
    4. 3rd Type GET, "http://localhost:8082/api/students" it provides details of all the students in the database.

--> Reset database:
    1. Goto "/storage"
    2. Delete the file named "storage.db".
    3. Create another file named "storage.db".
    4. Start the application.

--> TablePlus, it can be used to observe the database creation and data manipulation inside the database.

--> Using SQLite which is very fast because it stores the data in our local machine.