package concourse

type Source struct {
  ConsumerKey string `json:"consumer_key"`
  ConsumerSecret string `json:"consumer_secret"`
  AccessToken string `json:"access_token"`
  AccessTokenSecret string `json:"access_token_secret"`
}

type Version struct {
  TweetId string `json:tweet_id,omitempty`
}

type MetadataPair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type OutParams struct {
  Status string `json:"status"`
  Media []string `json:"media,omitempty"`
}

type OutRequest struct {
  Source Source `json:"source"`
  Params OutParams `json:"params"`
}

type OutResponse struct {
  Version Version `json:"version"`
  Metadata []MetadataPair `json:"metadata"`
}
