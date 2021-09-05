// package holds a number of constants used within the application.
package constant

var PORT = "9091"
var PREFIX = "MARVEL_"
var MARVEL_URL = "https://gateway.marvel.com/v1/public/characters"
var SCHEME = "https"
var TTL = "In TTL caching strategy, the cache keys will automatically expire after the configured TTL in config.yml configuration. Subsequent API requests after the cache expiry will incur the cache miss and re-populate the cache."
var PREFETCH = "In PREFETCH caching strategy, the cache keys will have no cache expiry and cache will automatically be populated by pre-fetching with latest copy of the data after the configured time in config.yml."
