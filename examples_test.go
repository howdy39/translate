// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package translate_test

import (
	"fmt"
	"testing"

	"github.com/howdy39/translate"
	"context"
	"golang.org/x/text/language"
	"strings"
)

func TestNewClient(t *testing.T) {
	ctx := context.Background()
	client, err := translate.NewClient(ctx)
	if err != nil {
		t.Fatal(err)
	}
	// Use the client.

	// Close the client when finished.
	if err := client.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestTranslate(t *testing.T) {
	ctx := context.Background()
	client, err := translate.NewClient(ctx)
	if err != nil {
		t.Fatal(err)
	}
	translations, err := client.Translate(ctx,
		[]string{"こんにちは。お元気ですか？"}, language.English,
		&translate.Options{
			Source: language.Japanese,
			Format: translate.Text,
		})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(translations[0].Text)
}


func TestTranslateJA1000(t *testing.T) {
	ctx := context.Background()
	client, err := translate.NewClient(ctx)
	if err != nil {
		t.Fatal(err)
	}
	translations, err := client.Translate(ctx,
		[]string{strings.Repeat("おはよう。",200)}, language.English,
		&translate.Options{
			Source: language.Japanese,
			Format: translate.Text,
		})
	if err != nil {
		t.Fatal(err)
		/*
		examples_test.go:73: googleapi: got HTTP response code 400 with body: <!DOCTYPE html>
		<html lang=en>
		  <meta charset=utf-8>
		  <meta name=viewport content="initial-scale=1, minimum-scale=1, width=device-width">
		  <title>Error 400 (Bad Request)!!1</title>
		  <style>
		    *{margin:0;padding:0}html,code{font:15px/22px arial,sans-serif}html{background:#fff;color:#222;padding:15px}body{margin:7% auto 0;max-width:390px;min-height:180px;padding:30px 0 15px}* > body{background:url(//www.google.com/images/errors/robot.png) 100% 5px no-repeat;padding-right:205px}p{margin:11px 0 22px;overflow:hidden}ins{color:#777;text-decoration:none}a img{border:0}@media screen and (max-width:772px){body{background:none;margin-top:0;max-width:none;padding-right:0}}#logo{background:url(//www.google.com/images/branding/googlelogo/1x/googlelogo_color_150x54dp.png) no-repeat;margin-left:-5px}@media only screen and (min-resolution:192dpi){#logo{background:url(//www.google.com/images/branding/googlelogo/2x/googlelogo_color_150x54dp.png) no-repeat 0% 0%/100% 100%;-moz-border-image:url(//www.google.com/images/branding/googlelogo/2x/googlelogo_color_150x54dp.png) 0}}@media only screen and (-webkit-min-device-pixel-ratio:2){#logo{background:url(//www.google.com/images/branding/googlelogo/2x/googlelogo_color_150x54dp.png) no-repeat;-webkit-background-size:100% 100%}}#logo{display:inline-block;height:54px;width:150px}
		  </style>
		  <a href=//www.google.com/><span id=logo aria-label=Google></span></a>
		  <p><b>400.</b> <ins>That’s an error.</ins>
		  <p>Your client has issued a malformed or illegal request.  <ins>That’s all we know.</ins>
		*/
	}
	fmt.Println(translations[0].Text)
}
