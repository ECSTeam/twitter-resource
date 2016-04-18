package main

import (
  "encoding/base64"
  "io/ioutil"
  "net/url"
  "os"
  "path"

  "github.com/ECSTeam/twitter-resource/concourse"
  "github.com/ChimeraCoder/anaconda"
)

func main() {
  if len(os.Args) < 2 {
    concourse.Fatal("Missing required working dir arg!")
  }

  workingDir := os.Args[1]

  var request concourse.OutRequest

  concourse.ReadRequest(&request)

  anaconda.SetConsumerKey(request.Source.ConsumerKey)
  anaconda.SetConsumerSecret(request.Source.ConsumerSecret)

  api := anaconda.NewTwitterApi(request.Source.AccessToken,
    request.Source.AccessTokenSecret)

  uploadedMedia := url.Values{}

  for _, imageFile := range request.Params.Media {
    if dir, direrr := os.Getwd(); direrr == nil {
      concourse.Sayf("Working in dir %v\n", dir)
    }

    imageFile = path.Join(workingDir, imageFile)

    concourse.Sayf("Uploading file %v\n", imageFile)
    bytes, fileErr := ioutil.ReadFile(imageFile)
    if fileErr != nil {
      concourse.Fatal("Error reading file: %v\n", fileErr)
    }

    if media, err := api.UploadMedia(base64.StdEncoding.EncodeToString(bytes)); err != nil {
      concourse.Fatal("Error uploading media: %v\n", err)
    } else {
      uploadedMedia.Add("media_ids", media.MediaIDString)
    }

    concourse.Sayf("Upload of %v complete\n", imageFile)
  }

  // expand any variables
  statusText := os.ExpandEnv(request.Params.Status)
  concourse.Sayf("Posting tweet '%s'\n", statusText)

  output := concourse.OutResponse{}
  if tweet, err := api.PostTweet(statusText, uploadedMedia); err != nil {
    concourse.Fatal("Error posting tweet: %v\n", err)
  } else {
    output.Version = concourse.Version{
      TweetId: tweet.IdStr,
    }
  }

  concourse.WriteResponse(output)
}
