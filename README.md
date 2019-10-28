# inga
inga's Not Google Analytics

# Client Developers
For now, either:
* make an HTTP GET request to: `https://inga.shef.ac.uk/api/v201910/?apikey=<APIKEY>&product=<PRODUCT>&version=<VERSION>&uuid=<UUID>`.
* or, make a POST request, e.g.
```
curl -v --data 'apikey=<APIKEY>&product=<PRODUCT>&version=<VERSION>&uuid=<UUID>' http://inga.shef.ac.uk/api/v201910/
```
These will return a 200 response.

The following key=value pairs are passed to inga.
* `APIkey` - will be granted on request.
* `PRODUCT` - is up to the vendor, i.e. the name of your software package.
* `VERSION` - your application's version.
* `UUID` - again, up to the vendor. We recommend it is used as a unique user ID per installation, user etc.

We will expect you to bake the API key into your application. We reserve the right to revoke API keys if they are abused and you may need to refresh your API key and application if so.

# Development (inga-server)

Install `make` and the Go tools.

To build:

    make

To run:

    ./inga
