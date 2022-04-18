# Sprint 4 - Summary of tasks achieved and Bonus

## UI Tasks achieved


## Backend Tasks achieved

- Added Swagger documentation for login apis
- Added Swagger documentation for lease apis
- Added Swagger documentation for listing apis
- Added Swagger documentation for apartment apis
- Added Swagger documentation for society apis
- UNIT tests written for society and listing apis
- Deployed services to remote server in HEROKU
- Integrated CICD to heroku with our current git repo
- Documented all remote server apis

## BONUS

- Deployed the app UI to remote server
- Deployed the app backend services to remote server

## Useful links of the project

- [Easy-Lease Repo Link](https://github.com/rahulvemula/SE-leasingPortal)
- [Discussions link on git](https://github.com/rahulvemula/SE-leasingPortal/discussions)
- [Actions link on git](https://github.com/rahulvemula/SE-leasingPortal/actions)
- [Sprint 4 User stories progress board link](https://github.com/rahulvemula/SE-leasingPortal/projects/5)
- [All user stories link](https://github.com/rahulvemula/SE-leasingPortal/issues)

## Sprint 4 Demo UI


## Backend Demo video(swagger)
![Alt text](Screenshots/SwaggerDocForBackend.gif?raw=true "Backend demo swagger")

## Backend GoLang unit tests demo
![Alt text](Screenshots/BackendUnitTests.gif?raw=true "Back end Unit tests")

## Cypress automation tests Demo


## Screenshots for unit tests


#### UI Components tests



#### Go lang unit tests


## Api documentation of backend services

### A collapsible section with markdown

<details>
  <summary>User API</summary>
  
  ### GET ALL USERS
  - [http://murmuring-earth-87031.herokuapp.com/users ](http://murmuring-earth-87031.herokuapp.com/users)
  ### GET USER BY EMAIL
  - [http://murmuring-earth-87031.herokuapp.com/users/{email}](http://murmuring-earth-87031.herokuapp.com/users/{email})
  ### CREATE A USER
  - [http://murmuring-earth-87031.herokuapp.com/users](http://murmuring-earth-87031.herokuapp.com/users)
  * Payload
  ``` json
   {
      "name":"vamsi",
      "email":"vbethamsetty@ufl.edu",
      "password": "vamsi"
   }
   ```
  ### UPDATE AN USER
  - [http://murmuring-earth-87031.herokuapp.com/users/{userId}](http://murmuring-earth-87031.herokuapp.com/users/{userId})
  * Payload
  ``` json
   {
      "name":"vamsi",
      "email":"vbethamsetty@ufl.edu",
      "password": "vamsi"
   }
   ```
  ### DELETE AN USER
  - [http://murmuring-earth-87031.herokuapp.com/users/{id}](http://murmuring-earth-87031.herokuapp.com/users/{id})
</details>
<details>
  <summary>Lease API</summary>
  
  ### GET  ALL LEASES
  - [http://murmuring-earth-87031.herokuapp.com/leases ](http://murmuring-earth-87031.herokuapp.com/leases)
  ### CREATE A LEASE
  - [http://murmuring-earth-87031.herokuapp.com/leases](http://murmuring-earth-87031.herokuapp.com/leases)
  * Payload
  ``` json
    {
        "listingId":1,
        "userId":"1",
        "leaseStartDate": "28 Jan",
        "leaseEndDate" : "14 Feb"
    }
   ```
  ### UPDATE A LEASE
  - [http://murmuring-earth-87031.herokuapp.com/leases/{leaseId}](http://murmuring-earth-87031.herokuapp.com/leases/{leaseId})
  * Payload
  ``` json
   {
        "listingId":1,
        "userId":"1",
        "leaseStartDate": "28 Jan",
        "leaseEndDate" : "14 Feb"
    }
   ```
  ### DELETE A LEASE
  - [http://murmuring-earth-87031.herokuapp.com/leases/{leaseId}](http://murmuring-earth-87031.herokuapp.com/leases/{leaseId})
</details>
<details>
  <summary>Apartment API</summary>
  
  ### GET ALL APARTMENTS
  - [http://murmuring-earth-87031.herokuapp.com/apartments ](http://murmuring-earth-87031.herokuapp.com/apartments)
  ### CREATE AN APARTMENT
  - [http://murmuring-earth-87031.herokuapp.com/apartments](http://murmuring-earth-87031.herokuapp.com/apartments)
  * Payload
  ``` json
    {
        "name":"",
        "address":"3800 SW",
        "amenities": "counter top, new appliances"
    }
   ```
  ### UPDATE AN APARTMENT
  - [http://murmuring-earth-87031.herokuapp.com/apartments/{apartmentId}](http://murmuring-earth-87031.herokuapp.com/apartments/{apartmentId})
  * Payload
  ``` json
   {
        "name":"",
        "address":"3800 SW",
        "amenities": "counter top, new appliances"
    }
   ```
  ### DELETE AN APARTNMENT
  - [http://murmuring-earth-87031.herokuapp.com/apartments/{apartmentId}](http://murmuring-earth-87031.herokuapp.com/apartments/{apartmentId})
</details>
<details>
  <summary>Listing API</summary>
  
  ### GET ALL LISTINGS
  - [http://murmuring-earth-87031.herokuapp.com/listings ](http://murmuring-earth-87031.herokuapp.com/listings)
  ### CREATE A LISTING
  - [http://murmuring-earth-87031.herokuapp.com/listings](http://murmuring-earth-87031.herokuapp.com/listings)
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
  ### UPDATE A LISTING
  - [http://murmuring-earth-87031.herokuapp.com/listings/{listingId}](http://murmuring-earth-87031.herokuapp.com/listings/{listingId})
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
  ### DELETE A LISTING
  - [http://murmuring-earth-87031.herokuapp.com/listings/{listingId}](http://murmuring-earth-87031.herokuapp.com/listing/{listingId})
</details>
