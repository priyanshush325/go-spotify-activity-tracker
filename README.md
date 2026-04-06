# Spotify Recently Played Widget API 

A simple REST API written in Go using Gin for getting your most recently listened song from Spotify (to display on Personal Site, etc.). 

Built for my personal website as part of my efforts to learn Go (and test the impact of AI on learning computer science fundamentals). You can read more here: https://www.priyanshu.org/blog/go

## Setup

1. Create a Spotify app at https://developer.spotify.com/dashboard
2. Add `http://127.0.0.1:8888/callback` as a redirect URI in your app settings
3. Copy `.env.example` to `.env` and add your `CLIENT_ID` and `CLIENT_SECRET`
4. Run `./setup.sh` to authenticate and generate your tokens

