package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type header struct {
	Alg string `json:"alg"`
	Type string `json:"typ"`
}

type PJWToken struct {}

func ( t *PJWToken )setHeader() ( string, error ) {
	data, err := json.Marshal( header{ Alg: "HS512", Type: "JWT" } )
	if err != nil { return "", err }
	return base64.RawURLEncoding.EncodeToString( data ), nil
}

func ( t *PJWToken )setPlayload( i interface{} )  ( string, error ){
	data, err := json.Marshal( i )
	if err != nil { return "", err }
	return base64.RawStdEncoding.EncodeToString( data ), nil
}

func ( t *PJWToken )forge( head, playload string, key []byte ) string {
	src := fmt.Sprintf( "%s.%s", head, playload )
	h := hmac.New( sha512.New, key )
	h.Write( []byte(src) )
	sha := base64.RawURLEncoding.EncodeToString( h.Sum( nil ) )
	return fmt.Sprintf( "%s.%s", src, sha )
}

func ( t *PJWToken )CreateToken( i interface{}, key []byte ) ( string, error ) {
	header, err := t.setHeader()
	if err != nil { return "", fmt.Errorf( "CreateToken::setHeader() error : %s", err ) }
	playload, err := t.setPlayload( i )
	if err != nil { return "", fmt.Errorf( "CreateToken::setPlayload() error : %s", err )  }
	return t.forge( header, playload, key ), nil
}

func ( t *PJWToken )ValidToken( bearer string, key []byte ) bool {
	parts := strings.Split( bearer, "." )
	if len(parts) != 3 { return false }
	if header, err := t.setHeader(); (header != parts[0]) || err != nil { return false }
	return t.forge( parts[0], parts[1], key ) == bearer
}

func ( t *PJWToken )ExtractPlayloadFromToken( bearer string, i interface{} ) error{
	parts := strings.Split( bearer, "." )
	if len(parts) != 3 { return fmt.Errorf( "bearer format error") }

	playload, err := base64.RawURLEncoding.DecodeString( parts[1] )
	if err != nil { return err }
	return json.Unmarshal( playload, i )
}
