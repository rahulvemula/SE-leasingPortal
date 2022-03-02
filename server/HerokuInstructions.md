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
  - [https://murmuring-earth-87031.herokuapp.com/users ](https://murmuring-earth-87031.herokuapp.com/user/)
   ### GET USER BY EMAIL
  - [https://murmuring-earth-87031.herokuapp.com/users/{email}](https://murmuring-earth-87031.herokuapp.com/user/{email})
  ### POST
  - [https://murmuring-earth-87031.herokuapp.com/users/](https://murmuring-earth-87031.herokuapp.com/user/)
  * Payload
  ``` json
   {
      "name":"vamsi",
      "email":"vbethamsetty@ufl.edu",
      "password": "vamsi"
   }
   ```
  ### PUT
  - [https://murmuring-earth-87031.herokuapp.com/users/](https://murmuring-earth-87031.herokuapp.com/users/{userId})
  * Payload
  ``` json
   {
      "name":"vamsi",
      "email":"vbethamsetty@ufl.edu",
      "password": "vamsi"
   }
   ```
  ### DELETE
  - [https://murmuring-earth-87031.herokuapp.com/users/{id}](https://murmuring-earth-87031.herokuapp.com/user/{id})
</details>
<details>
  <summary>Lease API</summary>
  
  ### GET
  - [https://murmuring-earth-87031.herokuapp.com/leases ](https://murmuring-earth-87031.herokuapp.com/lease/)
  ### POST
  - [https://murmuring-earth-87031.herokuapp.com/leases](https://murmuring-earth-87031.herokuapp.com/lease/)
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
  - [https://murmuring-earth-87031.herokuapp.com/leases/{leaseId}](https://murmuring-earth-87031.herokuapp.com/)
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
  - [https://murmuring-earth-87031.herokuapp.com/leases/{id}](https://murmuring-earth-87031.herokuapp.com/lease/{leaseId})
</details>
<details>
  <summary>Apartment API</summary>
  
  ### GET
  - [https://murmuring-earth-87031.herokuapp.com/apartments ](https://murmuring-earth-87031.herokuapp.com/apartment/)
  ### POST
  - [https://murmuring-earth-87031.herokuapp.com/apartments](https://murmuring-earth-87031.herokuapp.com/apartment/)
  * Payload
  ``` json
    {
        "name":"",
        "address":"3800 SW",
        "amenities": "counter top, new appliances"
    }
   ```
  ### PUT
  - [https://murmuring-earth-87031.herokuapp.com/apartments/{apartmentId}](https://murmuring-earth-87031.herokuapp.com/)
  * Payload
  ``` json
   {
        "name":"",
        "address":"3800 SW",
        "amenities": "counter top, new appliances"
    }
   ```
  ### DELETE
  - [https://murmuring-earth-87031.herokuapp.com/apartments/{id}](https://murmuring-earth-87031.herokuapp.com/apartment/{apartmentId})
</details>
<details>
  <summary>Listing API</summary>
  
  ### GET
  - [https://murmuring-earth-87031.herokuapp.com/listings ](https://murmuring-earth-87031.herokuapp.com/listing/)
  ### POST
  - [https://murmuring-earth-87031.herokuapp.com/listings](https://murmuring-earth-87031.herokuapp.com/listing/)
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
  - [https://murmuring-earth-87031.herokuapp.com/listings/{listingId}](https://murmuring-earth-87031.herokuapp.com/)
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
  - [https://murmuring-earth-87031.herokuapp.com/listings/{id}](https://murmuring-earth-87031.herokuapp.com/listing/{listingId})
</details>
