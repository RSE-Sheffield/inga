# inga
inga's Not Google Analytics

# Client Developers
For now: make an HTTP GET request to: `https://inga.shef.ac.uk/api/v201910/?api=<APIKEY>&product=<PRODUCT>&version=<VERSION>&uuid=<UUID>`
This will return a 200 response.

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
