# Sprint 1 - Summary of tasks achieved

## Project - Overview

The idea is to provide a web-application to lease out the apartments where customers can find a suitable room to sign a lease and manage their lease.

## Useful links of the project
- [Easy-Lease Repo Link](https://github.com/rahulvemula/SE-leasingPortal) 
- [Discussions link on git](https://github.com/rahulvemula/SE-leasingPortal/discussions)
- [Actions link on git](https://github.com/rahulvemula/SE-leasingPortal/actions)
- [Sprint 1 User stories progress(JIRA) board link](https://github.com/rahulvemula/SE-leasingPortal/projects/1)
- [All user stories link](https://github.com/rahulvemula/SE-leasingPortal/issues)

## UI Tasks achieved - (React js)
- Performed the initial setup required for the frontend
- Added the Registration screen for the user
- Added the Login screen for the user
- Provision to create a dummy lease for now for the user
- Provision to show the listing of leases for the user

## Backend Tasks achieved - (Go-lang)
- Created API's for Lease CRUD operations
- Created API's for User Login/Registration CRUD operations
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
