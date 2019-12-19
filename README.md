# inga

# Client Developers
For now:
make a POST request.
From the command line, using `curl`:
```
curl -v --data 'apikey=<APIKEY>&product=<PRODUCT>&version=<VERSION>&uuid=<UUID>&eventID=<EVENTID>&dateTime=<DATETIME>' http://inga.shef.ac.uk/api/v201910/
```
This will return a 200 response.

or from Python using the [`requests`](https://requests.kennethreitz.org/en/master/) library:

```
import requests

apikey = "apikey-supplied-by-us"
product = "your-product-name"
version = "your-version-number"
uuid = "uniquely-identify-the-user-installation"
eventID = "your-event-identifier"
dateTime = "a-date-time-stamp"

data = dict(apikey=apikey, product=product, version=version, uuid=uuid, eventID=eventID, dateTime=dateTime)

requests.post("https://inga.shef.ac.uk/api/v201910/", data=data)
```

The following key=value pairs are passed to **inga**.
* `APIkey` - will be granted on request.
* `PRODUCT` - is up to the vendor, i.e. the name of your software package.
* `VERSION` - your application's version. (can be left blank)
* `UUID` - again, up to the vendor. We recommend it is used as a unique user ID per installation, user etc.
* `EVENTID` - use as an application run ID or launch ID, whatever you prefer, or leave blank
* `DATETIME` - generate a date/time stamp for the ping for **inga** to log.

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
