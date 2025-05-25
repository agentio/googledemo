
call "google" {
  name     = "Google"
  port     = 4848
  endpoint = "www.googleapis.com"
  auth {
    name       = "authorization"
    location   = "header"
    credential = "session:google"
  }
}

host "YOUR-HOST-NAME" {
  name    = "googledemo"
  backend = "localhost:3000"
  oauth "google" {
    client_id        = "YOUR-CLIENT-ID"
    client_secret    = "YOUR-CLIENT-SECRET"
    authorize_url    = "https://accounts.google.com/o/oauth2/auth"
    access_token_url = "https://accounts.google.com/o/oauth2/token"
    scopes           = ["profile", "email"]
    authorize_parameter {
      name = "approval_prompt"
      value = "force"
    } 
  }
}

