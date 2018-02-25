/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"flag"
	"log"

	"golang.org/x/oauth2/google"

	"github.com/GoogleCloudPlatform/image-rebase/pkg/rebase"
)

const scope = "https://www.googleapis.com/auth/devstorage.read_write"

var (
	robotEmail = flag.String("robot_email", "", "Robot account email address.")
	deriv      = flag.String("deriv", "", "Image to rebase")
	oldBase    = flag.String("old_base", "", "Old base to remove")
	newBase    = flag.String("new_base", "", "New base to replace with")
	rebased    = flag.String("tag", "", "Rebased image to push")
)

func main() {
	flag.Parse()
	ctx := context.Background()

	// TODO: Shell out to docker-credential-gcr to get auth for any registry.
	c, err := google.DefaultClient(ctx, scope)
	if err != nil {
		log.Fatalf("Failed to set up Google auth client: %v", err)
	}
	r := rebase.Rebaser{c}
	if err := r.Rebase(
		rebase.FromString(*deriv),
		rebase.FromString(*oldBase),
		rebase.FromString(*newBase),
		rebase.FromString(*rebased),
	); err != nil {
		log.Fatalf("Rebase: %v", err)
	}
}
