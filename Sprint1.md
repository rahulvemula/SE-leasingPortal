# Sprint 1 - Summary of tasks achieved

## Project - Overview

The idea is to provide a web-application to lease out the apartments where customers can find a suitable room to sign a lease and manage their lease.

## Useful links of the project
- [Easy-Lease Repo Link](https://github.com/rahulvemula/SE-leasingPortal) 
- [Discussions link on git](https://github.com/rahulvemula/SE-leasingPortal/discussions)
- [Actions link on git](https://github.com/rahulvemula/SE-leasingPortal/actions)
- [Sprint 1 User stories progress board link](https://github.com/rahulvemula/SE-leasingPortal/projects/1)
- [All user stories link](https://github.com/rahulvemula/SE-leasingPortal/issues)

## UI Tasks achieved - (React js)
- Performed the initial setup required for the frontend
- Added home page to show all available societies
- Added listings page to show houses available at a particular society
- Provision to show a dummy lease confirmation for now
- Provision to show the listings of properties available

## UI Demo
[Link to video](https://www.youtube.com/watch?v=g1fCuswK1mo)

## Backend Demo
[Link to video](https://www.youtube.com/watch?v=f7cEDfhIKnI)

## Backend Tasks achieved - (Go-lang)
- Created API's for Lease CRUD operations
- Created API's for User Registration CRUD operations
- Created API's for Apartment CRUD operations
- Created API'S for Listing Lease CRUD operations

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
