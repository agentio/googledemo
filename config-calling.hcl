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
