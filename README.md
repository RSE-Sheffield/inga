# inga
**inga**'s Not Google Analytics

# Client Developers
For now:
make a POST request.
From the command line, using `curl`:
```
curl -v --data 'apikey=<APIKEY>&product=<PRODUCT>&version=<VERSION>&uuid=<UUID>' http://inga.shef.ac.uk/api/v201910/
```
This will return a 200 response.

or from Python using the [`requests`](https://requests.kennethreitz.org/en/master/) library:

```
import requests

apikey = "apikey-supplied-by-us"
product = "your-product-name"
version = "your-version-number"
uuid = "uniquely-identify-the-user-installation"

data = dict(apikey=apikey, product=product, version=version, uuid=uuid)

requests.post("https://inga.shef.ac.uk/api/v201910/", data=data)
```

The following key=value pairs are passed to **inga**.
* `APIkey` - will be granted on request.
* `PRODUCT` - is up to the vendor, i.e. the name of your software package.
* `VERSION` - your application's version.
* `UUID` - again, up to the vendor. We recommend it is used as a unique user ID per installation, user etc.

We will expect you to bake the API key into your application. We reserve the right to revoke API keys if they are abused and you may need to refresh your API key and application if so.

## Requesting your users' consent to collect data
Whilst **inga** does not collect any personalised or identifiable data, it is essential that your application gets the user's consent prior to sending data. We recommend that you use a statement such as:
> **Consent to collect anonymised usage data**
> In order to help this project obtain future funding for research and development, it is helpful for us to collect a small amount of anonymous, non-identifiable data. This is only done at {installation,launch,login} and will not disrupt the normal functioning of the application. We use a service called **inga** to aggregate numbers of {installations,launches,logins}, their date and time, and country of origin. If you do not consent to data collection, no information will be stored.

# Data Protection Statement
> TBD

# Development (inga-server)

Install `make` and the Go tools.

To build:

    make

To run:

    ./inga
