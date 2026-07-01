calling "google" {
  name   = "Google Identity API"
  target = "www.googleapis.com"
  port   = 4848
  apply_header "authorization" {
    secret = "session:google"
  }
  operation "v1-userinfo" {
    method      = "GET"
    path        = "/oauth2/v1/userinfo"
    description = "Get user info"
  }
  operation "v3-userinfo" {
    method      = "GET"
    path        = "/oauth2/v3/userinfo"
    description = "Get user info"
  }
}

ingress "YOUR-HOST-NAME" {
  name    = "googledemo"
  backend = "nomad:googledemo"
  oauth_client "google" {
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
