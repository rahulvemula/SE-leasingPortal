# Heroku deployment instructions

## URL of heroku deployment

- [https://murmuring-earth-87031.herokuapp.com/] (https://murmuring-earth-87031.herokuapp.com/)

## Changes to be done while you are running your project in local for testing
- In main.go file just change one below line
- log.Fatal(http.ListenAndServe("localhost:9010", r))
- While you want to deploy to heroku server then revert back the above line to below line
- log.Fatal(http.ListenAndServe(":"+port, nil))  *(port := os.Getenv("PORT"))


## Instructions to deploy to heroku

- Follow this link to get the idea about heroku deployment [link](https://www.youtube.com/watch?v=9ea2-J9vCy4)
- Basically we have to initialize a new git repo inside server folder just for heroku. This git is different from our project git repo
- We can also check logs in heroku server in case of any failure

## Api documentation of backend services
### A collapsible section with markdown
<details>
  <summary>User API</summary>
  
  ### GET
  - [http://localhost:9010/user/ ](http://localhost:9010/user/)
   ### GET USER BY EMAIL
  - [http://localhost:9010/user/{email}](http://localhost:9010/user/{email})
  ### POST
  - [http://localhost:9010/user/](http://localhost:9010/user/)
  * Payload
  ``` json
   {
      "name":"vamsi",
      "email":"vbethamsetty@ufl.edu",
      "password": "vamsi"
   }
   ```
  ### PUT
  - [http://localhost:9010/user/](http://localhost:9010/user/)
  * Payload
  ``` json
   {
      "name":"vamsi",
      "email":"vbethamsetty@ufl.edu",
      "password": "vamsi"
   }
   ```
  ### DELETE
  - [http://localhost:9010/user/{id}](http://localhost:9010/user/{id})
</details>
<details>
  <summary>Lease API</summary>
  
  ### GET
  - [http://localhost:9010/lease/ ](http://localhost:9010/lease/)
  ### POST
  - [http://localhost:9010/lease/](http://localhost:9010/lease/)
  * Payload
  ``` json
    {
        "listingId":1,
        "userId":"1",
        "leaseStartDate": "28 Jan",
        "leaseEndDate" : "14 Feb"
    }
   ```
  ### PUT
  - [http://localhost:9010/lease/{leaseId}](http://localhost:9010/)
  * Payload
  ``` json
   {
        "listingId":1,
        "userId":"1",
        "leaseStartDate": "28 Jan",
        "leaseEndDate" : "14 Feb"
    }
   ```
  ### DELETE
  - [http://localhost:9010/lease/{id}](http://localhost:9010/lease/{leaseId})
</details>
<details>
  <summary>Apartment API</summary>
  
  ### GET
  - [http://localhost:9010/apartment/ ](http://localhost:9010/apartment/)
  ### POST
  - [http://localhost:9010/apartment/](http://localhost:9010/apartment/)
  * Payload
  ``` json
    {
        "name":"",
        "address":"3800 SW",
        "amenities": "counter top, new appliances"
    }
   ```
  ### PUT
  - [http://localhost:9010/apartment/{apartmentId}](http://localhost:9010/)
  * Payload
  ``` json
   {
        "name":"",
        "address":"3800 SW",
        "amenities": "counter top, new appliances"
    }
   ```
  ### DELETE
  - [http://localhost:9010/apartment/{id}](http://localhost:9010/apartment/{apartmentId})
</details>
<details>
  <summary>Listing API</summary>
  
  ### GET
  - [http://localhost:9010/listing/ ](http://localhost:9010/listing/)
  ### POST
  - [http://localhost:9010/listing/](http://localhost:9010/listing/)
  * Payload
  ``` json
    {
        "listingType":"bedroom",
        "houseType":"1",
        "rent": 500,
        "userId": 1,
        "isleased": true
    }
   ```
  ### PUT
  - [http://localhost:9010/listing/{listingId}](http://localhost:9010/)
  * Payload
  ``` json
   {
        "listingType":"bedroom",
        "houseType":"1",
        "rent": 500,
        "userId": 1,
        "isleased": true
    }
   ```
  ### DELETE
  - [http://localhost:9010/listing/{id}](http://localhost:9010/listing/{listingId})
</details>
