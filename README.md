**Studio Classes and Booking API**

**Overview**
This project is a simple RESTful API built with Go and the Gin framework. It allows a studio owner to manage classes and members to book classes. The data is stored in an in-memory data store for simplicity, making it ideal for testing and development.



**Features**
- Create new classes with details like name, start date, end date, and capacity.
- View all available classes.
- Book a class for a specific date.


**API Endpoints**
**1. Create a Class**
**Endpoint**: `/classes`  
**Method**: `POST`  
**Query Parameters**:  
- `class_name` (string, required): Name of the class.  
- `start_date` (string, required): Start date of the class (format: `YYYY-MM-DD`).  
- `end_date` (string, required): End date of the class (format: `YYYY-MM-DD`).  
- `capacity` (int, required): Maximum number of attendees per class.  

**Response**:  
- `201 Created`: Class created successfully.  
- `400 Bad Request`: Missing or invalid parameters.  

---

**2. Get All Classes**
**Endpoint**: `/getclasses`  
**Method**: `GET`  



**3. Book a Class**
**Endpoint**: `/bookings`  
**Method**: `GET`  
**Query Parameters**:  
- `name` (string, required): Name of the member booking the class.  
- `date` (string, required): Date of the booking (format: `YYYY-MM-DD`).  



**Installation**
1. Extract the zip file:
   
   

2. Install dependencies:
   run below commands in terminal
   go mod tidy

3. Run the server:
   then run this command in terminal
   go run main.go



**Project Structure**
```
├── controllers/       # Contains API controllers for handling requests
├── models/            # Contains data models for classes and bookings
├── services/          # Business logic for classes and bookings
├── router/            # Route definitions
├── main.go            # Entry point of the application
├── README.md          # Project documentation
```



**Testing**
Use tools like `Postman` or `cURL` to test the API endpoints. Example `cURL` command to create a class:
```bash
curl -X POST "http://localhost:8080/classes?class_name=Yoga&start_date=2025-01-20&end_date=2025-01-30&capacity=15"
```

---

**Future Enhancements**
- Add authentication for endpoints.
- Introduce a database for persistent storage.
- Implement booking capacity limits.
- Add some more business logic


