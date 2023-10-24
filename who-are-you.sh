#!/bin/bash

# Fetch JSON data from the URL using curl and parse it with jq
curl -s https://learn.01founders.co/assets/superhero/all.json | jq '.[] | select(.id==70) | .name'
