# Twitter [Concourse](http://concourse.ci) Resource

Send tweets when your [Concourse](http://concourse.ci) builds finish.

## Source Configuration

### Required
* `consumer_key`: The consumer key from registering an application with
  [Twitter](https://apps.twitter.com).
* `consumer_secret`: The consumer secret associated with the key above
* `access_token`: An OAuth 1.0a access token for the twitter account that will
  be posting tweets.
* `access_token_secret`: The secret associated with the access token above.

## Behavior

### `check`, `in`

Currently this resource only supports the `put` phase of a job plan, so these
are effectively no-ops. This will likely change in the future.

### `out`: Post a tweet

Posts a tweet with the given parameters. Note that no validation is done on the
parameters to ensure that they meet Twitter's specifications; that is left to
the API calls to Twitter itself, and builds will fail if the data does not
conform.

#### Parameters

##### Required:
* `status`: The text of the tweet itself. Any
  [metadata](http://concourse.ci/implementing-resources.html#resource-metadata)
  in the status will be evaluated prior to sending the tweet. Use `media` to
  include references to photos or video

##### Optional:
* `media`: An array of paths to images or video to upload. Any media referenced
  here will be uploaded and referenced by the resulting tweet.
