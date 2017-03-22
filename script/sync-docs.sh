#!/usr/bin/env bash

# Ensures that we run on Travis
if [ "$TRAVIS_BRANCH" != "master" ] || [ "$BUILD_DOCS" != "yes" ] || [ "$TRAVIS_SECURE_ENV_VARS" == "false" ] || [ "$TRAVIS_PULL_REQUEST" != "false" ] ; then
    echo "Must be: a merged pr on the master branch, BUILD_DOCS=yes, TRAVIS_SECURE_ENV_VARS=false"
    exit 0
fi

DOCS_REPO_NAME="kompose"
DOCS_REPO_URL="git@github.com:kubernetes-incubator/kompose.git"
DOCS_KEY="script/deploy_key"
DOCS_USER="Travis CI"
DOCS_BRANCH="gh-pages"
DOCS_FOLDER="docs"

# decrypt the private key
openssl aes-256-cbc -K $encrypted_1069b8f9eeff_key -iv $encrypted_1069b8f9eeff_iv -in "$DOCS_KEY.enc" -out "$DOCS_KEY" -d
chmod 600 "$DOCS_KEY"
eval `ssh-agent -s`
ssh-add "$DOCS_KEY"

# clone the repo
git clone "$DOCS_REPO_URL" "$DOCS_REPO_NAME"

# change to that directory (to prevent accidental pushing to master, etc.)
cd "$DOCS_REPO_NAME"

# switch to gh-pages and grab the docs folder from master
git checkout gh-pages
git checkout master docs

# clean-up the docs and convert to jekyll-friendly docs
cd docs
for filename in *.md; do
    if cat $filename | head -n 1 | grep "\-\-\-";
    then
    echo "$filename already contains Jekyll format"
    else
    echo "Adding Jekyll file format to $filename"
    jekyll="---
layout: default
permalink: /$filename/
---
"
    echo -e "$jekyll\n$(cat $filename)" > $filename
    fi
done
cd ..

# remove README.md from docs folder as it isn't relevant
rm docs/README.md

# add relevant user information
git config user.name "$DOCS_USER"
git config user.email "<>"
git add --all

# Check if anything changed, and if it's the case, push to origin/master.
if git commit -m 'Update docs' -m "Commit: https://github.com/kubernetes-incubator/kompose/commit/$TRAVIS_COMMIT" ; then
  git push
fi

# cd back to the original root folder
..
