# Weather API go client

[weatherapi.com](https://www.weatherapi.com/)

## Testing

There is no HTTP mocks in the tests. Tests are to used to validate that the JSON decoding is still working from the live API. Tests should be run periodically to validate that the decoding still works. The Weather API is subject to change their response JSON schema.

To run tests, you must provide a valid Weather API key in a text file. After creating a free account, find the API key [here](https://www.weatherapi.com/my/). Create a new text file called `apikey.txt` and paste the API key into it.
