# Sprint 4 - Summary of tasks achieved and Bonus

## Description of application
The application leasingPortal is a web application developed using react Js and goLang, which lets customers explore different available housing options and lease one of them.
The functionalities our application provides are as follows:

- **Register:** For first-timers, the app provides a UI, takes in the user name, email, and password, and creates an account.
- **Login:**  Once the users are registered, they can log in using email and password. The details are validated using login API.
- **Property Listing:** Every client, irrespective of registration can browse through all available properties in a given city
- **Housing Listing:** For a given property, the application lists various types of housing options
- **Lease:** 
  - Signing the lease: Once a house is selected, the customer can lease it by providing all required details and accepting the given terms and conditions.
  - Lease Document: Once the lease is confirmed, a lease document is generated, which can be downloaded in a PDF form
- **Support:** 
  A client can register a complaint with the property from the support page
- **About Us:** 
  Our application also has an About page which tell users about the people responsible.
  
  
## Group Members
SNO | Name                          | Github username| Type of development|
--- | -------------                 |:-------------: | :------------------:
1   | Vamsi Viswanath Bethamsetty   | vamsibvv       | Backend (Go lang)  |
2   | Rahul Vemula                  | rahulvemula    | Frontend(React)    |
3   | Lahari Barad                  | lbarad         | Frontend(React)    |
4   | Mitul Mandaliya               | mitulmandaliya | Backend (Go lang)  |


## UI Tasks achieved
- Deployed the app to github pages
- Added a support page to let users register a complaint
- Added a terms and conditions page, with out accepting these, a lease can not be signed
- Added lease details to redux store to be able to share between components
- Added cypress tests for support page and terms and conditions page
- Added component tests for new components
- Bug fixes


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
- Deployed App Url: [https://rahulvemula.github.io/SE-leasingPortal/](https://rahulvemula.github.io/SE-leasingPortal/)
### Bonus Details
- Deployed the app UI to remote server using git hub pages
- Deployed the app backend services to remote server using heroku apis.  
- Backendservies-domain: [https://murmuring-earth-87031.herokuapp.com/]
- [Remote-apis-documentation-link](#api-documentation-of-backend-services)         

## Useful links of the project

- [Easy-Lease Repo Link](https://github.com/rahulvemula/SE-leasingPortal)
- [Discussions link on git](https://github.com/rahulvemula/SE-leasingPortal/discussions)
- [Actions link on git](https://github.com/rahulvemula/SE-leasingPortal/actions)
- [Sprint 4 User stories progress board link](https://github.com/rahulvemula/SE-leasingPortal/projects/5)
- [All user stories link](https://github.com/rahulvemula/SE-leasingPortal/issues)

## Sprint 4 UI Demo
![UI Demo](Screenshots/Sprint4UIDemo.gif?raw=true "UI demo")

## Backend Demo video(swagger)
![Alt text](Screenshots/SwaggerDocForBackend.gif?raw=true "Backend demo swagger")

## Backend GoLang unit tests demo
![Alt text](Screenshots/BackendUnitTests.gif?raw=true "Back end Unit tests")

## Cypress automation tests Demo
![cypress](Screenshots/finalCypressTests.gif?raw=true)

## Screenshots for unit tests
![Alt text](Screenshots/component-tests.png?raw=true "UI Component tests")

## Go lang unit tests
![Alt text](Screenshots/GoLangUnitTestsSpring4.png?raw=true "Back end Unit tests")

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
