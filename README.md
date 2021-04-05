# PJWT : A simple JWT implementation

PJWT is a very tiny golang package to implement and verify JWT Token quickly.
This is not a library for full manipulation JWT and __do not implements__ full RFC specifications.
PJWT can just create, valid and extract playload from token. Algorithm for signature is `HMAC512`, you can't change it (except if you modify code btw).

[jwt.io](https://jwt.io) was used during dev to test tokens.

### What pjwt can :
- create signed token for your app whith `HMAC512`
- verify signature and integrity datas from given token into your app
- extract playload from token direct into json struct

### What pjwt can't :
- create unsigned token
- use public/private key for encryption
- explore full JWT implementation or subtilities
- verify expiration date or whatever. You have to control by yourself playload datas

# Examples

## - Fill playload and create token
```golang
type PlayloadExample struct {
	User 		string 	`json:"user"`
	Admin 		bool 	`json:"admin"`
	ExpiresAt	int64 	`json:"exp"`
}

key := []byte( "your-256-bit-secret" )
playload := PlayloadExample{
    User: "toto",
    Admin: false,
    // ExpiresAt: time.Now().Add( 24 * time.Hour ).Unix()
    ExpiresAt: 123456 // for README.md example
}
pjwt := new( PJWToken )

tk, err := pjwt.CreateToken( &playload, key )
if err != nil { panic( err ) }

fmt.Println( tk )
```

Will done :

```
eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoidG90byIsImFkbWluIjpmYWxzZSwiZXhwIjoxMjM0NTZ9.hk3aPoxn2X8gN0ZJYFZhom1tBsOKihiu6FRmEIfg1wF1Kikn8-o86E0OyiLMEz8Xn1LyEP-mNSC_z-8L1YpcKA
```

[Try it on jwt.io !](https://jwt.io/#debugger-io?token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoidG90byIsImFkbWluIjpmYWxzZSwiZXhwIjoxMjM0NTZ9.hk3aPoxn2X8gN0ZJYFZhom1tBsOKihiu6FRmEIfg1wF1Kikn8-o86E0OyiLMEz8Xn1LyEP-mNSC_z-8L1YpcKA)

## - Verify And Extract playload from given token
```golang
type PlayloadExample struct {
	User 		string 	`json:"user"`
	ExpiresAt	int64 	`json:"exp"`
	Rights		struct {
		Read 	bool 	`json:"read"`
		Write 	bool 	`json:"write"`
		Execute	bool 	`json:"execute"`
	} `json:"rights"`
}

key := []byte( "your-256-bit-secret" )
bearer := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoidG90byIsImV4cCI6OTg2NTQzMjEsInJpZ2h0cyI6eyJyZWFkIjp0cnVlLCJ3cml0ZSI6ZmFsc2UsImV4ZWN1dGUiOnRydWV9fQ.W4pRb-PrV4UH6drgYaFhznvAdqTrN4finADlgs239uLXT0TT_ezKLpcb1mhCO4IB2P36c3I3jIYIyj2UAQ5E0Q"

var playload PlayloadExample
pjwt := new( PJWToken )

if ok:= pjwt.ValidToken( bearer, key ); !ok { panic( "invalid token" ) }
if err := pjwt.ExtractPlayloadFromToken( bearer, &playload ); err != nil  { panic( err ) }

fmt.Println( playload )
```

Will done
```
{toto 98654321 {true false true}}
```

[Try it on jwt.io !](https://jwt.io/#debugger-io?token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoidG90byIsImV4cCI6OTg2NTQzMjEsInJpZ2h0cyI6eyJyZWFkIjp0cnVlLCJ3cml0ZSI6ZmFsc2UsImV4ZWN1dGUiOnRydWV9fQ.W4pRb-PrV4UH6drgYaFhznvAdqTrN4finADlgs239uLXT0TT_ezKLpcb1mhCO4IB2P36c3I3jIYIyj2UAQ5E0Q)