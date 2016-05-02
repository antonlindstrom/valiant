# valiant

This is a tool that takes a configuration file and sends the request to an
upstream HTTP server and validates the response.

## Usage

Execute tests:

	valiant execute --address=http://example.com

You can automatically update responses if you see the expected output but
didn't update the test file:

	valiant update --address=http://example.com --test-file=tests/00_example.yml

Generate an example test file:

	valiant generate --test-file=tests/00_generated_example.yml
