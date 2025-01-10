--> Run the application: <br />
    . "go run cmd/Rest_APIs_In_Golang/main.go -config config/local.yaml". <br />
    . go to browser "http://localhost:8082/api/students". Provides the list of all the students. <br /> <br />

--> Use of Postman: <br />
    1. Type: POST <br />
    2. Inside body write the below Json formatted data and send <br />
    { <br />
    \t "name": "Kumar", <br />
    \t "email": "kumar@gmail.com", <br />
    \t "age": 25 <br />
    } <br />
    3. 2nd Type: GET, "http://localhost:8082/api/students/id" where id = 1,2,3... till no. of students in the database. <br />
    4. 3rd Type GET, "http://localhost:8082/api/students" it provides details of all the students in the database. <br /> <br />

--> Reset database: <br />
    1. Goto "/storage" <br />
    2. Delete the file named "storage.db". <br />
    3. Create another file named "storage.db". <br />
    4. Start the application. <br /> <br />

--> TablePlus, it can be used to observe the database creation and data manipulation inside the database. <br /> <br />

--> Using SQLite which is very fast because it stores the data in our local machine. <br />
