# Cities

A small project that works with storing data in a file of type CSV.

The application writes data from a file to RAM and then works with it in this form. All requests are made through API and JSON data type.

### Functions provided by this application:
1) Create a new city entry
2) Getting data about the city through ID
3) Deleting information about the city according to the specified ID
4) Update information about the population of the city for the specified id;
5) Getting a list of cities for the specified region;
6) Getting a list of cities for the specified district;
7) Get a list of cities for the specified population range;
8) Get a list of cities in the specified range of the year of foundation.

## Launch

You can start the service with a specific address or leave it as default - localhost:4000. 
`go run ./cmd/srv`
